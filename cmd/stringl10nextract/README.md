# stringl10nextract
A programmer's tool to extract string literals from source files.

This command generates output data that can be used as input to the
string localization tool stringl10n. Data must be edited.

### Limitations
There is not much intelligence implemented. A lot of useless
literals get extracted. Remove them by editing.

### Installation
Provided that your Go environment is ready, i.e. $GOPATH is set, you just do:

`$ go get github.com/hwheinzen/stringl10n/cmd/stringl10nextract`

### Usage
In your project directory:

`$ find . | stringl10nextract > example.json`

or:

`$ stringl10nextract -dir=. -o=example.json`

If you control the source you could prepare all strings you want translated with
let's say a prefix 'I18N:'. Then try this one:

`$ find . | stringl10nextract -regexp=^I18N: > example.json`

You'll get a file `example.json` that looks something like:

```
# ... some instructions here ...
{
	 "Copyright": "2015 Itts Mee"
	,"Package":   "example"
	,"GenFile":   "stringl10n_generated.go"

	,"Vars": [
			 {"Name": "", "Type": ""}
			,{"Name": "", "Type": ""}
	]

	,"Funcs": [
			 {"Name": "", "Function": "", "Path": ""}
			,{"Name": "", "Function": ""}
	]

	,"Texts":	{
		 "I18N:programmer's words 1": [
			 {"Lang": "", "Value": ""}

		 # somesource.go:112:22

		]
		,"I18N:programmer's words 2": [
			 {"Lang": "", "Value": ""}

		 # somesource.go:128:22

		]
	}
}
```

### TODO
- think about workflow
