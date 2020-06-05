package patterns

//import "fmt"


func ExampleHybridsLine() {
	p := NewLine(10*unitX, unitX, Filling{unitY})
	Output(p,unitX)
	/* Output:
Graph
      -15	--------------------------------
     */
}

func ExampleBorderedInverse(){
	p:=NewBorderedInverse(NewLine(10*unitX, unitX, Filling{unitY}),unitX)
	Output(p,unitX)
	/* Output:
Graph
      -15	--------------------------------
     */

}