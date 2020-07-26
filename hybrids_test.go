package pattern

//import "fmt"

func ExampleHybridsFrame() {
	Output(NewFrame(5, 2, Filling(unitY)), unitX)
	/* Output:
	Graph
       -8	XXXXXXXXXXXXXXXXX
       -7	XXXXXXXXXXXXXXXXX
       -6	XXXXXXXXXXXXXXXXX
       -5	XXXXXXXXXXXXXXXXX
       -4	XXXX--------XXXXX
       -3	XXXX--------XXXXX
       -2	XXXX--------XXXXX
       -1	XXXX--------XXXXX
        0	XXXX--------XXXXX
        1	XXXX--------XXXXX
        2	XXXX--------XXXXX
        3	XXXX--------XXXXX
        4	XXXXXXXXXXXXXXXXX
        5	XXXXXXXXXXXXXXXXX
        6	XXXXXXXXXXXXXXXXX
        7	XXXXXXXXXXXXXXXXX
        8	XXXXXXXXXXXXXXXXX
	*/
}

