# stringl10n
String localization tools for Go.

Two commands are available:

1: stringl10n generates code to be included in a project. 
A function l10nTranslate (alias t) translates strings.
A function l10nSubstitute resolves text/template expressions in a
string and substitutes variables. Stringl10n needs a JSON input
file.

2: stringl10nextract extracts string literals from a go project
and puts them into a JSON-encoded file. After editing this can
be input for stringl10n.

### Installing
Provided that your Go environment is ready, i.e. $GOPATH is set, you need to:

`$ go get github.com/hwheinzen/stringl10n/...`
