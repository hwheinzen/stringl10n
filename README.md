# stringl10n
A simple string localization tool for Go

This command generates Go code. It allows to initially choose a language for text strings at compile time, and to reset globally the language at run time.

## Limitations
There is neither variable substitution nor plural handling.

Because the language gets set globally stringl10n is not a solution for multi-user-multi-language services.

## Installing
Provided that your Go environment is ready, i.e. $GOPATH is set, you need to:

`$ go get github.com/hwheinzen/stringl10n`

## Usage
Scan your projects code base for string literals, replace them by unique variable names, and map these names and strings in a JSON file:

```
{
	"Copyright": "<year> <copyright owner>"
	,"Package":  "example"
	,"GenFile":  "stringl10n_generated.go"
	,"Default":  "en"
	,"Start":    "de"
	,"Text":	[
		{
			"Name":     "localizedStringVariable1"
			,"Locs": [
				 {"ID": "en", "Value": "english words 1"}
			]
		}
		,{
			"Name":     "localizedStringVariable2"
			,"Locs": [
				 {"ID": "en", "Value": "english words 2"}
			]
		}
	]
}
```

Add translations:

```
			"Name":     "localizedStringVariable1"
			,"Locs": [
				 {"ID": "en", "Value": "english words 1"}
				,{"ID": "de", "Value": "deutsche Wörter 1"}
				,{"ID": "fr", "Value": "mots françaises 1"}
				...
```

Add one line to your go code:

`//go:generate stringl10n -json=example.json`

Then run `go generate` before building your package or command.

## TODO
- Mutex
