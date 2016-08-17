package patterns

import (
//"fmt"
//"io/ioutil"
//"strings"
//"testing"
)

func ExampleComposite() {
	Output(Limiter{NewComposite(Shrunk{Disc{Filling{unitY}}, .25}),unitX*4})
	/* Output:
Graph
       -5	-----------
       -4	-----X-----
       -3	---XXXXX---
       -2	--XXXXXXX--
       -1	--XXXXXXX--
        0	-XXXXXXXX--
        1	--XXXXXXX--
        2	--XXXXXXX--
        3	---XXXXX---
        4	-----------
        5	-----------
	*/
}
