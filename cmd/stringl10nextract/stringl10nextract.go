// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Command stringl10nextract generates output data that can be used
// as input to the string localization tool stringl10n.
// Data must be edited.
//
// Usage:
//  $ stringl10nextract -o=example.txt
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

const (
	pgmname = "stringl10nextract"
	head    = `// This file was generated by stringl10nextract.
//
// Edit to make it the JSON input file for stringl10n.
//
// - Change copyright owner
// - Change package name
// - Change GenFile name (optional)
// - Delete entries which need no translation
// - Add translations
// - Add all variables from template expressions to Vars
// - Remove all comments (JSON doesn't like them)
//   including these introductory lines
// - Save with a new file name
//
{
	"Copyright": "2015 Itts Mee"
	,"Package":  "example"
	,"GenFile":  "stringl10n_generated.go"

	,"Vars": [
			{ "Name": "", "Type": ""}
			,{"Name": "", "Type": ""}
	]

	,"Texts": {
		"Dummy": [] // FIRST ENTRY MUST NOT HAVE A LEADING COMMA
`
	foot = "\t}\n}\n"
)

func main() {
	root, out := args()

	err := createOutFile(root, out)
	if err != nil {
		log.Fatal(err)
	}
}

func createOutFile(dir, out string) (err error) {
	var oFile *os.File
	if out != "" {
		oFile, err = os.Create(out)
		if err != nil {
			return
		}
		defer oFile.Close()
	} else {
		oFile = os.Stdout
	}

	_, err = oFile.Write([]byte(head))
	if err != nil {
		return
	}

	err = filepath.Walk(dir, visitFile) // walk all files
	if err != nil {
		return
	}

	_, err = oFile.Write([]byte(foot))
	if err != nil {
		return
	}

	return
}

func visitFile(path string, fi os.FileInfo, err error) error {
	if err == nil && isGoFile(fi) { // only .go files
		if path == fi.Name() { // no subdirs
			err = processInFile(path)
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func isGoFile(fi os.FileInfo) bool {
	name := fi.Name()

	return !fi.IsDir() &&
		!strings.HasPrefix(name, ".") &&
		strings.HasSuffix(name, ".go") &&
		!strings.Contains(name, "_test.") &&
		!strings.Contains(name, "_generated")
}

func processInFile(path string) (err error) {

	// Create the AST by parsing src.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return
	}

	// Inspect the AST and print string literals.
	ast.Inspect(
		f,
		func(n ast.Node) bool {
			switch x := n.(type) {
			case *ast.GenDecl:
				if x.Tok == token.IMPORT {
					return false
				}
			case *ast.BasicLit:

				if x.Kind == token.STRING {

					runes := []rune(x.Value)
					if len(runes) < 5 { // probably not worth translating
						return false
					}
					ok := false
					for _, v := range runes {
						if unicode.IsLetter(v) { // letters inside?
							ok = true
							break
						}
					}
					if !ok { // No. nothing to translate
						return false
					}

					fmt.Fprintf(oFile, "\t\t,%s: [\t\t// %s\n", x.Value, fset.Position(n.Pos()))
					fmt.Fprintln(oFile, "\t\t\t{ \"Lang\":\"en\",\"Value\":\"\"}")
					fmt.Fprintln(oFile, "\t\t\t,{\"Lang\":\"  \",\"Value\":\"\"}")
					fmt.Fprintln(oFile, "\t\t]")
				}
			}
			return true
		},
	)
	return
}