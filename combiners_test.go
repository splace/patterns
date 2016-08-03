package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExampleComposite() {
	p := NewComposite(Disc{5, Filling{unitY}}, Square{4, Filling{unitY}})
	PrintGraph(p, -10, 10, -10, 10, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}
