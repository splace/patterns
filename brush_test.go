package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExampleBrushLine() {
	b := Brush{Width: 2, In: unitY, Relative: true}
	p:=b.Line(0, 0, 10, 10)
	PrintGraph(p, -p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}
