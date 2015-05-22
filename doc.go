// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Command stringl10n (string localization) generates a 
// go source file for localizing text strings.
//
// The code contains:
//  - (global) declarations of the localized string variables
//  - init function that fills these variables with text strings according to the chosen start language code
//  - function that refills these variables with text strings according to a chosen language code
//  - constant containing all locale versions of all localized strings
// 
// Information must be provided via a JSON file:
/*
 {
	"Copyright": "<year> <copyright owner>"
	,"Package":  "example"
	,"GenFile":  "stringl10n_generated.go"
	,"Default":  "en"
	,"Start":    "de"
	,"Text":	[
		{
			"Name":     "localizedStringVariable1"
			,"Locs": [
				{ "Lang": "de", "Value": "deutsche Wörter 1"}
				,{"Lang": "en", "Value": "english words 1"}
				,{"Lang": "fr", "Value": "mots françaises 1"}
			]
		}
		,{
			"Name":     "localizedStringVariable2"
			,"Locs": [
				{ "Lang": "de", "Value": "deutsche Wörter 2"}
				,{"Lang": "en", "Value": "english words 2"}
				,{"Lang": "fr", "Value": "mots françaises 2"}
			]
		}
	]
 }
*/
// Usage:
//  $ stringl10n -json=example.json
//
// Usage inside source code:
//  //go:generate stringl10n -json=example.json
//
// The variables localizedStringVariable1 and localizedStringVariable2 can now
// be used instead of string literals.
package main
