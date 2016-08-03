package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExampleShifted() {
	PrintGraph(Shifted{Square{3, Filling{unitY}}, 2, 1}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleScaled() {
	PrintGraph(Scaled{Square{3, Filling{unitY}}, .5, 1}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleRotated() {
	PrintGraph(Rotated{Square{2, Filling{unitY}}, .707, .707}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleTranslated() {
	p := Translated{Square{2, Filling{unitY}}, 2, 0}
	PrintGraph(p, -p.maxX(), p.maxX(), -p.maxX(), p.maxX(), 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}


func ExampleInverted() {
	p := LimitedInverted{Square{2, Filling{unitY}}}
	PrintGraph(p, -p.maxX()-2, p.maxX()+2, -p.maxX()-2, p.maxX()+2, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}




