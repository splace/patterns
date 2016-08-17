package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExampleBrushLine() {
	b := Brush{Width: 2, In: unitY, Relative: true}
	Output(b.Line(0, 0, 10, 10))
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}


