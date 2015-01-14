// pyylisp.go
// Data structure representing Lisp code in a way that makes it convenient to convert between standard and Python-like forms.
// License: Public domain / Unlicense (unlicense.org)
// Author: Mark Haferkamp

package pyylisp

// This must be able to contain a string or a list of PyyLisp objects.
type Pyylisp interface {
	String() string
}
