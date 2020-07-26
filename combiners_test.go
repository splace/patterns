package pattern

import (
"fmt"
//"io/ioutil"
//"strings"
//"testing"
)


func ExampleCombinersCentred() {
	fmt.Printf("%#v\n",Recentre(Composite{Translated{Disc(unitY), 10*unitX,10*unitX},Translated{Disc(unitY), 12*unitX,12*unitX}}))
	/* Output:
	pattern.Limiter{Unlimited:pattern.UnlimitedTranslated{Unlimited:pattern.Limiter{Unlimited:pattern.Translated{Limited:pattern.Composite{pattern.Translated{Limited:true, X:10000, Y:10000}, pattern.Translated{Limited:true, X:12000, Y:12000}}, X:-11000, Y:-11000}, Max:2000}, X:11000, Y:11000}, Max:13000}
	*/
}

