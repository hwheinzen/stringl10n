// Copyright 2020 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	. "github.com/hwheinzen/stringl10n/mistake"
)

// arguments
// ---------
// -input-dir		<source directory>		(default stdin)
// -i  				<curremt JSON file>
// -o				<output file>			(default stdout)
// -lang			programmers language code	(default "en")
// -regexp			<only if string conforms to regexp>
// -inclist			<only if string contains one of the words>
// -deep			<include subdirs>		(default false)
// -min				<min string length>		(default 3)
// -max				<max string length>		(default 200)
// -deep			<include subdirs>		(default false)
// -help
//
// example
// -------
// $ l10nextract -input-dir=$GOPATH/hawe/bgzdb

// All arguments are global.
var (
	argDir      string
	argOut      string
	argJSON     string
	argLang     string
	argKeywords []string // from inclist
	argDeep     bool
	argMin      int
	argMax      int

	argRgx *regexp.Regexp
)

func args(buildtime string) {

	var version bool
	flag.BoolVar(&version, "version", false, "(if built with -ldflags \"-X main.buildtime='```date -Iseconds`'\"") // ``` seem to be necessary for PrintDefaults()

	var help bool
	flag.BoolVar(&help, "help", false, "Usage information")

	flag.StringVar(&argDir, "input-dir", "", "directory instead of stdin")
	flag.StringVar(&argJSON, "i", "", "current JSON file")
	flag.StringVar(&argOut, "o", "", "output file instead of stdout")
	flag.StringVar(&argLang, "lang", "en", "programmers language")

	var update string
	flag.StringVar(&update, "update", "", "update JSON file; same as: -i=X -o=Y where X == Y")

	var inclist string
	flag.StringVar(&inclist, "inclist", "", "string must contain these")

	var rgx string
	flag.StringVar(&rgx, "regexp", "", "string must conform to")

	flag.BoolVar(&argDeep, "deep", false, "dive into sub-directories")
	flag.IntVar(&argMin, "min", 3, "string contains at least min runes")
	flag.IntVar(&argMax, "max", 200, "string contains at most max runes")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if version {
		if buildtime == "" {
			inf := Err{
				Fix: "L10NEXTRACT:{{.Name}}:unknown version",
				Var: []struct {
					Name  string
					Value interface{}
				}{
					{"Name", pgm},
				},
			}
			fmt.Println(translate(inf, argLang))
		} else {
			inf := Err{
				Fix: "L10NEXTRACT:{{.Name}}:version of {{.Nam2}}",
				Var: []struct {
					Name  string
					Value interface{}
				}{
					{"Name", pgm},
					{"Nam2", buildtime},
				},
			}
			fmt.Println(translate(inf, argLang))
		}
		os.Exit(0)
	}

	if update != "" { // -i == -o == -update
		argJSON = update
		argOut = update
	}

	if inclist != "" {
		argKeywords = strings.Split(inclist, " ")
	}

	//var err error
	if rgx != "" {
		var err error
		argRgx, err = regexp.Compile(rgx)
		if err != nil {
			log.Fatalln(pgm+":", err)
		}
	}

	return
}
