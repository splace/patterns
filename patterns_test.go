package patterns

import (
	"fmt"
	//"io/ioutil"
	//"strings"
	"testing"
)

func ExampleXscan(){
	var v x
	_,err:=fmt.Sscan("  ,  13.45.  ",&v)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(v)
	// Output:
	// 13.45
}


func ExampleXscanMuilti(){
	var v1,v2,v3,v4 x
	_,err:=fmt.Sscan("  ,  13.45.5, -0.3-.091  ",&v1,&v2,&v3,&v4)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(&v1,&v2,&v3,&v4)
	// Output:
	// 13.45 0.5 -0.3 -0.091
}

// one step margin 
func Output(p LimitedPattern,step x){
	PrintGraph(p,-p.MaxX()-step, p.MaxX()+step, -p.MaxX()-step, p.MaxX()+step, step)
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
	Output(Limiter{Constant(unitY),5*unitX},unitX)
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
	Output(Shrunk{Disc(unitY),0.1},unitX)
	/* Output:
Graph
      -11	-----------------------
      -10	-----------X-----------
       -9	-------XXXXXXXXX-------
       -8	-----XXXXXXXXXXXXX-----
       -7	----XXXXXXXXXXXXXXX----
       -6	---XXXXXXXXXXXXXXXXX---
       -5	---XXXXXXXXXXXXXXXXX---
       -4	--XXXXXXXXXXXXXXXXXXX--
       -3	--XXXXXXXXXXXXXXXXXXX--
       -2	--XXXXXXXXXXXXXXXXXXX--
       -1	--XXXXXXXXXXXXXXXXXXX--
        0	-XXXXXXXXXXXXXXXXXXXXX-
        1	--XXXXXXXXXXXXXXXXXXX--
        2	--XXXXXXXXXXXXXXXXXXX--
        3	--XXXXXXXXXXXXXXXXXXX--
        4	--XXXXXXXXXXXXXXXXXXX--
        5	---XXXXXXXXXXXXXXXXX---
        6	---XXXXXXXXXXXXXXXXX---
        7	----XXXXXXXXXXXXXXX----
        8	-----XXXXXXXXXXXXX-----
        9	-------XXXXXXXXX-------
       10	-----------X-----------
       11	-----------------------
	*/
}

func ExampleSquare() {
	Output(Square(unitY),unitX)
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


