// parse.go
// Parse Pyylisp or Lisp-formatted code into an abstract syntax tree.

package pyylisp

import (
	"fmt"
	"io"
	"strings"
)

// Parse purely parentheses-grouped Lisp code into an abstract syntax tree for output in either traditional or Pyylisp-formatted Lisp.
func ParseLisp(code string) (Pyylisp, error) {
	return parse(code, false)
}

// Parse Pyylisp-formated Lisp code into an abstract syntax tree for output in either traditional or Pyylisp-formatted Lisp.
// NOTE: Although this /should/ "just work" with regular Lisp by ignoring redundant grouping indicators, it can break with irregular indentation. Please use ParseLisp() to parse Lisp and save yourself some frustration.
func ParsePyylisp(code string) (Pyylisp, error) {
	return parse(code, true)
}

// This contains a regex pattern that should be matched by the tokenType.
type tokenType string

const (
	openp  tokenType = `\(`
	closep tokenType = `\)`
	tabs   tokenType = "\t+"
	spaces tokenType = " +"
	symbol tokenType = "[^()\t ]+" // Catch everything not caught above. TODO: Match Lisp spec.
)

type token struct {
	tokenType
	content string // the raw text that the tokenType expression matched
}

// All types of tokens, this time priority-irrelevant.
var tokenTypes = []tokenType{openp,
	closep,
	tabs,
	spaces,
	symbol,
}

// Convert raw code into tokens representing whitespace, parentheses, and strings.
func tokenize(code string) ([]token, error) {

}

// Do the actual parsing work. If whitespace is true, then infer grouping from indentation levels (Pyylisp style); otherwise only infer grouping from parentheses (traditional Lisp style).
func parse(code string, whitespace bool) (Pyylisp, error) {
	t, err := tokenize(code)
	if err != nil {
		return nil, err
	}
}
