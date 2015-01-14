#| example.(pyy)?lisp
This is the (Pyylisp-formatted )? example code for testing the Pyylisp conversion tool.
TODO: Switch to a more sophisticated example than a solution to problem 1 of Project Euler.
|#

(defun pe1-brute-force (stopat) (
	let ((sum 0))
	(dotimes (x stopat)
		(if
			(or (= 0 (mod x 3)) (= 0 (mod x 5)))
			(setf sum (+ sum x))
		)
	)
	sum
))

(print "Brute force solution:")
(print (pe1-brute-force 1000))
