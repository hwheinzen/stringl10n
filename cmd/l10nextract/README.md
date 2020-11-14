# l10nextract
A programmer's tool to extract string literals from source files.

Command `l10nextract` generates JSON output (or updates a JSON file) that can be used as input to the string localization tool `l10n`. JSON data has to be edited by adding translations, variables for `text/template` expressions, and eventually functions that can be used in `text/template` expressions.

Parameter `-regexp` helps to limit extraction. Since the primary use of `l10nextract` and `l10n` is the localization of error messages, using a searchable prefix to your error messages is a good idea.

### Installation
Provided that your Go environment is ready, you run:

`$ go get github.com/hwheinzen/l10n/cmd/l10nextract`

### Usage
Initally run in your project directory:

`$ find . | l10nextract -regexp=^YOURPREFIX > l10n.json`

You'll get a file `l10n.json` that looks something like:

```
# ... some #-comments here ...
{
	"Copyright": "2015 Itts Mee",
	"Package":   "example",
	"ErrorType": "Err",
	"ErrorPref": "",
	"GenFile":   "l10n_generated.go",
	"Vars": [
			{"Name": "VarName", "Type": ""}
	],
	"Funcs": [],
	"Texts": {
		"YOURPREFIX:programmer's words 1 {{.VarName}}": [
			 {"Lang": "", "Value": ""}
		# somesource.go:112:22
		],
		"YOURPREFIX:programmer's words 2": [
			 {"Lang": "", "Value": ""}
		# somesource.go:128:22
		]
	}
}
```

Later, that is after using `l10n` to generate additional go code (that will contain some more error message strings), run:

`$ find . | l10nextract -regexp=^YOURPREFIX -update=l10n.json`

(You have to edit tje JSON file. Find further information with tool `l10n`.)
