// pyylisp.go
// Data structure representing Lisp code in a way that makes it convenient to convert between standard and Python-like forms.
// License: Public domain / Unlicense (unlicense.org)
// Author: Mark Haferkamp

package pyylisp

// An abstract syntax tree representation of Lisp code, allowing for output in whatever format makes sense.
// Note that lineLimit in each of these is a best-effort guideline to the output functions to avoid too much non-whitespace in a line. Nodes will be compressed to single lines and use parentheses up to lineLimit characters. Use 0 if you want to completely eliminate parentheses.
type Pyylisp interface {
	// Return Lisp-formatted code, for use with existing computer-side Lisp stuff.
	Lisp(lineLimit int) string
	// Return Pyylisp-formatted code, for greater clarity in human-side Lisp stuff.
	Pyylisp(lineLimit int) string
}

// These are the two types of Pyylisp thing.
type PyylispLeaf string    // the stuff without the structure
type PyylispNode []Pyylisp // the structure of the stuff

// The string is itself in any representation.
func (p *PyylispLeaf) Lisp(_ int) string {
	return string(*p)
}
func (p *PyylispLeaf) Pyylisp(_ int) string {
	return string(*p)
}

// These constants are the only difference between Lisp and Pyylisp format.
var sepLisp = [2]string{"(", ")"}
var sepPyylisp = [2]string{"", ""}

// Call the real formatting code.
func (p *PyylispNode) Lisp(lineLimit int) string {
	return p.code(lineLimit, sepLisp)
}
func (p *PyylispNode) Pyylisp(lineLimit int) string {
	return p.code(lineLimit, sepPyylisp)
}

// Generate Pyylisp output, with redundant separators as given by sep and trying to keep lines within n characters.
func (p *PyylispNode) code(n int, sep [2]string) string {

}
