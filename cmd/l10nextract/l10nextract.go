// Copyright 2020 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Command stringl10nextract generates output data that can be used
// as input to the string localization tool stringl10n.
//
// Output data must be edited!
//
// Usage:
//  $ find . | l10nextract > tmp.json
// or:
//  $ l10nextract -input-dir=. -update=l10n.json
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
	"regexp"
	"strings"

	. "github.com/hwheinzen/stringl10n/mistake"
)

const pgm = "l10nextract"

var gPoss = make(map[string][]string, 50) // occurrences per text

// buildtime serves 'l10n -version' if l10n was built with:
// -ldflags "-X 'main.buildtime=`date -Iseconds`'"
var buildtime string

func main() {

	args(buildtime) // all arguments are global

	err := readJSON()
	if err != nil {
		err = translate(err, argLang) // ******** l10n ********
		log.Fatalln(pgm + ":" + err.Error())
	}

	err = extract()
	if err != nil {
		err = translate(err, argLang) // ******** l10n ********
		log.Fatalln(pgm + ":" + err.Error())
	}

	err = addVars()
	if err != nil {
		err = translate(err, argLang) // ******** l10n ********
		log.Fatalln(pgm + ":" + err.Error())
	}

	err = addFuncs()
	if err != nil {
		err = translate(err, argLang) // ******** l10n ********
		log.Fatalln(pgm + ":" + err.Error())
	}

	bytes, err := makeJSON()
	if err != nil {
		err = translate(err, argLang) // ******** l10n ********
		log.Fatalln(pgm + ":" + err.Error())
	}

	err = createOutFile(bytes)
	if err != nil {
		err = translate(err, argLang) // ******** l10n ********
		log.Fatalln(pgm + ":" + err.Error())
	}
}

func extract() error {
	fnc := "extract"

	if argDir != "" { // examine files
		err := filepath.Walk(argDir, visitFile)
		if err != nil {
			return fmt.Errorf(fnc+":%w", err)
		}
	} else { // examine stdin
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			path := s.Text()
			fi, err := os.Stat(path)
			if err != nil {
				e := Err{
					Fix: "L10NEXTRACT:get fileinfo for {{.Name}} failed",
					Var: []struct {
						Name  string
						Value interface{}
					}{
						{"Name", path},
					},
				}
				return fmt.Errorf(fnc+":%w:"+err.Error(), e)
			}
			err = visitFile(path, fi, err)
			if err != nil {
				return fmt.Errorf(fnc+":%w", err)
			}
		}
	}

	return nil
}

func visitFile(path string, fi os.FileInfo, inErr error) error {
	fnc := "visitFile"

	if inErr != nil {
		return fmt.Errorf(fnc+":%w", inErr)
	}

	if len(path) > 1 && path[:2] == "./" {
		path = path[2:]
	}

	if isGoFile(fi) {
		if path == fi.Name() || argDeep {
			// process file or dive into subdir
			//
			// processInFile returns error from parser.ParseFile
			// which makes filepath.Walk go on to the next ...
			err := processInFile(path)
			if err != nil {
				return fmt.Errorf(fnc+":%w", err)
			}
		}
	}
	return nil
}

func isGoFile(fi os.FileInfo) bool {
	name := fi.Name()
	return !fi.IsDir() &&
		!strings.HasPrefix(name, ".") &&
		strings.HasSuffix(name, ".go") &&
		!strings.HasSuffix(name, "test.go")
}

func processInFile(path string) error {
	fnc := "processInFile"

	// create AST
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil { // directories return error!
		e := Err{
			Fix: "L10NEXTRACT:parse file {{.Name}} failed",
			Var: []struct {
				Name  string
				Value interface{}
			}{
				{"Name", path},
			},
		}
		return fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	// inspect AST
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
					return processString( // process string literal
						n,
						fset,
						x.Value[1:len(x.Value)-1],
					)
				}
			}
			return true
		},
	)
	return nil
}

// processString adds relevant strings JSON-formatted to gData.
func processString(n ast.Node, fset *token.FileSet, s string) bool {
	//	fnc := "processString"

	runes := []rune(s)
	if len(runes) < argMin+2 { // probably not worth translating
		return false
	}
	if len(runes) > argMax+2 { // probably not intended for translation
		return false
	}

	if argKeywords != nil {
		ok := false
		for _, v := range argKeywords {
			if strings.Contains(s, v) { // keyword inside
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}

	if argRgx != nil && !argRgx.MatchString(s) { // match
		return false
	}

	_, ok := gData.Texts[s]
	if !ok { // add new item
		gData.Texts[s] = []struct {
			Lang  string
			Value string
		}{{Lang: "en", Value: ""}}
	}

	// memorize position as #-comment
	gPoss[s] = append(gPoss[s], "\t\t# "+fmt.Sprint(fset.Position(n.Pos())))

	return true
}

func addVars() (err error) {
	//	fnc := "addVars"

	vars := make([]string, 0, 5)
	for key, _ := range gData.Texts {
		rxp := regexp.MustCompile(`{{\.[A-Z][a-z0-9]*}}`)
		poss := rxp.FindAllIndex([]byte(key), -1)
		for _, pos := range poss {
			vars = append(vars, key[pos[0]+3:pos[1]-2])
		}
	}

	for _, v := range vars {
		var found bool
		for _, w := range gData.Vars {
			if w.Name == v {
				found = true
			}
		}
		if !found {
			gData.Vars = append(gData.Vars, struct {
				Name string
				Type string
				Path string `json:"Path,omitempty"`
			}{Name: v})
		}
	}

	return nil
}

func addFuncs() (err error) {
	//	fnc := "addFuncs"

	// ?

	return nil
}

func createOutFile(bytes []byte) (err error) {
	fnc := "createOutFile"

	var file *os.File
	if argOut != "" {
		file, err = os.Create(argOut) // to file
		if err != nil {
			e := Err{
				Fix: "L10NEXTRACT:create file {{.Name}} failed",
				Var: []struct {
					Name  string
					Value interface{}
				}{
					{"Name", argOut},
				},
			}
			return fmt.Errorf(fnc+":%w:"+err.Error(), e)
		}
		defer file.Close()
	} else {
		file = os.Stdout // to stdout
	}

	_, err = fmt.Fprint(file, head) // print leading comment lines
	if err != nil {
		e := Err{
			Fix: "L10NEXTRACT:print to file {{.Name}} failed at: {{.Nam2}}",
			Var: []struct {
				Name  string
				Value interface{}
			}{
				{"Name", argOut},
				{"Nam2", head},
			},
		}
		return fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	lines := strings.Split(string(bytes), "\n")
	var actTexts bool
	var comments []string

	for _, line := range lines {
		s := strings.TrimSpace(line)

		if strings.HasPrefix(s, "\"Texts\": {") { // start of Texts
			actTexts = true
		}

		if actTexts && strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\": [") { // get position comments at start of item
			key := strings.TrimLeft(s, "\"")
			key = strings.TrimRight(key, "\": [")
			comments = gPoss[key]
		}

		if actTexts && strings.HasPrefix(s, "]") { // print at end of item
			for _, c := range comments {
				_, err = fmt.Fprintln(file, c)
				if err != nil {
					e := Err{
						Fix: "L10NEXTRACT:print to file {{.Name}} failed at: {{.Nam2}}",
						Var: []struct {
							Name  string
							Value interface{}
						}{
							{"Name", argOut},
							{"Nam2", c},
						},
					}
					return fmt.Errorf(fnc+":%w:"+err.Error(), e)
				}
			}
		}

		_, err = fmt.Fprintln(file, line) // print current line
		if err != nil {
			e := Err{
				Fix: "L10NEXTRACT:print to file {{.Name}} failed at: {{.Nam2}}",
				Var: []struct {
					Name  string
					Value interface{}
				}{
					{"Name", argOut},
					{"Nam2", line},
				},
			}
			return fmt.Errorf(fnc+":%w:"+err.Error(), e)
		}
	}

	return nil
}
