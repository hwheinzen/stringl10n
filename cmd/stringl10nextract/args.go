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
// -root	<source directory>	(default ".")
// -o		<output>			(default stdout)
// -help
//
// Example
// -------
// $ stringl10nextract -root=$GOPATH/hawe/bgzdb

func args() (root, out string) {

	var help bool

	flag.StringVar(&root, "root", ".", "source directory")
	flag.StringVar(&out, "o", "", "output file instead of stdout")
	flag.BoolVar(&help, "help", false, "usage information")

	flag.Usage = argsUsage
	flag.Parse()

	if help {
		argsUsage()
		os.Exit(0)
	}

	return
}

func argsUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", pgmname, "[OPTION]...")
	flag.PrintDefaults()
}
