// Copyright 2020 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"

	. "github.com/hwheinzen/stringl10n/mistake"
)

// arguments
// ---------
// -json	<strings file name>	(MUST)
// -lang	<default: en>
// -version
// -help
//
// example
// -------
// $ l10n -json=l10n.json
//
// or
// --
// //go:generate l10n -json=l10n.json

func args(buildtime string) (jsonFile, lang string) {

	var version bool
	flag.BoolVar(&version, "version", false, "(if built with -ldflags \"-X main.buildtime '`date -Iseconds`'\"")

	var help bool
	flag.BoolVar(&help, "help", false, "Usage information")

	flag.StringVar(&jsonFile, "json", "", "input file name")

	flag.StringVar(&lang, "lang", "en", "language of error messages")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if version {
		if buildtime == "" {
			inf := Err{
				Fix: "L10N:{{.Name}}:unknown version",
				Var: []struct {
					Name  string
					Value interface{}
				}{
					{"Name", pgm},
				},
			}
			fmt.Println(translate(inf, lang))
		} else {
			inf := Err{
				Fix: "L10N:{{.Name}}:version of {{.Nam2}}",
				Var: []struct {
					Name  string
					Value interface{}
				}{
					{"Name", pgm},
					{"Nam2", buildtime},
				},
			}
			fmt.Println(translate(inf, lang))
		}
		os.Exit(0)
	}

	if jsonFile == "" {
		err := Err{
			Fix: "L10N:{{.Name}}:{{.Nam2}} argument missing",
			Var: []struct {
				Name  string
				Value interface{}
			}{
				{"Name", pgm},
				{"Nam2", "-json"},
			},
		}
		fmt.Fprintln(os.Stderr, translate(err, lang))
		flag.Usage()
		os.Exit(2)
	}

	return jsonFile, lang
}
