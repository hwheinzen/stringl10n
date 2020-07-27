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
	"time"

	"github.com/dullgiulio/jsoncomments"
)

const (
	pgmname = "stringl10n"
)

type All struct {
	Copyright string
	Generator string
	Generated string
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
}

func main() {
	filename := args()

	all, err := getAll(filename)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}

	code, err := makeCode(all)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}

	testcode, err := makeTestcode(all)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}

	err = gofmt(code)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}
	err = gofmt(testcode)
	if err != nil {
		log.Fatalln(pgmname+":", err)
	}
}

func gofmt(filename string) error {
	cmd := exec.Command("gofmt", "-w", filename)
	err := cmd.Run()
	return err
}

func getAll(filename string) (all All, err error) {

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
	all.Generated = time.Now().String()[:40]

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
	_, err = t.Parse(code)
	if err != nil {
		return
	}
	err = t.Execute(file, all)
	if err != nil {
		return
	}

	err = addJSON(file, all)
	if err != nil {
		return
	}

	return
}

func addJSON(file *os.File, all All) (err error) {

	// begin init function
	_, err = file.Write([]byte(`
	
// init fills the translation map.
func init() {
	var l10nJSON = `))
	if err != nil {
		return
	}

	_, err = file.Write([]byte("`")) // raw string delimiter
	if err != nil {
		return
	}
	bytes, err := json.MarshalIndent(all.Texts, "", " ") // make JSON + indent
	if err != nil {
		return
	}
	_, err = file.Write(bytes) // write JSON
	if err != nil {
		return
	}
	_, err = file.Write([]byte("`")) // raw string delimiter
	if err != nil {
		return
	}

	// end init function
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

	return
}

func makeTestcode(all All) (filename string, err error) {

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
	_, err = t.Parse(test)
	if err != nil {
		return
	}
	err = t.Execute(file, all)
	if err != nil {
		return
	}

	return
}
