package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)


func ExampleShifted() {
	Output(Translated{Shrunk{Square{Filling{unitY}}, .5}, 2*unitX, 2*unitX})
	/* Output:
Graph
       -5	-----------
       -4	-----------
       -3	-----------
       -2	-----------
       -1	-----------
        0	-----XXXX--
        1	-----XXXX--
        2	-----XXXX--
        3	-----XXXX--
        4	-----------
        5	-----------
	*/
}

func ExampleTranslated() {
	Output(Translated{Shrunk{Square{Filling{unitY}}, .5}, 2*unitX, 0})
	/* Output:
Graph
       -5	-----------
       -4	-----------
       -3	-----------
       -2	-----XXXX--
       -1	-----XXXX--
        0	-----XXXX--
        1	-----XXXX--
        2	-----------
        3	-----------
        4	-----------
        5	-----------
	*/
}

func ExampleZoomed() {
	Output(Shrunk{Square{Filling{unitY}}, .25})
	/* Output:
Graph
       -5	-----------
       -4	-XXXXXXXX--
       -3	-XXXXXXXX--
       -2	-XXXXXXXX--
       -1	-XXXXXXXX--
        0	-XXXXXXXX--
        1	-XXXXXXXX--
        2	-XXXXXXXX--
        3	-XXXXXXXX--
        4	-----------
        5	-----------
	*/
}

func ExampleScaled() {
	Output(Limiter{Reduced{Square{Filling{unitY}}, .125, 1},8*unitX})
	/* Output:
Graph
       -9	-------------------
       -8	-------------------
       -7	-------------------
       -6	-------------------
       -5	-------------------
       -4	-------------------
       -3	-------------------
       -2	-------------------
       -1	-XXXXXXXXXXXXXXXX--
        0	-XXXXXXXXXXXXXXXX--
        1	-------------------
        2	-------------------
        3	-------------------
        4	-------------------
        5	-------------------
        6	-------------------
        7	-------------------
        8	-------------------
        9	-------------------
	*/
}

func ExampleRotated() {
	Output(Limiter{Rotated{Shrunk{Square{Filling{unitY}}, .5}, .707, .707},3*unitX})
	/* Output:
Graph
       -4	---------
       -3	---------
       -2	----X----
       -1	---XXX---
        0	--XXXXX--
        1	---XXX---
        2	----X----
        3	---------
        4	---------
	*/
}

func ExampleInverted() {
	Output(LimitedInverted{Shrunk{Square{Filling{unitY}}, .5}})
	/* Output:
Graph
       -3	XXXXXXX
       -2	X----XX
       -1	X----XX
        0	X----XX
        1	X----XX
        2	XXXXXXX
        3	XXXXXXX
	*/
}
