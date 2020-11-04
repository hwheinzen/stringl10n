// Copyright 2020 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// THIS FILE HAS BEEN GENERATED BY l10n using l10n.json.
// ON 2020-11-04 11:10:07.265868711 +0100 CET . DO NOT EDIT.
// CHANGES WILL DISAPPEAR AFTER NEXT RUN OF l10n.

/*
 l10n_generated.go contains all localized strings and
 functions for translating (L10nTranslate),
 replacing text/template variables (L10nReplace),
 and a conveniance function for localizing errors (L10nLocalizeError)
*/

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"text/template"

	. "github.com/hwheinzen/stringl10n/mistake"
)

// Type l10nPair is used during string localization.
type l10nPair struct {
	Lang  string
	Value string
}

// l10nMap contains all key strings and all translations.
var l10nMap = make(map[string][]l10nPair, 10)

// L10nTranslate returns a translation of a given text
// according to the chosen language code.
func L10nTranslate(key, lang string) (out string, err error) {
	fnc := "L10nTranslate"

	pairs, ok := l10nMap[key]
	if !ok {
		err := Err{
			Fix: "L10N:no entry for '{{.Name}}'",
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

	err = Err{
		Fix: "L10N:no {{.Nam2}} translation for '{{.Name}}'",
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

// l10nVars declares all possible variables needed for substitution.
type l10nVars struct {
	Name string
	Nam2 string
	Type string
}

// L10nReplace replaces text/template expressions and returns
// the changed text string. Variables in these text/template
// expressions are substituted by values.
func L10nReplace(tmpl string, vars []struct {
	Name  string
	Value interface{}
}) (txt string, err error) {
	fnc := "L10nReplace"

	t := template.New("t")
	_, err = t.Parse(tmpl)
	if err != nil {
		e := Err{
			Fix: "L10N:error parsing '{{.Name}}'",
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

	for _, pair := range vars {
		switch pair.Name {
		case "Name":
			v, ok := pair.Value.(string)
			if !ok {
				e := Err{
					Fix: "L10N:wrong variable type {{.Type}} for '{{.Name}}', expected: {{.Nam2}}",
					Var: []struct {
						Name  string
						Value interface{}
					}{
						{"Type", fmt.Sprintf("%T", pair.Value)},
						{"Name", tmpl},
						{"Nam2", "string"},
					},
				}
				return "", fmt.Errorf(fnc+":%w:"+err.Error(), e)
			}
			allVars.Name = v
		case "Nam2":
			v, ok := pair.Value.(string)
			if !ok {
				e := Err{
					Fix: "L10N:wrong variable type {{.Type}} for '{{.Name}}', expected: {{.Nam2}}",
					Var: []struct {
						Name  string
						Value interface{}
					}{
						{"Type", fmt.Sprintf("%T", pair.Value)},
						{"Name", tmpl},
						{"Nam2", "string"},
					},
				}
				return "", fmt.Errorf(fnc+":%w:"+err.Error(), e)
			}
			allVars.Nam2 = v
		case "Type":
			v, ok := pair.Value.(string)
			if !ok {
				e := Err{
					Fix: "L10N:wrong variable type {{.Type}} for '{{.Name}}', expected: {{.Nam2}}",
					Var: []struct {
						Name  string
						Value interface{}
					}{
						{"Type", fmt.Sprintf("%T", pair.Value)},
						{"Name", tmpl},
						{"Nam2", "string"},
					},
				}
				return "", fmt.Errorf(fnc+":%w:"+err.Error(), e)
			}
			allVars.Type = v
		default:
			err = Err{
				Fix: "L10N:variable {{.Name}} not declared",
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
			Fix: "L10N:error executing template {{.Name}}",
			Var: []struct {
				Name  string
				Value interface{}
			}{
				{"Name", tmpl},
			},
		}
		return "", fmt.Errorf(fnc+":%w:"+err.Error(), e)
	}

	return b.String(), nil
}

// L10nLocalizeError takes the innermost wrapped error of in and tries
// to translate the error message and then tries to replace text/template
// expressions with variable values if available. It creates a new error
// and returns it wrapped with the input error message as prefix.
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
			Name  string
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
}

// init fills the translation map.
func init() {
	fnc := "init"

	var l10nJSON = `{
 "ARGS:{{.Name}}:unknown version": [
  {
   "Lang": "en",
   "Value": "{{.Name}}: unknown version"
  },
  {
   "Lang": "de",
   "Value": "{{.Name}}: Version unbekannt"
  }
 ],
 "ARGS:{{.Name}}:version of {{.Nam2}}": [
  {
   "Lang": "en",
   "Value": "{{.Name}}: version of {{.Nam2}}"
  },
  {
   "Lang": "de",
   "Value": "{{.Name}}: Version vom {{.Nam2}}"
  }
 ],
 "ARGS:{{.Name}}:{{.Nam2}} argument missing": [
  {
   "Lang": "en",
   "Value": "{{.Name}}: {{.Nam2}} argument missing"
  },
  {
   "Lang": "de",
   "Value": "{{.Name}}: Argument {{.Nam2}} fehlt"
  }
 ],
 "Dummy": [
  {
   "Lang": "Dummy",
   "Value": "Dummy"
  }
 ],
 "L10N:create file {{.Name}} failed": [
  {
   "Lang": "en",
   "Value": "create file {{.Name}} failed"
  },
  {
   "Lang": "de",
   "Value": "Anlegen der Datei {{.Name}} fehlgeschlagen"
  }
 ],
 "L10N:decode JSON from {{.Name}} failed": [
  {
   "Lang": "en",
   "Value": "decode JSON from {{.Name}} failed"
  },
  {
   "Lang": "de",
   "Value": "JSON-Dekodierung von {{.Name}} fehlgeschlagen"
  }
 ],
 "L10N:error executing template {{.Name}}": [
  {
   "Lang": "en",
   "Value": "error executing template {{.Name}}"
  },
  {
   "Lang": "de",
   "Value": "execute Template {{.Name}} fehlgeschlagen"
  }
 ],
 "L10N:error parsing '{{.Name}}'": [
  {
   "Lang": "en",
   "Value": "error parsing {{.Name}}"
  },
  {
   "Lang": "de",
   "Value": "parse Template {{.Name}} fehlgeschlagen"
  }
 ],
 "L10N:error unmarshaling 'l10nJSON'": [
  {
   "Lang": "en",
   "Value": "error unmarshaling l10nJSON"
  },
  {
   "Lang": "de",
   "Value": "unmarshali l10nJSON fehlgeschlagen"
  }
 ],
 "L10N:execute template {{.Name}} failed": [
  {
   "Lang": "en",
   "Value": "error executing template {{.Name}}"
  },
  {
   "Lang": "de",
   "Value": "execute Template {{.Name}} fehlgeschlagen"
  }
 ],
 "L10N:no entry for '{{.Name}}'": [
  {
   "Lang": "en",
   "Value": "no entry for '{{.Name}}'"
  },
  {
   "Lang": "de",
   "Value": "Fehlertext '{{.Name}}' inbekannt"
  }
 ],
 "L10N:no {{.Nam2}} translation for '{{.Name}}'": [
  {
   "Lang": "en",
   "Value": "no {{.Nam2}} translation for '{{.Name}}'"
  },
  {
   "Lang": "de",
   "Value": "keine {{.Nam2}}-Übersetzung für '{{.Name}}' verfügbar"
  }
 ],
 "L10N:open {{.Name}} failed": [
  {
   "Lang": "en",
   "Value": "open file {{.Name}} failed"
  },
  {
   "Lang": "de",
   "Value": "Öffnen der Datei {{.Name}} fehlgeschlagen"
  }
 ],
 "L10N:parse template {{.Name}} failed": [
  {
   "Lang": "en",
   "Value": "error parsing {{.Name}}"
  },
  {
   "Lang": "de",
   "Value": "parse Template {{.Name}} fehlgeschlagen"
  }
 ],
 "L10N:variable {{.Name}} not declared": [
  {
   "Lang": "en",
   "Value": "variable {{.Name}} not declared"
  },
  {
   "Lang": "de",
   "Value": "kein Variable {{.Name}} deklariert"
  }
 ],
 "L10N:write to file {{.Name}} failed": [
  {
   "Lang": "en",
   "Value": "write to file {{.Name}} failed"
  },
  {
   "Lang": "de",
   "Value": "Schreiben in Datei {{.Name}} fehlgeschlagen"
  }
 ],
 "L10N:wrong variable type {{.Type}} for '{{.Name}}', expected: {{.Nam2}}": [
  {
   "Lang": "en",
   "Value": "wrong variable type '{{.Type}}' for '{{.Name}}', expected: {{.Nam2}}"
  },
  {
   "Lang": "de",
   "Value": "falscher Typ '{{.Type}}' für Variable {{.Name}}, sollte {{.Nam2}} sein"
  }
 ],
 "L10N:{{.Nam2}} in {{.Name}} is missing": [
  {
   "Lang": "en",
   "Value": "{{.Nam2}} in {{.Name}} is missing"
  },
  {
   "Lang": "de",
   "Value": "{{.Nam2}} fehlt in {{.Name}}"
  }
 ]
}`

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
