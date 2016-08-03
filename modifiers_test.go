package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExamplePatternsShifted() {
	PrintGraph(Shifted{Square{3, unitY}, 2, 1}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExamplePatternsScaled() {
	PrintGraph(Scaled{Square{3, unitY}, .5, 1}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExamplePatternsRotated() {
	PrintGraph(Rotated{Square{2, unitY}, .707, .707}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExamplePatternsTranslated() {
	p:=Translated{Square{2, unitY},2, 0}
	PrintGraph(p, -p.maxX(), p.maxX(), -p.maxX(), p.maxX(), 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}
