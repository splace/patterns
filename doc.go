/*
Package patterns generates, stores, downloads and manipulates abstract patterns, when imported it can then be used with specific real-world quantities.

Definition of 'pattern'

A varying property as it depends, uniquely, on two parameters.

The controlling parameters are generally unbounded.

Fundamental Types

x :- 'parameters' designed to be used as if unbounded (+ve and -ve), with unitX near the centre of its precision range.

y :- 'property' can have a value between limits, +unitY and -unitY.

Interfaces 

Pattern :- as method at(x,x)y which returns a 'y' value from two 'x' value parameters.

LimitedPattern :- a Pattern with a MaxX() method returning the the 'x' value outside which the Pattern can be assumed to return zero, it effectively has an size.

*/
package patterns

/*
Implementation details.

*/
