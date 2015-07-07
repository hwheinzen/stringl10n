// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Command stringl10nextract generates output data that can be used
// as input to the string localization tool stringl10n.
// Data must be edited.
//
// Usage:
//  $ find . | stringl10nextract > example.txt
// or:
//  $ stringl10nextract -root=. -o=example.txt
package main

import (
	"bufio"
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
// - Add translations to Texts
// - Add variables from template expressions to Vars
//  (if needed)
// - Add functions from template expressions to Funcs
//   (if needed)
//   (identical paths only once)
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

	,"Funcs": [
			{ "Name": "", "Function": "", "Path": ""}
			,{"Name": "", "Function": ""}
	]

	,"Texts": {
		"Dummy": [] // FIRST ENTRY MUST NOT HAVE A LEADING COMMA
`
	foot = "\t}\n}\n"
)

var oFile *os.File

func main() {
	args()
	err := createOutFile()
	if err != nil {
		log.Fatal(err)
	}
}

func createOutFile() (err error) {

	if argOut != "" {
		oFile, err = os.Create(argOut)
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

	if argRoot != "" { // walk from argRoot
		err = filepath.Walk(argRoot, visitFile)
		if err != nil {
			return
		}
	} else {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() { // use stdin
			path := s.Text()
			fi, err := os.Stat(path)
			if err != nil {
				return err
			}
			err = visitFile(path, fi, err)
		}
	}

	_, err = oFile.Write([]byte(foot))
	if err != nil {
		return
	}

	return
}

func visitFile(path string, fi os.FileInfo, err error) error {
	if len(path) > 1 && path[:2] == "./" {
		path = path[2:]
	}
	if err == nil && isGoFile(fi) {
		if argDeep || path == fi.Name() { // dive into subdir or not
			err = processInFile(path)
		}
	}
	return err
}

func isGoFile(fi os.FileInfo) bool {
	name := fi.Name()

	return !fi.IsDir() &&
		!strings.HasPrefix(name, ".") &&
		strings.HasSuffix(name, ".go") &&
		!strings.Contains(name, "test.") &&
		!strings.Contains(name, "_generated")
}

func processInFile(path string) (err error) {

	// Create the AST by parsing src.
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return
	}

	// Inspect the AST and process string literals.
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
					return processString(n, fset, x.Value)
				}
			}
			return true
		},
	)
	return
}

// processString prints relevant strings JSON-formatted to the output file.
func processString(n ast.Node, fset *token.FileSet, s string) bool {
	runes := []rune(s)
	if len(runes) < argMin+2 { // probably not worth translating
		return false
	}
	if len(runes) > argMax+2 { // probably not intended for translation
		return false
	}

	ok := false
	for _, v := range runes {
		if unicode.IsLetter(v) { // letter inside?
			ok = true
			break
		}
	}
	if !ok {
		return false
	}

	if argKeywords != nil {
		ok = false
		for _, v := range argKeywords {
			if strings.Contains(s, v) { // keyword inside?
				ok = true
				break
			}
		}
	}
	if !ok {
		return false
	}

	fmt.Fprintf(oFile, "\t\t,%s: [\t\t// %s\n", s, fset.Position(n.Pos()))
	fmt.Fprintln(oFile, "\t\t\t{ \"Lang\":\"en\",\"Value\":\"\"}")
	fmt.Fprintln(oFile, "\t\t\t,{\"Lang\":\"de\",\"Value\":\"\"}")
	fmt.Fprintln(oFile, "\t\t]")

	return true
}
