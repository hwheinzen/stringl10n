// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

var test = `// Copyright {{.Copyright}}. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// THIS FILE HAS BEEN GENERATED BY {{.Generator}}.
// DO NOT EDIT.
// CHANGES WILL DISAPPEAR AFTER NEXT RUN OF {{.Generator}}.

package {{.Package}}

import (
	"testing"

	"github.com/hwheinzen/stringl10n/message"
)

type InT struct {
	key string
	lang string
}
type tTest struct {
	in  InT
	out string
}

var tTests = []tTest{ {{range $key, $pair := .Texts}}{{range $pair}}
	{
		in: InT{key: "{{$key}}", lang: "{{.Lang}}"},
		out: "{{.Value}}",
	},{{end}}{{end}}
}

func TestTrans(test *testing.T) {
	for _, v := range tTests {
		temp := l10nTrans(v.in.key, v.in.lang)
		if temp != v.out {
			test.Error("Key:"+v.in.key+" Lang:"+v.in.lang+"\nexpected:"+v.out+"\ngot:     "+temp)
		}
	}
}

func TestManualSubst(t *testing.T) { // run with go test -v -run=Manual
	var err error
	for _, v := range tTests {
		err = getErrorMsg(v.in.key)
		trans := l10nTrans(v.in.key, v.in.lang)
		subst := l10nSubst(trans, err.(*message.Message))
		t.Log(subst)
	}
}

func getErrorMsg(s string) error {
	msg := message.New(s){{range .Vars}}
	msg.AddVar("{{.Name}}", {{/*
		--------*/}}{{if eq .Type "int"}}123{{/*
		*/}}{{else}}{{if eq .Type "int8"}}123{{/*
		*/}}{{else}}{{if eq .Type "int16"}}123{{/*
		*/}}{{else}}{{if eq .Type "int32"}}123{{/*
		*/}}{{else}}{{if eq .Type "int64"}}123{{/*
		*/}}{{else}}{{if eq .Type "uint"}}123{{/*
		*/}}{{else}}{{if eq .Type "uint8"}}123{{/*
		*/}}{{else}}{{if eq .Type "uint16"}}123{{/*
		*/}}{{else}}{{if eq .Type "uint32"}}123{{/*
		*/}}{{else}}{{if eq .Type "int64"}}123{{/*
		*/}}{{else}}{{if eq .Type "float32"}}123.456789{{/*
		*/}}{{else}}{{if eq .Type "float64"}}123.456789{{/*
		*/}}{{else}}{{if eq .Type "string"}}"SOMESTRING"{{/*
		*/}}{{else}}{{.Type}}{}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}{{end}}){{end}}
	return msg
}

// THIS FILE HAS BEEN GENERATED BY {{.Generator}}.
// DO NOT EDIT.
// CHANGES WILL DISAPPEAR AFTER NEXT RUN OF {{.Generator}}.
`
