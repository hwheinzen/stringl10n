// Copyright 2020-21 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/template"
	"time"

	"github.com/dullgiulio/jsoncomments"
)

const pgm = "l10n"

type All struct {
	Copyright string
	Generator string
	Generated string
	Input     string
	Package   string
	GenFile   string
	Vars      []struct {
		Name string
		Type string
		Path string
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
	// ---
	TypeTemplate string
	NameTemplate string
	Nam2Template string
	ValTemplate  string
	Val2Template string
}

// buildtime serves 'l10n -version' if l10n was built with:
// -ldflags "-X 'main.buildtime=`date -Iseconds`'"
var buildtime string

func main() {

	jsonFile := args(buildtime)

	all, err := getAll(jsonFile)
	if err != nil {
		log.Fatalln(pgm + ":" + err.Error())
	}

	codeFn, err := makeCode(all)
	if err != nil {
		log.Fatalln(pgm + ":" + err.Error())
	}

	err = gofmt(codeFn)
	if err != nil {
		log.Fatalln(pgm + ":" + err.Error())
	}
}

func gofmt(fn string) error {
	fnc := "gofmt"

	cmd := exec.Command("gofmt", "-w", fn)
	stdOutErr, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf(fnc+":"+string(stdOutErr)+"\n"+fnc+":%w", err)
	}

	return nil
}

func getAll(jsonFile string) (all All, err error) {
	fnc := "getAll"

	file, err := os.Open(jsonFile)
	if err != nil {
		e := errors.New("L10N:open "+jsonFile+" failed")
		return all, fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}
	defer file.Close()

	reader := jsoncomments.NewReader(file) // filters #-comments
	dec := json.NewDecoder(reader)

	err = dec.Decode(&all)
	if err != nil {
		e := errors.New("L10N:decode JSON from "+jsonFile+" failed")
		return all, fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	if all.Copyright == "" {
		err = errors.New("L10N:'Copyright' missing in "+jsonFile)
		return all, fmt.Errorf(fnc+":%w", err)
	}
	if all.Package == "" {
		err = errors.New("L10N:'Package' missing in "+jsonFile)
		return all, fmt.Errorf(fnc+":%w", err)
	}
	if all.Texts == nil || len(all.Texts) == 0 {
		err = errors.New("L10N:'Texts' missing in "+jsonFile)
		return all, fmt.Errorf(fnc+":%w", err)
	}

	all.Generator = pgm
	all.Generated = time.Now().String()[:40]
	all.Input = jsonFile

	all.TypeTemplate = "{{.Type}}"
	all.NameTemplate = "{{.Name}}"
	all.Nam2Template = "{{.Nam2}}"
	all.ValTemplate = "{{.Val}}"
	all.Val2Template = "{{.Val2}}"

	return all, nil
}

func makeCode(all All) (codeFn string, err error) {
	fnc := "makeCode"

	codeFn = all.GenFile
	file, err := os.Create(codeFn)
	if err != nil {
		e := errors.New("L10N:create file "+codeFn+" failed")
		return codeFn, fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}
	defer file.Close()

	t := template.New("t")
	_, err = t.Parse(code) // code.go
	if err != nil {
		e := errors.New("L10N:parse template 'code' failed")
		return codeFn, fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}
	err = t.Execute(file, all)
	if err != nil {
		e := errors.New("L10N:execute template 'code' failed")
		return codeFn, fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	err = addTexts(file, all)
	if err != nil {
		return codeFn, fmt.Errorf(fnc+":%w", err)
	}

	return codeFn, nil
}

func addTexts(file *os.File, all All) error {
	fnc := "addTexts"

	// begin init function
	_, err := file.Write([]byte(`
	
// init fills the translation map.
func init() {
	fnc := "init"

	var l10nJSON = `))

	if err != nil {
		e := errors.New("L10N:write to file "+all.GenFile+" failed")
		return fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	_, err = file.Write([]byte("`")) // raw string delimiter
	if err != nil {
		e := errors.New("L10N:write to file "+all.GenFile+" failed")
		return fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}
	bytes, err := json.MarshalIndent(all.Texts, "", " ") // JSON + indent
	if err != nil {
		e := errors.New("L10N:write to file "+all.GenFile+" failed")
		return fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}
	_, err = file.Write(bytes) // write JSON
	if err != nil {
		e := errors.New("L10N:write to file "+all.GenFile+" failed")
		return fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}
	_, err = file.Write([]byte("`")) // raw string delimiter
	if err != nil {
		e := errors.New("L10N:write to file "+all.GenFile+" failed")
		return fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	// end init function
	_, err = file.Write([]byte(`

	err := json.Unmarshal([]byte(l10nJSON), &l10nMap)
	if err != nil {
		e := Err{Fix: "L10N:error unmarshaling 'l10nJSON'"}
		log.Fatalln(fnc+":%w:"+err.Error(), e)
	}
	l10nJSON = "" // no longer needed
}
// THIS FILE HAS BEEN GENERATED.
// DO NOT EDIT.
// CHANGES WILL DISAPPEAR AFTER NEXT RUN OF go generate.
`))

	if err != nil {
		e := errors.New("L10N:write to file "+all.GenFile+" failed")
		return fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	return nil
}
