package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExampleBrushLine() {
	b := Brush{Width: 2, In: unitY, Relative: true}
	PrintGraph(b.Line(0, 0, 3, 3), -10, 10, -10, 10, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}
