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

func ExampleSVGDrawPathCubicBezierToRelative() {
	b := Brush{LineBrush:LineBrush{Width: 2000, In: unitY}}
	p := Path{MoveTo{217021, 167042}, CubicBezierToRelative{18631, -9483, 30288, -26184, 27565, -54007}}
	PrintGraph(p.Draw(&b),210*unitX,300*unitX,110*unitX,200*unitX,1*unitX)
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
	_,err:=fmt.Sscan("m 0 0 l 5 5 h -5 z",&p)
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
	p:=Path{}
	_,err:=fmt.Sscan("m -5 -5 l 10 10 0 -10 z",&p)
	if err!=nil{
		fmt.Println(err)
	}
	b := Brush{LineBrush:LineBrush{Width: 2*unitX, In: unitY}}
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

func ExampleSVGpathScanCubicOverlap() {
	p:=Path{}
	_,err:=fmt.Sscan("M20,15 C-30,-30 -30,30 20,-15",&p)
	if err!=nil{
		fmt.Println(err)
	}
	b := Brush{LineBrush:LineBrush{Width: 2*unitX, In: unitY,CurveDivision:3}}
	Output(Limiter{p.Draw(&b),20*unitX})
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

func ExampleSVGpathScanSmoothCubic() {
	f:="M-%[1]v,0C-%[1]v,-%[1]v %[1]v,-%[1]v %[1]v,0S-%[1]v,%[1]v -%[1]v,0z"
	
	radius:=15
	cpath:=fmt.Sprintf(f,radius)
	radius=10
	cpath+=fmt.Sprintf(f,radius)
	radius=5
	cpath+=fmt.Sprintf(f,radius)

	p:=Path{}
	_,err:=fmt.Sscan(cpath,&p)
	if err!=nil{
		fmt.Println(err)
	}

	b := Brush{LineBrush:LineBrush{Width: 4*unitX, In: unitY, CurveDivision:3}}
	Output(Shrunk{Limiter{p.Draw(&b),18*unitX},0.5})
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




func ExampleSVGpathScanMulti() {
	b := Brush{LineBrush:LineBrush{Width: 18*unitX, In: unitY, CurveDivision:2}}
	p:=Path{}
	_,err:=fmt.Sscan(
`M117.021,167.042
c18.631,-9.483 30.288,-26.184 27.565,-54.007
c-3.667,-38.023 -36.526,-50.773 -78.006,-54.404
l-0.008,-52.741
h-32.139
l-0.009,51.354
c-8.456,0 -17.076,0.166 -25.657,0.338
L8.76,5.897
l-32.11,-0.003
l-0.006,52.728
c-6.959,0.142 -13.793,0.277 -20.466,0.277
v-0.156
l-44.33,-0.018
l0.006,34.282
c0,0 23.734,-0.446 23.343,-0.013
c13.013,0.009 17.262,7.559 18.484,14.076
l0.01,60.083 v84.397
c-0.573,4.09 -2.984,10.625 -12.083,10.637
c0.414,0.364 -23.379,-0.004 -23.379,-0.004
l-6.375,38.335
h41.817
c7.792,0.009 15.448,0.13 22.959,0.19
l0.028,53.338
l32.102,0.009
l-0.009,-52.779
c8.832,0.18 17.357,0.258 25.684,0.247
l-0.009,52.532 h32.138
l0.018,-53.249
c54.022,-3.1 91.842,-16.697 96.544,-67.385
C166.916,192.612 147.692,174.396 117.021,167.042
z
M9.535,95.321
c18.126,0 75.132,-5.767 75.14,32.064
c-0.008,36.269 -56.996,32.032 -75.14,32.032
V95.321
z
M9.521,262.447
l0.014,-70.672
c21.778,-0.006 90.085,-6.261 90.094,35.32
C99.638,266.971 31.313,262.431 9.521,262.447
z`,
		&p)
	if err!=nil{
		fmt.Println(err)
	}
//	fmt.Printf("%#v\n",p)
	Output(Limiter{UnlimitedShrunk{p.Draw(&b),6},60*unitX})
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


