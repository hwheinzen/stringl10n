// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
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
// -help
//
// example
// -------
// $ stringl10n -json=example.json
//
// or
// --
// //go:generate stringl10n -json=example.json

func args() (argJson string) {

	var argHelp bool

	flag.StringVar(&argJson, "json", "", "Input file name")
	flag.BoolVar(&argHelp, "help", false, "Usage information")

	flag.Usage = argsUsage
	flag.Parse()

	if argHelp {
		argsUsage()
		os.Exit(0)
	}

	if argJson == "" {
		fmt.Fprintln(os.Stderr, pgmname+":", "flag is missing:", "-json")
		argsUsage()
		os.Exit(2)
	}

	return
}

func argsUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", pgmname, "[OPTION]...")
	flag.PrintDefaults()
}
