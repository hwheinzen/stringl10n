// Copyright 2020 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package mistake

// Err is an error type that carries variables.
// Err is supposed to be used together with l10n.
//
// Fix is the error string; it may contain text/template
//  variables like {{.Myvar}} .
// Var then should contain pair structures with e.g.
//  Name="Myvar" and Value=<the corresponding Value>.
type Err struct {
	Fix string
	Var []struct {
		Name  string
		Value interface{}
	}
}

// Error returns the fix error string;
// this satisfies the error interface.
func (m Err) Error() string {
	return m.Fix
}

// Vars returns the Name/Value pairs attached to the error.
func (m Err) Vars() []struct {
	Name  string
	Value interface{}
} {
	return m.Var
}
