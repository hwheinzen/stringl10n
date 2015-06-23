# stringl10n
A simple string localization tool for Go.

This command generates two source files that can be included in a Go project:
- one that provides a function that translates strings
- one that provides a unit test

### Limitations
There is neither variable substitution nor plural handling.

### Installing
Provided that your Go environment is ready, i.e. $GOPATH is set, you need to:

`$ go get github.com/hwheinzen/stringl10n`

### Usage
Scan your projects code base for string literals, and map these strings with their translations inside a JSON file (e.g. example.json):

```
{
	"Copyright": "2015 Itts Mee"
	,"Package":  "example"
	,"GenFile":  "stringl10n_generated.go"

	,"Texts":	{
		"programmer's words 1": [
			{ "Lang": "de", "Value": "deutsche Wörter 1"}
			,{"Lang": "en", "Value": "english words 1"}
			,{"Lang": "fr", "Value": "mots françaises 1"}
		]
		,"programmer's words 2": [
			{ "Lang": "de", "Value": "deutsche Wörter 2"}
			,{"Lang": "en", "Value": "english words 2"}
			,{"Lang": "fr", "Value": "mots françaises 2"}
		]
	}
}
```

Add one line to your go code:

`//go:generate stringl10n -json=example.json`

Run `go generate`.

You can now use function t, for example:

`	err := errors(t("programmer's words 1", "de"))`

After translating you can substitute text/template expressions inside
the string with matching variables. More info is in the API docs.