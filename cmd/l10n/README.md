# l10n
A simple localization tool -- 
particularly suited to errors of type `mistake.Err`

This command is a code generator. 
It takes a JSON file containing translations et cetera
and returns a Go source file containing functions that can be used
to translate texts, to substitute `text/template` expressions, and to
localize errors of type `mistake.Err`.

### Limitations
You have to handle advanced features like plural handling and number formatting yourself: use `text/template` with functions.

### Install
Provided that your Go environment is ready, just do:

`$ go get github.com/hwheinzen/stringl10n/cmd/l10n`

### Usage
Scan your projects code base for string literals; the tool `l10nextract` helps you.

Map these strings with translations inside a JSON file (e.g. l10n.json):

```
{
...
		,"programmer's words": [
			{"Lang": "en", "Value": "english words"},
			{"Lang": "de", "Value": "deutsche Wörter"},
			{"Lang": "fr", "Value": "mots françaises"}
		]
...
}
```

Add one line to your go code:

`//go:generate l10n -json=l10n.json`

Run `go generate`.

You can now use these functions in your project:

```
func L10nTranslate(in, lang string) (out string, err error)

func L10nSubstitute(in string, vars []struct {
	Name  string
	Value interface{}
}) (out string, err error)

func L10nLocalizeError(in error, lang string) (out, err error)
```

`L10nTranslate` returns the matching string in the requested language.

`L10nSubstitute` returns the string with `text/template` expressions replaced using the matching variables and perhaps functions.

`L10nLocalizeError` operates on a `mistake.Err` and combines the two former functions.

### Example
See code in directory `example`. There run:

```
$ go generate
$ go build
$ ./example
...
$ ./example -lang=de
...
```
