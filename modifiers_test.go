package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExamplePatternsShifted() {
	PrintGraph(Shifted{Square{3, Filling{unitY}}, 2, 1}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExamplePatternsScaled() {
	PrintGraph(Scaled{Square{3, Filling{unitY}}, .5, 1}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExamplePatternsRotated() {
	PrintGraph(Rotated{Square{2, Filling{unitY}}, .707, .707}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExamplePatternsTranslated() {
	p := Translated{Square{2, Filling{unitY}}, 2, 0}
	PrintGraph(p, -p.maxX(), p.maxX(), -p.maxX(), p.maxX(), 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}


func ExamplePatternsInverted() {
	p := LimitedInverted{Square{2, Filling{unitY}}}
	PrintGraph(p, -p.maxX()-2, p.maxX()+2, -p.maxX()-2, p.maxX()+2, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}




