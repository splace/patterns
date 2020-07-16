package pattern

import (
	//"fmt"
	//"io/ioutil"
	//"strings"
	"testing"
)

func TestModifierMax(t *testing.T) {
	if max4(-2500, 2500, 2500, -2500) != 2500 {
		t.Error(max4(-2500, 2500, 2500, -2500))
	}
}

func ExampleModifiersTranslated() {
	Output(Translated{Square(unitY), 4 * unitX, 3 * unitX}, unitX)
	/* Output:
		Graph
	       -6	-------------
	       -5	-------------
	       -4	-------------
	       -3	-------------
	       -2	-------------
	       -1	-------------
	        0	-------------
	        1	-------------
	        2	---------XXX-
	        3	---------XXX-
	        4	---------XXX-
	        5	-------------
	        6	-------------

	*/
}

func ExampleModifiersShrunkTranslated() {
	Output(Translated{Shrunk{Square(unitY), .5}, 2 * unitX, 0}, unitX)
	/* Output:
		Graph
	       -5	-----------
	       -4	-----------
	       -3	-----------
	       -2	-----XXXXX-
	       -1	-----XXXXX-
	        0	-----XXXXX-
	        1	-----XXXXX-
	        2	-----XXXXX-
	        3	-----------
	        4	-----------
	        5	-----------
	*/
}

func ExampleModifiersScaled() {
	Output(Reduced{Square(unitY), .125, 1}, unitX)
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
	       -1	-XXXXXXXXXXXXXXXXX-
	        0	-XXXXXXXXXXXXXXXXX-
	        1	-XXXXXXXXXXXXXXXXX-
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

func ExampleHybridsFitted() {
	p := NewFitted( Square( Filling(unitY)) , 10, 2)
	Output(p, unitX)
	/* Output:
	Graph
	       -6	-------------
	       -5	-------------
	       -4	-------------
	       -3	-------------
	       -2	-------------
	       -1	-XXXXXXXXXXX-
	        0	-XXXXXXXXXXX-
	        1	-XXXXXXXXXXX-
	        2	-------------
	        3	-------------
	        4	-------------
	        5	-------------
	        6	-------------
	*/
}

func ExampleModifiersRotated() {
	Output(Limiter{Rotated{Shrunk{Square(unitY), .5}, .707, .707}, 3 * unitX}, unitX*.5)
	/* Output:
		Graph
	       -3	---------------
	       -3	---------------
	       -2	-------X-------
	       -2	------XXX------
	       -1	-----XXXXX-----
	       -1	----XXXXXXX----
	        0	---XXXXXXXXX---
	        0	--XXXXXXXXXXX--
	        0	---XXXXXXXXX---
	        1	----XXXXXXX----
	        1	-----XXXXX-----
	        2	------XXX------
	        2	-------X-------
	        3	---------------
	        3	---------------
	*/
}

func ExampleModifiersInverted() {
	Output(Inverted{Shrunk{Square(unitY), .5}}, unitX)
	/* Output:
		Graph
	       -3	XXXXXXX
	       -2	X-----X
	       -1	X-----X
	        0	X-----X
	        1	X-----X
	        2	X-----X
	        3	XXXXXXX
	*/
}
