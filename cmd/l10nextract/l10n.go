// Copyright 2020 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// l10n.json must be READY.
//
//go:generate l10n -json=l10n.json

package main

import (
	"log"
)

func translate(in error, lang string) (out error) {
	if in == nil {
		return nil
	}

	out, err := L10nLocalizeError(in, lang)
	if err == nil {
		return out
	}

	if err != nil {
		log.Println(err)
	}
	return in
}
