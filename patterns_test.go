package patterns

import (
	"fmt"
	//"io/ioutil"
	//"strings"
	"testing"
)

func PrintGraph(p Pattern, startx, endx,starty, endy, step x) {
	for py := starty; py < endy; py += step {
		row := make([]byte,int((endx-startx)/step))
		for px := startx; px < endx; px += step {
			row[int((px-startx)/step)]= p.property(px,py).String()[0]
		}
		fmt.Println(py,string(row))
	}
}

func ExamplePatternsConstant() {
	PrintGraph(Constant{true}, -3, 3, -3,3, 1)
	/* Output:
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/
}

func ExamplePatternsDisc() {
	PrintGraph(Disc{3,true}, -5,5, -5,5, 1)
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

func BenchmarkPattersSineSegmented(b *testing.B) {
	b.StopTimer()
	b.StartTimer()

}



