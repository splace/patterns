package patterns

import "fmt"


func ExampleSVGpath() {
	p:=Path{MoveTo{0,0},LineTo{10,10}}
	fmt.Printf("%#v\n",p)
	// Output:
	// patterns.Path{patterns.MoveTo{0, 0}, patterns.LineTo{10, 10}}
}

func ExampleSVGDrawPath() {
	p := Path{MoveTo{0,0},LineTo{10*unitX,10*unitX}}
	b := Brush{Pen:Pen{Nib:Facetted{Width: 2*unitX, In: unitY, CurveDivision:2}}}
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

//func ExampleSVGDrawPathCubicBezierToRelative() {
//	b := Brush{LineBrush:LineBrush{Width: 2000, In: unitY}}
//	p := Path{MoveTo{217021, 167042}, CubicBezierToRelative{18631, -9483, 30288, -26184, 27565, -54007}}
//	PrintGraph(p.Draw(&b),210*unitX,300*unitX,110*unitX,200*unitX,1*unitX)
//	/* Output:
//Graph
//      -10	---------------------
//       -9	---------------------
//       -8	---------------------
//       -7	---------------------
//       -6	---------------------
//       -5	---------------------
//       -4	---------------------
//       -3	---------------------
//       -2	---------------------
//       -1	---------------------
//        0	----------X----------
//        1	-----------X---------
//        2	------------X--------
//        3	-------------X-------
//        4	--------------X------
//        5	---------------X-----
//        6	----------------X----
//        7	-----------------X---
//        8	------------------X--
//        9	-------------------X-
//       10	--------------------X
//	*/
//}



//func ExampleSVGpathScan() {
//	b := Brush{LineBrush:LineBrush{Width: 2*unitX, In: unitY}}
//	p:=Path{}
//	_,err:=fmt.Sscan("m 0 0 5 5 h -10 z",&p)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	Output(Limiter{p.Draw(&b),10*unitX})
//	/* Output:
//Graph
//       -8	-----------------
//       -7	-----------------
//       -6	--XXXXXXXXXXXX---
//       -5	--XXXXXXXXXXXX---
//       -4	--XX--------XX---
//       -3	--XX--------XX---
//       -2	--XX--------XX---
//       -1	--XX--------XX---
//        0	--XX--------XX---
//        1	--XX--------XX---
//        2	--XX--------XX---
//        3	--XX--------XX---
//        4	--XXXXXXXXXXXX---
//        5	--XXXXXXXXXXXX---
//        6	-----------------
//        7	-----------------
//        8	-----------------
//	*/
//}

//func ExampleSVGpathScanMissingCommands() {
//	p:=Path{}
//	_,err:=fmt.Sscan("m -5 -5 l 10 10 0 -10 z",&p)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	b := Brush{LineBrush:LineBrush{Width: 2*unitX, In: unitY}}
//	Output(Limiter{p.Draw(&b),10*unitX})
//	/* Output:
//Graph
//       -8	-----------------
//       -7	-----------------
//       -6	--XXXXXXXXXXXX---
//       -5	--XXXXXXXXXXXX---
//       -4	--XX--------XX---
//       -3	--XX--------XX---
//       -2	--XX--------XX---
//       -1	--XX--------XX---
//        0	--XX--------XX---
//        1	--XX--------XX---
//        2	--XX--------XX---
//        3	--XX--------XX---
//        4	--XXXXXXXXXXXX---
//        5	--XXXXXXXXXXXX---
//        6	-----------------
//        7	-----------------
//        8	-----------------
//	*/
//}

//func ExampleSVGpathScanCubicOverlap() {
//	p:=Path{}
//	_,err:=fmt.Sscan("M20,15 C-30,-30 -30,30 20,-15",&p)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	b := Brush{LineBrush:LineBrush{Width: 2*unitX, In: unitY,CurveDivision:3}}
//	Output(Limiter{p.Draw(&b),20*unitX})
//	/* Output:
//Graph
//       -8	-----------------
//       -7	-----------------
//       -6	--XXXXXXXXXXXX---
//       -5	--XXXXXXXXXXXX---
//       -4	--XX--------XX---
//       -3	--XX--------XX---
//       -2	--XX--------XX---
//       -1	--XX--------XX---
//        0	--XX--------XX---
//        1	--XX--------XX---
//        2	--XX--------XX---
//        3	--XX--------XX---
//        4	--XXXXXXXXXXXX---
//        5	--XXXXXXXXXXXX---
//        6	-----------------
//        7	-----------------
//        8	-----------------
//	*/
//}

//func ExampleSVGpathScanSmoothQuadratic() {

//	p:=Path{}
//	_,err:=fmt.Sscan("M-50,30 Q-40,5 -20,30 T40,30",&p)
//	if err!=nil{
//		fmt.Println(err)
//	}
//	b := Brush{LineBrush:LineBrush{Width: 1*unitX, In: unitY, CurveDivision:3}}
//	Output(Limiter{p.Draw(&b),50*unitX})
//	/* Output:
//Graph
//       -8	-----------------
//       -7	-----------------
//       -6	--XXXXXXXXXXXX---
//       -5	--XXXXXXXXXXXX---
//       -4	--XX--------XX---
//       -3	--XX--------XX---
//       -2	--XX--------XX---
//       -1	--XX--------XX---
//        0	--XX--------XX---
//        1	--XX--------XX---
//        2	--XX--------XX---
//        3	--XX--------XX---
//        4	--XXXXXXXXXXXX---
//        5	--XXXXXXXXXXXX---
//        6	-----------------
//        7	-----------------
//        8	-----------------
//	*/
//}


//func ExampleSVGpathScanSmoothCubic() {
//	f:="M-%[1]v,0C-%[1]v,-%[1]v %[1]v,-%[1]v %[1]v,0S-%[1]v,%[1]v -%[1]v,0z"
//	
//	radius:=15
//	cpath:=fmt.Sprintf(f,radius)
//	radius=10
//	cpath+=fmt.Sprintf(f,radius)
//	radius=5
//	cpath+=fmt.Sprintf(f,radius)

//	p:=Path{}
//	_,err:=fmt.Sscan(cpath,&p)
//	if err!=nil{
//		fmt.Println(err)
//	}

//	b := Brush{LineBrush:LineBrush{Width: 4*unitX, In: unitY, CurveDivision:3}}
//	Output(Shrunk{Limiter{p.Draw(&b),18*unitX},0.5})
//	/* Output:
//Graph
//       -8	-----------------
//       -7	-----------------
//       -6	--XXXXXXXXXXXX---
//       -5	--XXXXXXXXXXXX---
//       -4	--XX--------XX---
//       -3	--XX--------XX---
//       -2	--XX--------XX---
//       -1	--XX--------XX---
//        0	--XX--------XX---
//        1	--XX--------XX---
//        2	--XX--------XX---
//        3	--XX--------XX---
//        4	--XXXXXXXXXXXX---
//        5	--XXXXXXXXXXXX---
//        6	-----------------
//        7	-----------------
//        8	-----------------
//	*/
//}




func ExampleSVGpathScanMulti() {
	p:=Path{}
	_,err:=fmt.Sscan(
`M 117 167 
c 19 -9 30 -26 28 -54
c -4 -38 -37 -51 -78 -54
l 0 -53
h -32
l 0 51
c -8 0 -17 0 -26 0
L 9 6
l -32 0
l 0 53
c -7 0 -14 0 -20 0
v 0
l -44 0
l 0 34
c 0 0 24 0 23 0
c 13 0 17 8 18 14
l 0 60
v 84
c -1 4 -3 11 -12 11
c 0 0 -23 0 -23 0
l -6 38
h 42
c 8 0 15 0 23 0
l 0 53
l 32 0
l 0 -53
c 9 0 17 0 26 0
l 0 53
h 32
l 0 -53
c 54 -3 92 -17 97 -67
C 167 193 148 174 117 167
z
M 10 95
c 18 0 75 -6 75 32
c 0 36 -57 32 -75 32
V 95
z
M 10 262
l 0 -71
c 22 0 90 -6 90 35
C 100 267 31 262 10 262
z
`,
		&p)
	if err!=nil{
		fmt.Println(err)
	}
//	fmt.Printf("%#v\n",p)
	b := Brush{Pen:Pen{Nib:Facetted{Width: 10*unitX, In: unitY, CurveDivision:2}}}
	Output(Limiter{UnlimitedShrunk{p.Draw(&b),6},60*unitX},unitX)
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

