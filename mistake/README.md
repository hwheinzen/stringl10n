# mistake
A useful error type 

Package `mistake` provides an error type that keeps
fix and variable parts of a text message apart.
It is meant to be used together with the `l10n` tool.

### Motivation
To translate error messages to any users language - even in a 
multi-user/multi-language environment - it is necessary either
to transport the language code to every place where an error message
is created, or to transport the variable parts of an error message
separated from the fix error string and do the translation 
at a place higher up in the logical hierarchy and substitute the
variables there.

### Properties
`Err` has two components:
- `Fix` is a string
- `Var` is a slice of name/value pairs

`Fix` is meant to hold a message text with optional placeholders
embedded. If you use the `l10n` tool, you will probably make
placeholders text/template expressions.

`Var` hold name/value pairs where the names correspond to the
text/template names in `Fix`.

The error type `Err` satifies the standard Go `error` interface,
meaning: it has an `Error` method which returns `Fix`.

The error type `Err` has an additional method `Vars` which returns `Var`.

### Install
Provided that your Go environment is ready, just do:

`$ go get github.com/hwheinzen/stringl10n/mistake`

### Usage
The `mistake` package is sufficiently small: you can use a dot import:

`import . "github.com/hwheinzen/stringl10n/mistake"`

To create an `Err` do:

```
	e := Err{
		Fix: "...text ... {{.Name}} ...text ...",
		Var: []struct {
			Name  string
			Value interface{}
		}{
			{"Name", variable},
		},
	}
```
