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
	examplelisp, err = ioutil.ReadFile("tests/example.lisp")
	if err != nil {
		panic()
	}
	examplepyylisp, err = ioutil.ReadFile("tests/example.pyylisp")
	if err != nil {
		panic()
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
	pl := ParseLisp(examplelisp).Pyylisp()
	if pl != examplepyylisp {
		t.Log("Generated Pyylisp doesn't match file.")
	}
	l := ParsePyylisp(pl).Lisp()
	if l != examplelisp {
		t.Error("Converting back to Lisp failed.")
	}
}
