// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// Arguments
// ---------
// -root	<source directory>	(default ".")
// -o		<output file>		(default stdout)
// -deep	<include subdirs>	(default false)
// -min		<min. length>		(default 3)
// -max		<max. length>		(default 200)
// -deep	<include subdirs>	(default false)
// -inclist	<only if string contains one of the words>
// -help
//
// Example
// -------
// $ stringl10nextract -root=$GOPATH/hawe/bgzdb

var (
	argRoot string
	argOut string
	argKeywords []string
	argDeep bool
	argMin int
	argMax int
)

func args() {

	var inclist string
	var help bool

	flag.StringVar(&argRoot, "root", ".", "source directory")
	flag.StringVar(&argOut, "o", "", "output file instead of stdout")
	flag.StringVar(&inclist, "inclist", "", "only string which contain these")
	flag.BoolVar(&argDeep, "deep", false, "dive into sub-directories")
	flag.IntVar(&argMin, "min", 3, "string contains at least min runes")
	flag.IntVar(&argMax, "max", 200, "string contains at most max runes")
	flag.BoolVar(&help, "help", false, "usage information")

	flag.Usage = argsUsage
	flag.Parse()

	if help {
		argsUsage()
		os.Exit(0)
	}

	if inclist != "" {
		argKeywords = strings.Split(inclist, " ")
	}

	return
}

func argsUsage() {
	fmt.Fprintln(os.Stderr, "Usage:", pgmname, "[OPTION]...")
	flag.PrintDefaults()
}
