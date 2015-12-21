// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

package message

import (
	"errors"
	"testing"
)

// TestPrepend checks with different input for Prepend().
func TestPrepend(t *testing.T) {
	fncname := "TestPrepend"

	tests := []struct {
		pre        string
		err        error
		out        bool
		outError   string
		outDetails string
	}{
		{ // nil
			"", nil,
			false, "", "",
		},
		{ // empty error
			"", errors.New(""),
			true, "", "",
		},
		{
			"Präfix", errors.New(""),
			true, "", "Präfix",
		},
		{ // error
			"", errors.New("Error"),
			true, "Error", "",
		},
		{
			"Präfix", errors.New("Error"),
			true, "Error", "Präfix",
		},
		{ // empty message
			"", New(""),
			true, "", "",
		},
		{
			"Präfix", New(""),
			true, "", "Präfix",
		},
		{ // message
			"", New("Message"),
			true, "Message", "",
		},
		{
			"Präfix", New("Message"),
			true, "Message", "Präfix",
		},
	}

	for i, v := range tests {

		var ok bool
		if v.err == nil {
			ok = Prepend(v.pre, nil)
			t.Log("Testfall:", i, "", v.err)
		} else {
			ok = Prepend(v.pre, &v.err)
			t.Log("Testfall:", i, "", v.err)
		}

		if ok != v.out {
			t.Errorf(fncname+"\nTestfall %3d: %t\nexpected:     %t", i, ok, v.out)
		}
		if ok {
			if (v.err).(*Message).Text() != v.outError {
				t.Errorf(
					fncname+"\n Testfall %3d: %s\nexpected:     %s",
					i, (v.err).(*Message).Text(), v.outError,
				)
			}
			if (v.err).(*Message).Details() != v.outDetails {
				t.Errorf(
					fncname+"\n Testfall %3d: %s\nexpected:     %s",
					i, (v.err).(*Message).Details(), v.outDetails,
				)
			}
		}
	}
}

// TestAppend checks with different input for Append().
func TestAppend(t *testing.T) {
	fncname := "TestAppend"

	tests := []struct {
		pre        string
		err        error
		out        bool
		outError   string
		outDetails string
	}{
		{ // nil
			"", nil,
			false, "", "",
		},
		{ // empty error
			"", errors.New(""),
			true, "", "",
		},
		{
			"Suffix", errors.New(""),
			true, "", "Suffix",
		},
		{ // error
			"", errors.New("Error"),
			true, "Error", "",
		},
		{
			"Suffix", errors.New("Error"),
			true, "Error", "Suffix",
		},
		{ // empty message
			"", New(""),
			true, "", "",
		},
		{
			"Suffix", New(""),
			true, "", "Suffix",
		},
		{ // message
			"", New("Message"),
			true, "Message", "",
		},
		{
			"Suffix", New("Message"),
			true, "Message", "Suffix",
		},
	}

	for i, v := range tests {

		var ok bool
		if v.err == nil {
			ok = Append(v.pre, nil)
			t.Log("Testfall:", i, "", v.err)
		} else {
			ok = Append(v.pre, &v.err)
			t.Log("Testfall:", i, "", v.err)
		}

		if ok != v.out {
			t.Errorf(fncname+"\nTestfall %3d: %t\nexpected:     %t", i, ok, v.out)
		}
		if ok {
			if (v.err).(*Message).Text() != v.outError {
				t.Errorf(
					fncname+"\n Testfall %3d: %s\nexpected:     %s",
					i, (v.err).(*Message).Text(), v.outError,
				)
			}
			if (v.err).(*Message).Details() != v.outDetails {
				t.Errorf(
					fncname+"\n Testfall %3d: %s\nexpected:     %s",
					i, (v.err).(*Message).Details(), v.outDetails,
				)
			}
		}
	}
}

// TestNilError checks with nil input to Append() and Prepend().
func TestNilError(t *testing.T) {
	fncname := "TestNilError"

	err := func() error { return nil }()
	var ok bool

	ok = Append("doesntmatter", &err)
	if ok {
		t.Error(fncname + ": Append to nil error returned true; expected: false")
	}

	ok = Prepend("doesntmatter", &err)
	if ok {
		t.Error(fncname + ": Prepend to nil error returned true; expected: false")
	}
}

var msgText = "test message #{{.Num}} from {{.Name}}"
var num = 4711
var name = "Nullachtfuffzehn"
var numVars = 2

// TestCascade does several checks including Append() and Prepend() to existing details.
func TestCascade(t *testing.T) {

	err := cascade1()
	if err == nil {
		t.Fatal("error expected")
	}

	m, ok := err.(*Message)

	if !ok {
		t.Fatal("message expected")
	}
	if m.text != msgText {
		t.Fatal("message.txt:", m.text, " expected:", msgText)
	}
	if len(m.vars) != numVars {
		t.Fatal("number of variables:", len(m.vars), " expected:", numVars)
	}
	for i, v := range m.vars {
		if v.Name == "Num" && v.Value.(int) != num {
			t.Fatal(
				"\nmessage.vars[", i, "]:", v.Name, "=", v.Value.(int),
				"\nexpected:         ", "Num", "=", num,
			)
		}
		if v.Name == "Name" && v.Value.(string) != name {
			t.Fatal(
				"\nmessage.vars[", i, "]:", v.Name, "=", v.Value.(string),
				"\nexpected:         ", "Name", "=", name,
			)
		}
	}
	if m.details != "cascade1:cascade2 suffix" {
		t.Fatal("message.details:", m.details, " expected:", "cascade1:cascade2 suffix")
	}

	t.Log(m.Text())
	t.Log(m.Vars())
	t.Log(m.Details())
}

func cascade1() error {
	fncname := "cascade1"

	err := cascade2()
	if Prepend(fncname+":", &err) {
		Append(" suffix", &err)
		return err
	}
	return nil
}

func cascade2() error {
	fncname := "cascade2"

	msg := New(msgText)
	msg.SetDetails(fncname)
	msg.AddVar("Num", num)
	msg.AddVar("Name", name)
	return msg
}
