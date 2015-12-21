// Copyright 2015 Itts Mee. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// THIS FILE HAS BEEN GENERATED BY stringl10n.
// DO NOT EDIT.
// CHANGES WILL DISAPPEAR AFTER NEXT RUN OF stringl10n.

package main

import (
	"testing"

	"github.com/hwheinzen/stringl10n/message"
)

type InT struct {
	key  string
	lang string
}
type tTest struct {
	in  InT
	out string
}

var tTests = []tTest{
	{
		in:  InT{key: "Int: {{printf \"%d\" .Int}} Float: {{printf \"%f\" .Flo}}", lang: "en"},
		out: "Integer: {{printf \"%d\" .Int}} Float: {{trimright (printf \"%f\" .Flo) \"0\"}}",
	},
	{
		in:  InT{key: "Int: {{printf \"%d\" .Int}} Float: {{printf \"%f\" .Flo}}", lang: "de"},
		out: "Gleitkomma: {{trimright (replace (printf \"%f\" .Flo | ) \".\" \",\" -1) \"0\"}} Ganzzahl: {{printf \"%d\" .Int}}",
	},
	{
		in:  InT{key: "Int: {{printf \"%d\" .Int}} Float: {{printf \"%f\" .Flo}}", lang: "xx"},
		out: "strüktür: {{print .Str}}",
	},
}

func TestTrans(test *testing.T) {
	for _, v := range tTests {
		temp := l10nTrans(v.in.key, v.in.lang)
		if temp != v.out {
			test.Error("Key:" + v.in.key + " Lang:" + v.in.lang + "\nexpected:" + v.out + "\ngot:     " + temp)
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
	msg := message.New(s)
	msg.AddVar("Int", 123)
	msg.AddVar("Flo", 123.456789)
	msg.AddVar("Str", Struct{})
	return msg
}

// THIS FILE HAS BEEN GENERATED BY stringl10n.
// DO NOT EDIT.
// CHANGES WILL DISAPPEAR AFTER NEXT RUN OF stringl10n.
