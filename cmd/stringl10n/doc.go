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

    ,"Vars": [
            { "Name": "I1", "Type": "int"}
            ,{"Name": "F1", "Type": "float64"}
        ]

    ,"Funcs": [
            { "Name": "replace", "Function": "strings.Replace", "Path": "strings"}
    ]

    ,"Text": {
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
        ,"A: {{printf \"%d\" .I1}} B: {{printf \"%f\" .F1}}": [
            { "Lang":"en","Value":"A: {{printf \"%d\" .I1}} B: {{printf \"%f\" .F1}}"}
            ,{"Lang":"de","Value":"A: {{replace (printf \"%f\" .F1) \".\" \",\" -1}} A: {{printf \"%d\" .I1}}"}
        ]
    }
 }
*/
// • Change copyright owner.
//
// • Change package name.
//
// • Change GenFile name (optional).
//
// • Delete entries which need no translation.
//
// • Add translations to Texts.
//
// • Add variables from template expressions to Vars.
//
// • Add functions from template expressions to Funcs, but identical paths only once.
//
// • Remove all comments including these introductory lines (JSON doesn't like them).
//
//
// Usage:
//  $ stringl10n -json=example.json
//
// Usage inside source code:
//  //go:generate stringl10n -json=example.json
package main
