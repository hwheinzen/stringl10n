// example.go starts the creation of stringl10n_generated.go.
//
//go:generate stringl10n -json=example.json

package main

import (
	"fmt"

	"github.com/hwheinzen/stringl10n/message"
)

var integer int = 42
var floatie float64 = 0.815
type Struct struct {
	Part1 string
	Part2 string
}
var structure = Struct{"Pürt1", "Pürt2"}
	

func main() {

	var err error

	err = makeErrorMessage()
	fmt.Println("  :", err.Error())

	var trans, subst string

	trans = l10nTrans(err.Error(), "en")
	subst = l10nSubst(trans, err.(*message.Message))
	fmt.Println("en:", subst)

	trans = l10nTrans(err.Error(), "de")
	subst = l10nSubst(trans, err.(*message.Message))
	fmt.Println("de:", subst)

	trans = l10nTrans(err.Error(), "xx")
	subst = l10nSubst(trans, err.(*message.Message))
	fmt.Println("xx:", subst)

	trans = l10nTrans(err.Error(), "unknown")
	fmt.Println("default:", trans)
}

func makeErrorMessage() error {
	msg := message.New("Int: {{printf \"%d\" .Int}} Float: {{printf \"%f\" .Flo}}")
	msg.AddVar("Int", integer)
	msg.AddVar("Flo", floatie)
	msg.AddVar("Str", structure) // language xx only
	return msg
}
