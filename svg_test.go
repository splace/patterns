package patterns

import "fmt"


func ExampleSVGpath() {
	p:=Path{MoveTo{0,0},LineTo{10,10}}
	fmt.Printf("%#v\n",p)
	// Output:
	// patterns.Path{patterns.MoveTo{0, 0}, patterns.LineTo{10, 10}}
}

func ExampleSVGDrawPath() {
	b := Brush{LineBrush:LineBrush{Width: 2*unitX, In: unitY}}
	p := Path{MoveTo{0,0},LineTo{10*unitX,10*unitX}}
	PrintGraph(p.Draw(&b),-10*unitX,10*unitX,-10*unitX,10*unitX,unitX)
	/* Output:
Graph
      -10	---------------------
       -9	---------------------
       -8	---------------------
       -7	---------------------
       -6	---------------------
       -5	---------------------
       -4	---------------------
       -3	---------------------
       -2	---------------------
       -1	---------------------
        0	----------X----------
        1	-----------X---------
        2	------------X--------
        3	-------------X-------
        4	--------------X------
        5	---------------X-----
        6	----------------X----
        7	-----------------X---
        8	------------------X--
        9	-------------------X-
       10	--------------------X
	*/
}


func ExampleSVGpathScan() {
	b := Brush{LineBrush:LineBrush{Width: 2*unitX, In: unitY}}
	p:=Path{}
	_,err:=fmt.Sscan("m 0 0 l 10 10 h -10 z",&p)
	if err!=nil{
		fmt.Println(err)
	}
	Output(Limiter{p.Draw(&b),10*unitX})
	/* Output:
Graph
       -8	-----------------
       -7	-----------------
       -6	--XXXXXXXXXXXX---
       -5	--XXXXXXXXXXXX---
       -4	--XX--------XX---
       -3	--XX--------XX---
       -2	--XX--------XX---
       -1	--XX--------XX---
        0	--XX--------XX---
        1	--XX--------XX---
        2	--XX--------XX---
        3	--XX--------XX---
        4	--XXXXXXXXXXXX---
        5	--XXXXXXXXXXXX---
        6	-----------------
        7	-----------------
        8	-----------------
	*/
}

func ExampleSVGpathScanMissingCommands() {
	b := Brush{LineBrush:LineBrush{Width: 2*unitX, In: unitY}}
	p:=Path{}
	_,err:=fmt.Sscan("m -5 -5 l 10 10 0 -10 z",&p)
	if err!=nil{
		fmt.Println(err)
	}
	Output(Limiter{p.Draw(&b),10*unitX})
	/* Output:
Graph
       -8	-----------------
       -7	-----------------
       -6	--XXXXXXXXXXXX---
       -5	--XXXXXXXXXXXX---
       -4	--XX--------XX---
       -3	--XX--------XX---
       -2	--XX--------XX---
       -1	--XX--------XX---
        0	--XX--------XX---
        1	--XX--------XX---
        2	--XX--------XX---
        3	--XX--------XX---
        4	--XXXXXXXXXXXX---
        5	--XXXXXXXXXXXX---
        6	-----------------
        7	-----------------
        8	-----------------
	*/
}



