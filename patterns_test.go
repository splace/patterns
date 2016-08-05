package patterns

import (
	"fmt"
	//"io/ioutil"
	//"strings"
	"testing"
)

const margin=4

func PrintGraph(p Pattern, startx, endx, starty, endy, step x) {
	//fmt.Printf("%#v\n",p)
	fmt.Println("Graph")
	row := make([]byte, int((endx-startx)/step)+1)
	for py := starty; py <= endy; py += step {
		for px := startx; px <= endx; px += step {
			row[int((px-startx)/step)] = p.at(px, py).String()[0]
		}
		fmt.Printf("% 9d\t%s\n",py, string(row))
	}
}

func ExampleConstant() {
	p:=Limiter{Constant{Filling{unitY}},5}
	PrintGraph(p, -p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleDisc() {
	p:=Disc{Filling{unitY}}
	PrintGraph(p, -p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleSquare() {
	p:=Square{Filling{unitY}}
	PrintGraph(p, -p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, 1)
	/* Output:
	   0.00%                                  X
	   0.00%                                  X
	   0.00%                                  X
	*/
}

func ExampleBox() {
	p := NewBox(5, 1, Filling{unitY})
	PrintGraph(p, -p.MaxX()-margin, p.MaxX()+margin, -p.MaxX()-margin, p.MaxX()+margin, 1)
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




