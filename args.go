// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
)

// Arguments
// ---------
// -json	<strings file name>	(compulsory)
// -help
//
// Example
// -------
// $ stringl10n -json=example.json
//
// or
// --
// //go:generate stringl10n -json=example.json

func args() (argJson string) {

	var argHelp bool

	flag.StringVar(&argJson, "json", "", "Input file name")
	flag.BoolVar(&argHelp, "help", false, "Help")

	flag.Usage = argsUsage
	flag.Parse()

	if argHelp {
		argsUsage()
		os.Exit(0)
	}

	if argJson == "" {
		fmt.Println(pgmname+":", "-json missing")
		argsUsage()
		os.Exit(2)
	}

	return
}

func argsUsage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION] ...\n", pgmname)
	flag.PrintDefaults()
}
