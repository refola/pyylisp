# pyylisp
NOTE: This project in its current state is abandoned. Please see [piklisp.py](https://github.com/refola/piklisp.py) for (abandoned as of mid 2015) progress on making a Python-integrated version of Lisp with Python-like indentation-based grouping of code. Also see [piklisp.go](https://github.com/refola/piklisp.go) for a new (as of 2015-11-14) experimental Go-integrated Lisp that I'm trying to make for the learning experience.
-

Program Lisp in a Python-like style.

I really want to learn Lisp (the macros, as much as I don't understand them, sound *extremely* useful), but all those parentheses are distracting. Also, I think a function's name should be after the first argument, in the "object-oriented" style. Knowing neither Lisp nor Python, I think what I'm describing is essentially a Python-style isomorphism of Lisp. Thus I think I can write a program to convert between styles, write Lisp in the "Python" style, use the program to convert to regular style, and then run it as Lisp.

As I know neither Lisp nor Python, I'm programming this in Go.

DONE:
* Made README file

TODO:
* Write Lisp parser
* Swap first and second things in parenthesis-grouped list
* Convert parentheses to combinations of tabs and line breaks
* Do the above backwards
* Learn Python and Lisp
* Fix whatever this conversion breaks
