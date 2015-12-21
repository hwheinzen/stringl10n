// Copyright 2015 Hans-Werner Heinzen. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

// Package message provides a message type that keeps
// fix and variable parts of a text message apart.
// (This enables "last-minute" localization.)
package message

// Message holds the fix and variable parts of a message.
type Message struct {
	text    string
	details string
	vars    []struct {
		Name  string
		Value interface{}
	}
}

// New returns the address of a new message with the fix part set to text;
// text may contain placeholders like text/template expressions.
func New(text string) *Message {
	return &Message{text: text}
}

// SetDetails sets the details part of the message
// which is thus variable.
func (m *Message) SetDetails(d string) {
	m.details = d
}

// AddVar adds a new Name/Value pair to the message which is meant
// to later on substitute a placeholder in the fix part.
func (m *Message) AddVar(name string, value interface{}) {
	m.vars = append(
		m.vars,
		struct {
			Name  string
			Value interface{}
		}{
			Name:  name,
			Value: value,
		},
	)
}

// Text returns the fix part of the message.
func (m *Message) Text() string {
	return m.text
}

// Error satisfies the built-in error interface,
// which is its main purpose.
// It also returns the fix part of the message.
func (m *Message) Error() string {
	return m.text
}

// Details returns the details string.
func (m *Message) Details() string {
	return m.details
}

// Vars returns the Name/Value pairs attached to the message.
func (m *Message) Vars() []struct {
	Name  string
	Value interface{}
} {
	return m.vars
}

// Prepend adds a prefix to the details of a message and
// returns true if an error is given, else it returns false.
func Prepend(pre string, errp *error) bool {
	if errp == nil { // no error
		return false
	}
	if *errp == nil { // errors can be nil
		return false
	}
	msg, ok := (*errp).(*Message)
	if ok {
		// already a Message
		(*errp).(*Message).SetDetails(pre + msg.Details())
	} else {
		// create a new Message
		msg = New((*errp).Error())
		msg.SetDetails(pre)
		var err error = msg
		*errp = err // this is important (*errp) !
	}
	return true
}

// Append adds a suffix to the details of a message and
// returns true if an error is given, else it returns false.
func Append(suf string, errp *error) bool {
	if errp == nil { // no error
		return false
	}
	if *errp == nil { // errors can be nil
		return false
	}
	msg, ok := (*errp).(*Message)
	if ok {
		// already a Message
		(*errp).(*Message).SetDetails(msg.Details() + suf)
	} else {
		// create a new Message
		msg = New((*errp).Error())
		msg.SetDetails(suf)
		var err error = msg
		*errp = err // this is important (*errp) !
	}
	return true
}
