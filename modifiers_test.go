package patterns

import (
	//"fmt"
	//"io/ioutil"
	//"strings"
	//"testing"
)

func ExamplePatternsShifted() {
	PrintGraph(Shifted{Square{3,true},2,1}, -5, 5, -5,5, 1)
	/* Output:
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/
}


func ExamplePatternsScaled() {
	PrintGraph(Scaled{Square{3,true},.5,1}, -5, 5, -5,5, 1)
	/* Output:
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/
}
func ExamplePatternsRotated() {
	PrintGraph(Rotated{Square{2,true},.707,.707}, -5, 5, -5,5, 1)
	/* Output:
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/
}
