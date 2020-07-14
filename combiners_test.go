package pattern

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExampleComposite() {
	Output(NewFrame(5, 2, Filling(unitY)), unitX)
	/* Output:
	Graph
	       -8	-----------------
	       -7	-----------------
	       -6	--XXXXXXXXXXXX---
	       -5	--XXXXXXXXXXXX---
	       -4	--XX--------XX---
	       -3	--XX--------XX---
	       -2	--XX--------XX---
	       -1	--XX--------XX---
	        0	--XX--------XX---
	        1	--XX--------XX---
	        2	--XX--------XX---
	        3	--XX--------XX---
	        4	--XXXXXXXXXXXX---
	        5	--XXXXXXXXXXXX---
	        6	-----------------
	        7	-----------------
	        8	-----------------
	*/
}
