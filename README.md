# stringl10n
String localization tools for Go.

Here we have one package and two commands for generating code:

1 . Package `message` provides a `Message` type with fix and variable parts;
it can be used as a kind of extended error type due to its `Error` method. 

A message can be created by function `New`.
Functions `Append` and `Prepend` can take an error interface and
transform the underlying concrete type to `Message`.

2 . Command `stringl10n` generates code to be included in a project. 
Function `l10nTrans` translates strings.
Function `l10nSubst` resolves the text/template expressions of a
text string and substitutes variables. 
(stringl10n needs a JSON input file.)

A message can be translated by `l10nTrans` because it is a text string.
Placeholders (in text/template format) in the messages text string can
be substituted  by `l10nSubst` because the message provides the necessary
variables.

3 . Comand `stringl10nextract` prepares input for stringl10n.
It extracts string literals from a go project and puts them into
a JSON-encoded file. (You have to edit his file.)

### Installing
Provided that your Go environment is ready, i.e. $GOPATH is set, you just do:

`$ go get github.com/hwheinzen/stringl10n/...`
