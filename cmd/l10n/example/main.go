// example: How to use l10n together with mistake.Err

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	. "github.com/hwheinzen/stringl10n/mistake"
)

const pgm = "example"

func main() {
	lang := args()

	var err error

	err = getSimpleError()
	if err != nil {
		err = translate(err, lang) // ******** l10n ********
		log.Println(pgm + ":" + err.Error())
	}

	err = getComplexError()
	if err != nil {
		err = translate(err, lang) // ******** l10n ********
		log.Println(pgm + ":" + err.Error())
	}
}

func args() (lang string) {
	var help bool
	flag.BoolVar(&help, "help", false, "usage information")

	flag.StringVar(&lang, "lang", "en", "language of error messages")

	flag.Parse()
	if help {
		flag.Usage()
		os.Exit(0)
	}

	return lang
}

func getSimpleError() error {
	fnc := "getSimpleError"

	e := Err{Fix: "EXAMPLE:simple error"}
	return fmt.Errorf(fnc+":%w", e) // wrap
}

func getComplexError() error {
	fnc := "getComplexError"

	err := errors.New("some error message from std library function")
	add := "variable information"

	e := Err{
		Fix: "EXAMPLE:complex error with '{{.Info}}'",
		Var: []struct {
			Name  string
			Value interface{}
		}{
			{"Info", add},
		},
	}
	return fmt.Errorf(fnc+":%w:"+err.Error(), e) // wrap
}
