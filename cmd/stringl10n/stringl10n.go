// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/dullgiulio/jsoncomments"
)

const (
	pgmname = "stringl10n"
)

type All struct {
	Copyright string
	Package   string
	GenFile   string
	Vars      []struct {
		Name string
		Type string
	}
	Funcs []struct {
		Name     string
		Function string
		Path     string
	}
	Texts map[string][]struct {
		Lang  string
		Value string
	}
	// ------------- computed values
	Generator string
}

func main() {

	filename := args()

	all, err := fillStruct(filename)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}

	var name string

	// code
	name, err = makeCode(all)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}
	err = addJSON(name, all)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}
	err = gofmt(name)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}

	// testcode
	name, err = makeTestCode(all)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}
	err = gofmt(name)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}
}

func gofmt(filename string) (err error) {
	cmd := exec.Command("gofmt", "-w", filename)
	err = cmd.Run()
	return
}

func fillStruct(filename string) (all All, err error) {

	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	reader := jsoncomments.NewReader(file) // filters #-comments
	dec := json.NewDecoder(reader)

	err = dec.Decode(&all)
	if err != nil {
		return
	}

	if all.Copyright == "" {
		err = errors.New("Copyright is missing")
		return
	}
	if all.Package == "" {
		err = errors.New("Package is missing")
		return
	}
	if all.Texts == nil || len(all.Texts) == 0 {
		err = errors.New("Texts are missing")
		return
	}

	all.Generator = pgmname

	return
}

func makeCode(all All) (filename string, err error) {

	filename = all.GenFile

	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	t := template.New("code")
	// parse code template
	_, err = t.Parse(code)
	if err != nil {
		return
	}
	// create code from template
	err = t.Execute(file, all)
	if err != nil {
		return
	}

	return
}

func addJSON(filename string, all All) (err error) {

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		return
	}
	defer file.Close()

	// begin raw string
	_, err = file.Write([]byte(`
	
// init fills the translation map.
func init() {
	var l10nJSON = `))
	if err != nil {
		return
	}

	_, err = file.Write([]byte("`"))
	if err != nil {
		return
	}

	// turn map Texts into JSON, indent slightly
	bytes, err := json.MarshalIndent(all.Texts, "", " ")
	if err != nil {
		return
	}
	// write JSON
	_, err = file.Write(bytes)
	if err != nil {
		return
	}

	// end raw string
	_, err = file.Write([]byte("`"))
	if err != nil {
		return
	}
	_, err = file.Write([]byte(`

	err := json.Unmarshal([]byte(l10nJSON), &l10nMap)
	if err != nil {
		log.Fatalln(err)
	}
	l10nJSON = "" // no longer needed
}`))
	if err != nil {
		return
	}
	//_, err = file.Write([]byte("`\n"))
	//if err != nil {
	//	return
	//}

	return
}

func makeTestCode(all All) (filename string, err error) {

	filename = all.GenFile[:len(all.GenFile)-3] + "_test.go"

	// prepare Texts of local structure All
	for k, v := range all.Texts {
		for i, val := range v {
			v[i].Value = strings.Replace(val.Value, `"`, "\\\"", -1)
		}
		delete(all.Texts, k)
		k = strings.Replace(k, `"`, "\\\"", -1)
		all.Texts[k] = v
	}

	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	t := template.New("test")
	// parse testcode template
	_, err = t.Parse(test)
	if err != nil {
		return
	}
	// write testcode
	err = t.Execute(file, all)
	if err != nil {
		return
	}

	return
}
