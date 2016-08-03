package patterns

import (
	"fmt"
	//"io/ioutil"
	//"strings"
	"testing"
)

func PrintGraph(p Pattern, startx, endx, starty, endy, step x) {
	row := make([]byte, int((endx-startx)/step))
	for py := starty; py < endy; py += step {
		for px := startx; px < endx; px += step {
			row[int((px-startx)/step)] = p.at(px, py).String()[0]
		}
		fmt.Println(py, string(row))
	}
}

func ExampleConstant() {
	PrintGraph(Constant{Filling{unitY}}, -3, 3, -3, 3, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleDisc() {
	PrintGraph(Disc{3, Filling{unitY}}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleSquare() {
	PrintGraph(Square{3, Filling{unitY}}, -5, 5, -5, 5, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}


func ExampleBox() {
	p := NewBox(5,2,Filling{unitY})
	PrintGraph(p, -p.maxX()-2, p.maxX()+2, -p.maxX()-2, p.maxX()+2, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
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
