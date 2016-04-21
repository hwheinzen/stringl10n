# stringl10n
A simple string localization tool for Go.

This command generates two source files that can be included in a Go project:
- one that provides functions to translate strings and to substitute variables
- one that provides a unit test

### Limitations
If you are looking for advanced features like plural handling et cetera
better have a look at [golang.org/x/text/...](https://golang.org/x/text).

### Installing
Provided that your Go environment is ready, i.e. $GOPATH is set, you just do:

`$ go get github.com/hwheinzen/stringl10n/cmd/stringl10n`

### Usage
Scan your projects code base for string literals.
(The tool stringl10nextract helps you.)

Map these strings with their translations inside a JSON file (e.g. example.json):

```
{
...
		,"programmer's words": [
			 {"Lang": "en", "Value": "english words"}
			,{"Lang": "de", "Value": "deutsche Wörter"}
			,{"Lang": "fr", "Value": "mots françaises"}
		]
...
}
```

Add one line to your go code:

`//go:generate stringl10n -json=example.json`

Run `go generate`.

You can now use these functions in your project:

`func l10nTrans(key, lang string) (value string)`

`func l10nSubst(tmpl string, vars Varser) (out string)`

l10nTrans returns the matching string in the requested language.
l10nSubst returns the string where placeholders are replaced by the matching variables.

Varser is an interface type:
```
type Varser interface {
    Vars() []struct{ // Vars returns Name-Value-Pairs
        Name string
        Value interface{}
    }
}
```
(message.Message satisfies this Interface.)

You can let stringl10n register fuctions to the template.FuncMap, so any kind of formatting is possible.
More info is in the API docs.

### TODO
Generate more tests.
