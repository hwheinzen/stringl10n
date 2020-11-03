// l10n.json must be READY.
//
//go:generate l10n -json=l10n.json

package main

import (
	"log"
	//	"otherpackage"
)

func translate(in error, lang string) (out error) {
	if in == nil {
		return nil
	}

	out, err := L10nLocalizeError(in, lang)
	if err == nil {
		return out
	}

	// If you use other selfmade packages that support l10n:
	// - import these and ...
	// - insert following code per package:
	//
	// 	out, err := otherpackage.L10nLocalizeError(in, lang)
	// 	if err == nil {
	// 		return out
	// 	}

	if err != nil {
		log.Println(err)
	}
	return in
}
