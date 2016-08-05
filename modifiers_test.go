package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)


func ExampleShifted() {
	p:=Translated{Shrunk{Square{Filling{unitY}}, .25}, 2, 1}
	PrintGraph(p,-p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin,1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleZoomed() {
	p:=Shrunk{Square{Filling{unitY}}, .25}
	PrintGraph(p, -p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleScaled() {
	p:=Limiter{Reduced{Shrunk{Square{Filling{unitY}}, .25}, 2, 1},8}
	PrintGraph(p, -p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleRotated() {
	p:=Limiter{Rotated{Shrunk{Square{Filling{unitY}}, .5}, .707, .707},3}
	PrintGraph(p, -p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin,1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleTranslated() {
	p := Translated{Shrunk{Square{Filling{unitY}}, .5}, 2, 0}
	PrintGraph(p,-p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin,1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleInverted() {
	p := LimitedInverted{Shrunk{Square{Filling{unitY}}, .5}}
	PrintGraph(p, -p.MaxX()-2, p.MaxX()+2, -p.MaxX()-2, p.MaxX()+2, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}
