// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Command stringl10n (string localization) generates two
// go source files that can be included in a Go project.
//
// First source file provides functions for string translations
// and variable substitution and defines an interface:
//
//  func l10nTrans(in, lang string) (out string)
//  func l10nRepl(tmpl string, vars []struct {
//          Name string
//          Value interface{}
//  }) (out string)
//
///Still there but no longer needed:
/// func l10nSubst(tmpl string, vars Varser) (out string)
/// type Varser interface {
///     Vars() []struct{ // Vars returns Name-Value-Pairs.
///         Name string
///         Value interface{}
///     }
/// }
///
//
// Second source file provides a unit tests.
//
// Information must be passed in via a JSON file (e.g. example.json):
/*
 {
	 "Copyright": "2015 Itts Mee"
	,"Package":  "example"
	,"GenFile":  "stringl10n_generated.go"

	,"Vars": [
		 {"Name": "FNum", "Type": "float64"}
	]

	,"Funcs": [
		 {"Name": "replace", "Function": "strings.Replace", "Path": "strings"}
	]

	,"Texts": {
		 "programmer's words": [
			 {"Lang": "en", "Value": "english words"}
			,{"Lang": "de", "Value": "deutsche Wörter"}
			,{"Lang": "fr", "Value": "mots françaises"}
		]
		,"float: {{printf \"%f\" .FNum}}": [
			 {"Lang":"en","Value":"floating-point number: {{printf \"%f\" .FNum}}"}
			,{"Lang":"de","Value":"Gleitkommazahl: {{replace (printf \"%f\" .FNum) \".\" \",\" -1}}"} # DezimalKOMMA
		]
	}
 }
*/
//
// • Add translations to Texts.
//
// • Add variables from template expressions to Vars.
//
// • Add functions from template expressions to Funcs, but identical paths only once. Path means import path.
//
//
// Usage:
//  $ stringl10n -json=example.json
//
// Usage inside source code:
//  //go:generate stringl10n -json=example.json
package main
