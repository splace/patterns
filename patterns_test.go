package patterns

import (
	"fmt"
	//"io/ioutil"
	//"strings"
	"testing"
)

const margin=unitX

func Output(p LimitedPattern){
	PrintGraph(p,-p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, unitX)
}

func PrintGraph(p Pattern, startx, endx, starty, endy, step x) {
	fmt.Println("Graph")
	row := make([]byte, int((endx-startx)/step)+1)
	for py := starty; py <= endy; py += step {
		for px := startx; px <= endx; px += step {
			row[int((px-startx)/step)] = p.at(px, py).String()[0]
		}
		fmt.Printf("% 9d\t%s\n",py/unitX, string(row))
	}
}

func ExampleConstant() {
	Output(Limiter{Constant{Filling{unitY}},5*unitX})
	/* Output:
Graph
       -6	XXXXXXXXXXXXX
       -5	XXXXXXXXXXXXX
       -4	XXXXXXXXXXXXX
       -3	XXXXXXXXXXXXX
       -2	XXXXXXXXXXXXX
       -1	XXXXXXXXXXXXX
        0	XXXXXXXXXXXXX
        1	XXXXXXXXXXXXX
        2	XXXXXXXXXXXXX
        3	XXXXXXXXXXXXX
        4	XXXXXXXXXXXXX
        5	XXXXXXXXXXXXX
        6	XXXXXXXXXXXXX
	*/
}

func ExampleDisc() {
	Output(Disc{Filling{unitY}})
	/* Output:
Graph
       -2	-----
       -1	--X--
        0	-XXX-
        1	--X--
        2	-----
	*/
}

func ExampleSquare() {
	Output(Square{Filling{unitY}})
	/* Output:
Graph
       -2	-----
       -1	-XX--
        0	-XX--
        1	-----
        2	-----
	*/
}


func BenchmarkPatternsSine(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

}

func BenchmarkPatternsSineSegmented(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

}


