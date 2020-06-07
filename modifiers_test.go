package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
"testing"
)

func TestModifierMax(t *testing.T){
	if max4(-2500,2500,2500,-2500)!=2500{
		t.Error( max4(-2500,2500,2500,-2500))
	}
}



func ExampleModifiersShifted() {
	Output(Translated{Square{Filling{unitY}}, 4*unitX, 3*unitX},unitX)
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

func ExampleModifiersTranslated() {
	Output(Translated{Shrunk{Square{Filling{unitY}}, .5}, 2*unitX, 0},unitX)
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

func ExampleModifiersZoomed() {
	Output(Shrunk{Square{Filling{unitY}}, .25},unitX)
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

func ExampleModifiersScaled() {
	Output(Limiter{Reduced{Square{Filling{unitY}}, .125, 1},8*unitX},unitX)
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

func ExampleModifiersRotated() {
	Output(Limiter{Rotated{Shrunk{Square{Filling{unitY}}, .5}, .707, .707},3*unitX},unitX*.5)
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

func ExampleModifiersInverted() {
	Output(Inverted{Shrunk{Square{Filling{unitY}}, .5}},unitX)
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



