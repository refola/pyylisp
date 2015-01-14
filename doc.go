// doc.go
// Documentation for Pyylisp, a Go program for writing Lisp in a Python-like style.

/*
Overview

This package parses and converts between normal and Pyylisp-formatted Lisp code.


Pyylisp Tutorial

Pyylisp aims to make learning/using Lisp easier to use for people who don't like seeing that many parentheses. It does this by using a simple matching rule to determine which parentheses can be inferred from the surrounding indentation and removing them. It also converts some functions (function mainarg arg2 arg3 ...) to the format of (mainarg.function arg2 arg3 ...), in order to simulate an "object-oriented" programming style.

As an example, consider this Lisp code that defines a function to get the integer part of m/n.
	(defun intdiv
		(m n)
		(nth-value
			0
			(floor (/ m n)); floor returns (quotient, remainder)
		)
	)
Notice how on each line the indentation level is equal to the number of unclosed parentheses before the line? This means that we can rewrite the code in "Pyylisp" format like so.
	defun intdiv
		m n
		nth-value
			0
			floor (/ m n); floor returns (quotient, remainder)
Isn't that easier to read? Multiple space-separated things on the same line are assumed to be part of the same group, so the only parentheses needed are in (/ m n). Simple enough so far. Now for a few more changes.
	intdiv.defun
		m n
		nth-value
			0
			floor; floor returns (quotient, remainder)
				m./ n
It's still clear from indentation what's grouped with what, but look at the dots! This is the "object-oriented" notation. This doesn't look like Lisp anymore, but it's still pretty clear how to convert it back. In case this seems a bit too spacious, Pyylisp still allows you to use Lisp's parenthetical-grouping syntax, like so.
	intdiv.defun (m n)
		nth-value 0
			floor (m./ n); floor returns (quotient, remainder)
The idea is to strike a balance between compactness and clarity. Parentheses are useful for compactness, but too many of them get confusing, especially for someone new to Lisp. Compare that compact representation to this version converted back to standard Lisp syntax.
	(defun intdiv (m n)
		(nth-value 0
			(floor (/ m n))))
It's still pretty clear, but that's a lot of parentheses. Even with only four closing parentheses, it was hard enough for me to parse that I relied on my text editor's bracket-matching to know when I had enough. But you don't have to deal with the messier code representation; that's what computers are for.


Pyylisp Format

An opening parenthesis is implied at the beginning of any line that's less-indented than the next line or that contains an item-separating space. A closing parenthesis is implied at the end of any line that's more-indented than the next line or that contains an item-separating space. I'm pretty sure that this covers all uses of parentheses except for singleton lists of the form (single-item).

A period that's not part of a string and is between two valid Lisp symbols swaps the symbols and is replaced with a space. So (thing.function arg arg) becomes (function thing arg arg). Thus you can use "object-oriented" syntax where it makes sense.


*/
package pyylisp
