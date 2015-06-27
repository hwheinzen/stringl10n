# stringl10n
A simple string localization tool for Go.

This command generates two source files that can be included in a Go project:
- one that provides functions for translating strings and for variable substitution
- one that provides a unit test

Errors during translation/substitution are logged.

### Limitations
There is no plural handling.

### Installing
Provided that your Go environment is ready, i.e. $GOPATH is set, you need to:

`$ go get github.com/hwheinzen/stringl10n/cmd/stringl10n`

### Usage
Scan your projects code base for string literals, and map these strings with their translations inside a JSON file (e.g. example.json):

```
{
...
		"programmer's words": [
			{ "Lang": "de", "Value": "deutsche Wörter"}
			,{"Lang": "en", "Value": "english words"}
			,{"Lang": "fr", "Value": "mots françaises"}
		]
...
}
```

Add one line to your go code:

`//go:generate stringl10n -json=example.json`

Run `go generate`.

You can now use function t, for example:

`	err := errors(t("programmer's words 1", "de"))`

After translating you can substitute text/template expressions in
the string with matching variables. 
You can let stringl10n register fuctions to the template.FuncMap,
so any kind of formatting is possible.
More info is in the API docs.