/*
Package patterns generates, and manipulates abstract patterns, no colour specified, when imported it can then be used with specific image types.

Definition of 'pattern'

A varying property depending, uniquely, on two parameters.

The controlling parameters are generally unbounded.

Fundamental Types

x :- 'parameters' designed to be used as if unbounded (+ve and -ve), with unitX near the centre of its precision range.

y :- 'property' can have a value interpolated between limits of +-unitY.

Interfaces 

Pattern :- as method at(x,x)y which returns a 'y' value from two 'x' value parameters.

LimitedPattern :- a Pattern with a MaxX() method returning the 'x' value range outside which the Pattern can be assumed to return a value interpreted by the Transparency() Method as completely see-through, it effectively has an size.

*/
package patterns

/*
Implementation details.

*/
