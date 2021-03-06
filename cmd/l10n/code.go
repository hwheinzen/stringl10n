// Copyright 2020 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

var code = `// Copyright {{.Copyright}}. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// THIS FILE HAS BEEN GENERATED BY {{.Generator}} using {{.Input}}.
// ON {{.Generated}}. DO NOT EDIT.
// CHANGES WILL DISAPPEAR AFTER NEXT RUN OF {{.Generator}}.

/*
 {{.GenFile}} contains all localized strings and
 functions for translating (L10nTranslate),
 replacing text/template variables (L10nReplace),
 and a conveniance function for localizing errors (L10nLocalizeError)
*/

package {{.Package}}

import (
	"bytes"
	"errors"
	"fmt"
	"encoding/json"
	"log"
	"strings"
	"text/template" {{if ne (len .Funcs) 0}}{{range .Funcs}}{{if ne .Path ""}}
	"{{.Path}}"{{end}}{{end}}{{end}}{{if ne (len .Vars) 0}}{{range .Vars}}{{if ne .Path ""}}
	"{{.Path}}"{{end}}{{end}}{{end}}

	. "github.com/hwheinzen/stringl10n/mistake"
)

// Type l10nPair is used during string localization.
type l10nPair struct {
	Lang  string
	Value string
}

// l10nMap contains all key strings and all translations.
var l10nMap = make(map[string][]l10nPair, 10)

// L10nTranslate returns the adequate translation of a given text
// according to the chosen language code.
func L10nTranslate(key, lang string) (out string, err error) {
	fnc := "L10nTranslate"
	
	pairs, ok := l10nMap[key]
	if !ok {
		err := {{.ErrorType}}{
			Fix: "{{.ErrorPref}}:no entry for '{{.NameTemplate}}'",
			Var: []struct {
				Name  string
				Value interface{}
			}{
				{"Name", key},
			},
		}
		return "", fmt.Errorf(fnc+":%w", err)
	}
	for _, v := range pairs {
		if len(lang) >= 5 { // assuming POSIX locales: language + country
			if v.Lang == lang[:5] {
				return v.Value, nil
			}
		}
	}
	for _, v := range pairs {
		if len(lang) >= 2 { // assuming POSIX locales: language only
			if v.Lang == lang[:2] {
				return v.Value, nil
			}
		}
	}

	err = {{.ErrorType}}{
		Fix: "{{.ErrorPref}}:no {{.Nam2Template}} translation for '{{.NameTemplate}}'",
		Var: []struct {
			Name  string
			Value interface{}
		}{
			{"Name", key},
			{"Nam2", lang},
		},
	}
	return "", fmt.Errorf(fnc+":%w", err)
}

// l10nVars declares all variables possibly needed for substitution.
type l10nVars struct { {{range .Vars}}
	{{.Name}} {{.Type}}{{end}}
}

// L10nReplace replaces text/template expressions and returns
// the changed text string. Variables in these text/template 
// expressions are substituted by values.
func L10nReplace(tmpl string, vars []struct {
	Name string
	Value interface{}
}) (out string, err error) {
	fnc := "L10nReplace"

	t := template.New("t"){{if eq (len .Funcs) 0}}
	_, err = t.Parse(tmpl){{else}}
	funcMap := template.FuncMap { {{range .Funcs}}
		"{{.Name}}": {{.Function}},{{end}}
	}
	_, err = t.Funcs(funcMap).Parse(tmpl){{end}}
	if err != nil {
		e := {{.ErrorType}}{
			Fix: "{{.ErrorPref}}:error parsing '{{.NameTemplate}}'",
			Var: []struct {
				Name  string
				Value interface{}
			}{
				{"Name", tmpl},
			},
		}
		return "", fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	allVars := l10nVars{}

{{$ErrorType := .ErrorType}}
{{$ErrorPref := .ErrorPref}}
{{$TypeTemplate := .TypeTemplate}}
{{$NameTemplate := .NameTemplate}}
{{$Nam2Template := .Nam2Template}}

	for _, pair := range vars {
		switch pair.Name { {{range .Vars}}
		case "{{.Name}}":
			v, ok := pair.Value.({{.Type}})
			if !ok {
				e := {{$ErrorType}}{
					Fix: "{{$ErrorPref}}:wrong variable type {{$TypeTemplate}} for '{{$NameTemplate}}', expected: {{$Nam2Template}}",
					Var: []struct {
						Name  string
						Value interface{}
					}{
						{"Type", fmt.Sprintf("%T", pair.Value)},
						{"Name", tmpl},
						{"Nam2", "{{.Type}}"},
					},
				}
				return "", fmt.Errorf(fnc+":%w:"+err.Error(), e)
			}
			allVars.{{.Name}} = v{{end}}
		default:
			err = {{.ErrorType}}{
				Fix: "{{.ErrorPref}}:variable {{.NameTemplate}} not declared",
				Var: []struct {
					Name  string
					Value interface{}
				}{
					{"Name", pair.Name},
				},
			}
			return "", fmt.Errorf(fnc+":%w", err)
		}
	}

	var b bytes.Buffer
	err = t.Execute(&b, allVars)
	if err != nil {
		e := Err{
			Fix: "{{$ErrorPref}}:error executing template {{$NameTemplate}}",
			Var: []struct {Name  string; Value interface{}}{
				{"Name", tmpl},
			},
		}
		return "", fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	return b.String(), nil
}

// L10nLocalizeError takes the innermost wrapped error of in and tries
// to translate the error message and then tries to replace text/template
// expressions with variable values if available.
// It creates a new error and returns it wrapped again.
func L10nLocalizeError(in error, lang string) (out, err error) {
	fnc := "L10nLocalizeError"

	// Unwrap
	var inner, e error
	var ss []string
	for inner, e = in, errors.Unwrap(in); e != nil; inner, e = e, errors.Unwrap(e) {
		ss = append(ss, strings.Replace(inner.Error(), e.Error(), "%w", 1))
	}
	ss = append(ss, inner.Error())

	// Translate
	txt, err := L10nTranslate(inner.Error(), lang)
	if err != nil {
		return nil, fmt.Errorf(fnc+":%w", err)
	}

	// Substitute
	type varser interface {
		Vars() []struct {
			Name string
			Value interface{}
		}
	}
	f, ok := inner.(varser)
	if ok {
		txt, err = L10nReplace(txt, f.Vars())
		if err != nil {
			return nil, fmt.Errorf(fnc+":%w", err)
		}
		out = Err{
			Fix: txt,
			Var: f.Vars(),
		}
	} else {
		out = Err{Fix: txt}
	}

	// Wrap again
	for i := len(ss) - 1; i != 0; i-- {
		out = fmt.Errorf(ss[i-1], out)
	}

	return out, nil
}`
