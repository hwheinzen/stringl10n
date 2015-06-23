// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"text/template"
)

const (
	pgmname = "stringl10n"
)

type Pair struct {
	Lang  string
	Value string
}
type All struct {
	Copyright string
	Package   string
	GenFile   string
	Text      map[string][]Pair
	MapLen    int
}

func main() {
	fn := args()
	all := decodeFile(fn)

	makeCode(all)
	makeTestCode(all)
}

func decodeFile(fn string) (all All) {

	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&all)
	if err != nil {
		log.Fatal(err)
	}

	if all.Copyright == "" {
		err = errors.New("Copyright is missing")
		log.Fatal(err)
	}
	if all.Package == "" {
		err = errors.New("Package is missing")
		log.Fatal(err)
	}
	if all.GenFile == "" {
		err = errors.New("GenFile is missing")
		log.Fatal(err)
	}

	all.MapLen = len(all.Text)

	return
}

func makeCode(all All) {

	out, err := os.Create(all.GenFile)
	if err != nil {
		log.Fatal(err)
	}
	t := template.New("tmpl")
	_, err = t.Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(out, all) // create code
	if err != nil {
		log.Fatal(err)
	}
	err = out.Close()
	if err != nil {
		log.Fatal(err)
	}

	bytes, err := json.Marshal(all.Text) // create data for dynamic resetting of the locale
	if err != nil {
		log.Fatal(err)
	}
	out, err = os.OpenFile(all.GenFile, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	_, err = out.Write([]byte("`"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = out.Write(bytes) // append to code
	if err != nil {
		log.Fatal(err)
	}
	_, err = out.Write([]byte("`\n"))
	if err != nil {
		log.Fatal(err)
	}
	err = out.Close()
	if err != nil {
		log.Fatal(err)
	}
}

var tmpl = `// Copyright {{.Copyright}}. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// THIS FILE HAS BEEN GENERATED BY stringl10n.
// DO NOT EDIT.
// CHANGES WILL BE LOST AFTER NEXT //go:generate stringl10n ...

/*
 {{.GenFile}} contains all localized strings
 and a translation function.
*/

package {{.Package}}

import (
	"encoding/json"
	"log"
)

// Type L10nPair is used during string localization.
type L10nPair struct {
	Lang  string
	Value string
}

// L10nMap contains all key strings and all translations.
var L10nMap = make(map[string][]L10nPair, {{.MapLen}})

// t returns a localized text for a given variable
// according to the chosen locale language code.
func t(key, lang string) string {
	pairs, ok := L10nMap[key]
	if !ok {
		log.Print("No entry for:", key)
		return key
	}
	for _, v := range pairs {
		if v.Lang == lang {
			return v.Value
		}
	}
	log.Print("No entry for (key/lang):", key, "/", lang)
	return key
}

func init() {
	err := json.Unmarshal([]byte(l10nJSON), &L10nMap)
	if err != nil {
		log.Fatal(err)
	}
	l10nJSON = "" // no longer needed
}

var l10nJSON = `

func makeTestCode(all All) {

	out, err := os.Create(
		all.GenFile[:len(all.GenFile)-3] + "_test.go",
	)
	if err != nil {
		log.Fatal(err)
	}
	t := template.New("testTmpl")
	_, err = t.Parse(testTmpl)
	if err != nil {
		log.Fatal(err)
	}
	err = t.Execute(out, all) // create code
	if err != nil {
		log.Fatal(err)
	}
	err = out.Close()
	if err != nil {
		log.Fatal(err)
	}
}

var testTmpl = `// Copyright {{.Copyright}}. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// THIS FILE HAS BEEN GENERATED BY stringl10n.
// DO NOT EDIT.
// CHANGES WILL BE LOST AFTER NEXT //go:generate stringl10n ...

package {{.Package}}

import (
	"testing"
)

type InT struct {
	key string
	lang string
}
type tTest struct {
	in  InT
	out string
}

var tTests = []tTest{ {{/*
*/}}{{range $key, $value := .Text}}{{/*
	*/}}{{range $i, $v := $value}}
	{
		in: InT{key: "{{$key}}", lang: "{{$v.Lang}}"},
		out: "{{$v.Value}}",
	},{{/*
	*/}}{{end}}{{/*
*/}}{{end}}
}

func TestT(test *testing.T) {
	for _, v := range tTests {
		temp := t(v.in.key, v.in.lang)
		if temp != v.out {
			test.Error("Key:"+v.in.key+" Lang:"+v.in.lang+"\nexpected:"+v.out+"\ngot:     "+temp)
		}
	}
}
`
