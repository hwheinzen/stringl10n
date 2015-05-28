// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Command stringl10n (string localization) generates two 
// go source files that can be included in a Go project.
//
// - One that provides a function that translates strings:
//  func t(in, lang string) (out string)
// 
// - And one that provides a unit test.
// 
// Information must be passed in via a JSON file (e.g. example.json):
/*
 {
	"Copyright": "2015 Itts Mee"
	,"Package":  "example"
	,"GenFile":  "stringl10n_generated.go"
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
	}
 }
*/
// Usage:
//  $ stringl10n -json=example.json
//
// Usage inside source code:
//  //go:generate stringl10n -json=example.json
package main
