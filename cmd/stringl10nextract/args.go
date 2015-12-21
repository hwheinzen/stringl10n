// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
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
)

// arguments
// ---------
// -dir		<source directory>	(default stdin)
// -o		<output file>		(default stdout)
// -lang	programmers language code	(default none)
// -deep	<include subdirs>	(default false)
// -min		<min. length>		(default 3)
// -max		<max. length>		(default 200)
// -deep	<include subdirs>	(default false)
// -inclist	<only if string contains one of the words>
// -regexp	<only if string conforms to regexp>
// -help
//
// example
// -------
// $ stringl10nextract -dir=$GOPATH/hawe/bgzdb

// global variables
var (
	argDir      string
	argOut      string
	argLang     string
	argKeywords []string // from inclist
	argRegexp   string
	argDeep     bool
	argMin      int
	argMax      int

	gRgx *regexp.Regexp
)

func args() {

	var inclist string
	var help bool

	flag.StringVar(&argDir, "dir", "", "directory instead of stdin")
	flag.StringVar(&argOut, "o", "", "output file instead of stdout")
	flag.StringVar(&argLang, "lang", "", "programmers language, e.g. en")
	flag.StringVar(&inclist, "inclist", "", "string must contain these")
	flag.StringVar(&argRegexp, "regexp", "", "string must conforms")
	flag.BoolVar(&argDeep, "deep", false, "dive into sub-directories")
	flag.IntVar(&argMin, "min", 3, "string contains at least min runes")
	flag.IntVar(&argMax, "max", 200, "string contains at most max runes")
	flag.BoolVar(&help, "help", false, "usage")

	flag.Usage = argsUsage
	flag.Parse()

	if help {
		argsUsage()
		os.Exit(0)
	}

	if inclist != "" {
		argKeywords = strings.Split(inclist, " ")
	}

	var err error
	if argRegexp != "" {
		gRgx, err = regexp.Compile(argRegexp)
		if err != nil {
			log.Fatalln(pgmname+":", err)
		}
	}

	return
}

func argsUsage() {
	fmt.Fprintln(os.Stderr, "usage:", pgmname, "[OPTION]...")
	flag.PrintDefaults()
}
