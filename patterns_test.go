package patterns

import (
	"fmt"
	//"io/ioutil"
	//"strings"
	"testing"
)

func PrintGraph(p Pattern, startx, endx,starty, endy, step x) {
	for py := starty; py < endy; py += step {
		row := make([]byte,int(4*(endx-startx)/step))
		for px := startx; px < endx; px += step {
			row[int(4*(px-startx)/step)]= byte(p.property(px,py))
			row[int(4*(px-startx)/step)+1]= byte(p.property(px,py)>>8)
			row[int(4*(px-startx)/step)+2]= byte(p.property(px,py)>>16)
			row[int(4*(px-startx)/step)+3]= byte(p.property(px,py)>>24)
		}
		fmt.Println(py,row)
	}
}

func ExamplePatternsConstant() {
	PrintGraph(Constant{1}, -3, 3, -3,3, 1)
	/* Output:
   0.00%                                  X
   0.00%                                  X
   0.00%                                  X
	*/
}

func ExamplePatternsDisc() {
	PrintGraph(Disc{4,0xff}, -5,5, -5,5, 1)
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

