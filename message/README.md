# message
A message type for Go -- OBSOLETE -- use package `mistake` instead

Package message provides a message type that keeps
fix and variable parts of a text message apart.
(This enables "last-minute" localization.)

### Motivation
To translate error messages to any users language - even in a 
multi-user/multi-language environment - it is necessary ... either
to transport the language code to every place an error message
is created ... or to transport the variable parts of an error message
separated from the fix error string and do the translation 
at one place high up in the logical hierarchy and substitute the
variables then.

### Properties

The message type satifies the standard Go error interface.

A message has three components:
- a fix part which is a string (function New, methods Text Error)
- any number of text/value pairs (methods AddVar, Vars)
- a variable part which is a string (methods SetDetails, Details)

The fix part of the message is meant to hold a message text
which can be translated later with the help of l10nTranslate
(generated by stringl10n into YOUR package).

The fix part may contain placeholders in the form of
text/template expressions
```... {{.StringVar}} ... {{printf \"%d\" .IntVar}} ...```
which can be substituted (after the translation) by the messages'
variables with the help of l10nSubstitute
(generated by stringl10n into YOUR package).

There are two convenience functions for adding a prefix or a suffix
to the details of a message; they take any error and make it a
message wrapped in the error interface.

The details part can serve as an error trace if you prefix it
with the function name every time a function returns an error:

```
    err := otherfunction()
    if message.Prepend(thisfunctionname+":", &err) {
        return err
    }

```

### Installation
Provided that your Go environment is ready, i.e. $GOPATH is set, you just do:

`$ go get github.com/hwheinzen/stringl10n/message`

### Usage
See `message_test.go`.
