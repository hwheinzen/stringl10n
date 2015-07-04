// example.go starts the creation of stringl10n_generated.go.
//
//go:generate stringl10n -json=example.json

package main

import (
	"fmt"

	"github.com/hwheinzen/mist" // extended error
)

var ganzzahl int = 42
var gleitkomma float64 = 0.815
type Struktur struct {
	Teil1 string
	Teil2 string
}
var struk = Struktur{"Teil1", "Teil2"}
	

func main() {
	xerr := makeXerror()

	fmt.Println("  :", xerr.Error())

	trans := l10nTranslate(xerr.Error(), "en")
	subst := l10nSubstitute(trans, xerr)
	fmt.Println("en:", subst)

	trans = l10nTranslate(xerr.Error(), "de")
	subst = l10nSubstitute(trans, xerr)
	fmt.Println("de:", subst)

	trans = l10nTranslate(xerr.Error(), "ex")
	subst = l10nSubstitute(trans, xerr)
	fmt.Println("ex:", subst)

	trans = l10nTranslate(xerr.Error(), "XX")
}

func makeXerror() (xerr mist.XError) {
	xerr = mist.New(
		"1: {{printf \"%d\" .I1}} 2: {{printf \"%f\" .Fl1}}",
		"ignored",
	).(mist.XError) // Typzusicherung, weil mist.New nur "error" zurückgibt
	xerr.AddVar("I1", ganzzahl)
	xerr.AddVar("Fl1", gleitkomma)
	xerr.AddVar("S1", struk)
	return
}
