package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExampleComposite() {
	p := Limiter{NewComposite(Shrunk{Disc{Filling{unitY}}, .2}, Shrunk{Square{Filling{unitY}}, .25}),5}
	PrintGraph(p,-p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}
