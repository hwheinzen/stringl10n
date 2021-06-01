// Copyright 2020-21 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
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

func args(buildtime string) (jsonFile string) {

	var version bool
	flag.BoolVar(&version, "version", false, "(if built with -ldflags \"-X main.buildtime '```date -Iseconds`'\")") // ``` seem to be necessary for PrintDefaults()

	var help bool
	flag.BoolVar(&help, "help", false, "Usage information")

	flag.StringVar(&jsonFile, "json", "", "input file name")

	flag.Parse()

	if help {
		flag.Usage()
		os.Exit(0)
	}

	if version {
		if buildtime == "" {
			fmt.Println(pgm+": unknown version")
		} else {
			fmt.Println(pgm+": version of "+ buildtime)
		}
		os.Exit(0)
	}

	if jsonFile == "" {
		fmt.Println(pgm+": -json argument missing")
		flag.Usage()
		os.Exit(2)
	}

	return jsonFile
}
