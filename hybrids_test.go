package patterns

//import "fmt"


func ExampleHybridsLine() {
	p := Rectangle(10*unitX, 2*unitX, Filling{unitY})
	Output(p,unitX)
	/* Output:
Graph
       -6	-------------
       -5	-------------
       -4	-------------
       -3	-------------
       -2	-------------
       -1	-XXXXXXXXXX--
        0	-XXXXXXXXXX--
        1	-------------
        2	-------------
        3	-------------
        4	-------------
        5	-------------
        6	-------------
       */
}

//func ExampleBorderedInverse(){
//	p:=NewBorderedInverse(NewLine(10*unitX, unitX, Filling{unitY}),unitX)
//	Output(p,unitX)
//	/* Output:
//Graph
//      -15	--------------------------------
//     */

//}