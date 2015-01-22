// parselisp_test.go
// Test parsing of Lisp

package pyylisp

import (
	"io/ioutil"
	"testing"
)

var (
	examplelisp    string
	examplepyylisp string // TODO: Handle multiple pyylisp files for different levels of conversion; parentheses are still sometimes useful for compactness.
)

func init() {
	var err error
	var b []byte

	lispfile := "tests/example.lisp"
	b, err = ioutil.ReadFile(lispfile)
	if err == nil {
		examplelisp = string(b)
	} else {
		panic("Could not read file: " + lispfile)
	}

	pyylispfile := "tests/example.pyylisp"
	b, err = ioutil.ReadFile(pyylispfile)
	if err == nil {
		examplepyylisp = string(b)
	} else {
		panic("Could not read file: " + pyylispfile)
	}
}

func TestParseLisp(t *testing.T) {
	code, err := ParseLisp(examplelisp)

	// does this even pretend to work?
	if code == nil || err != nil {
		t.Fail()
	}
}

func TestParsePyylisp(t *testing.T) {
	code, err := ParsePyylisp(examplepyylisp)

	// does this even pretend to work?
	if code == nil || err != nil {
		t.Fail()
	}
}

func TestLispToPyylispAndBack(t *testing.T) {
	parsed, err := ParseLisp(examplelisp)
	if err != nil {
		t.Errorf("Error generating syntax tree: %s", err)
	}
	pl := parsed.Pyylisp(0)
	if pl != examplepyylisp {
		t.Log("Generated Pyylisp doesn't match file.")
	}
	parsed, err = ParsePyylisp(pl)
	if err != nil {
		t.Errorf("Error generating syntax tree: %s", err)
	}
	l := parsed.Lisp(0)
	if l != examplelisp {
		t.Error("Converting back to Lisp failed.")
	}
}
