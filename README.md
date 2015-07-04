# stringl10n
String localization tools for Go.

Two commands are available:

1: stringl10n generates code to be included in a project. 
Function l10nTranslate translates strings.
Function l10nSubstitute resolves text/template expressions in a
string and substitutes variables. 
stringl10n needs a JSON input file.

2: stringl10nextract extracts string literals from a go project
and puts them into a JSON-encoded file. 
After editing can serve as input for stringl10n.

### Installing
Provided that your Go environment is ready, i.e. $GOPATH is set, you need to:

`$ go get github.com/hwheinzen/stringl10n/...`
