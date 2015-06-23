# stringl10n
String localization tools for Go.

Two commands are available:

One: stringl10n generates code to be included in a project. 
A function l10nTranslate (alias t) translates strings.
A function l10nSubstitute resolves text/template expressions in a
string and substitutes variables. It needs a JSON input file.

Two: stringl10nextract extracts string literals from a go project
and puts them into a JSON-encoded file. After editing this can
be input for stringl10n.
