// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Command stringl10n (string localization) generates two 
// go source files that can be included in a Go project.
//
// First source file provides functions for string translations
// and variable substitution and defines an interface:
// 
//  func l10nTranslate(in, lang string) (out string) // alias: t
//  func l10nSubstitute(tmpl string, vars Varser) (out string)
// 
//  type Varser interface {
//      // Vars returns Name-Value-Pairs.
//      Vars() []struct{
//          Name string
//          Value interface{}
//      }
//  }
// 
// Second source file provides a unit test.
// 
// Information must be passed in via a JSON file (e.g. example.json):
/*
 {
	"Copyright": "2015 Itts Mee"
	,"Package":  "example"
	,"GenFile":  "stringl10n_generated.go"

	,"Vars":     [
			{ "Name": "I1", "Type": "int"}
			,{"Name": "F1", "Type": "float64"}
		]

	,"Text":	{
		"programmer's words 1": [
			{ "Lang": "de", "Value": "deutsche Wörter 1"}
			,{"Lang": "en", "Value": "english words 1"}
			,{"Lang": "fr", "Value": "mots françaises 1"}
		]
		,"programmer's words 2": [
			{ "Lang": "de", "Value": "deutsche Wörter 2"}
			,{"Lang": "en", "Value": "english words 2"}
			,{"Lang": "fr", "Value": "mots françaises 2"}
		]
		,"A: {{printf \"%d\" .I1}} B: {{printf \"%d\" .F1}}": [
			{ "Lang":"en","Value":"A: {{printf \"%d\" .I1}} B: {{printf \"%d\" .F1}}"}
			,{"Lang":"de","Value":"A: {{printf \"%d\" .F1}} A: {{printf \"%d\" .I1}}"}
		]
	}
 }
*/
// Usage:
//  $ stringl10n -json=example.json
//
// Usage inside source code:
//  //go:generate stringl10n -json=example.json
package main
