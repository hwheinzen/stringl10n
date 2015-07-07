# stringl10nextract
A programmer's tool to extract string literals from source files.

This command generates output data that can be used as input to the
string localization tool stringl10n. Data must be edited.

### Limitations

There is not much intelligence implemented. A lot of useless
literals get extracted. Remove them by editing.

### Installing

Provided that your Go environment is ready, i.e. $GOPATH is set, you need to:

`$ go get github.com/hwheinzen/stringl10n/cmd/stringl10nextract`

### Usage
In your project directory:

`$ find . | stringl10nextract > example.txt`

or:

`$ stringl10nextract -root=. -o=example.txt`

and you'll get a file `example.txt` that looks something like:

```
// ... some instructions here ...
{
	"Copyright": "2015 Itts Mee"
	,"Package":  "example"
	,"GenFile":  "stringl10n_generated.go"

	,"Vars":     [
			{ "Name": "", "Type": ""}
			,{"Name": "", "Type": ""}
		]

	,"Funcs": [
			{ "Name": "", "Function": "", "Path": ""}
			,{"Name": "", "Function": "", "Path": ""}
	]

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
