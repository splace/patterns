package patterns

import "fmt"

func ExampleCurvesPrint(){
	for p:=range(Divide(8).QuadraticBezier(0,0,1*unitX,1*unitX,2*unitX,0)){
		fmt.Println(p)
	}
	// Output:
	/* [0.234 0.207]
[0.478 0.363]
[0.72 0.461]
[0.964 0.499]
[1.206 0.479]
[1.45 0.399]
[1.694 0.26]
[1.936 0.062]
*/
}

func ExampleCurvesLowResArcPrint(){
	for p:=range(Divide(4).Arc(-2,0,0,0,0,false,false,2,0)){
		fmt.Println(p)
	}
	// Output:
	/*
[-0.001 0.001]
[0 0.002]
[0.001 0.001]
*/
}



func ExampleCurvesCirculerArcPrint(){
	for p:=range(Divide(4).Arc(-1*unitX,0,2*unitX,2*unitX,0,false,false,1*unitX,0)){
		fmt.Println(p)
	}
	for p:=range(Divide(4).Arc(-1*unitX,0,2*unitX,2*unitX,0,false,true,1*unitX,0)){
		fmt.Println(p)
	}
	for p:=range(Divide(4).Arc(-1*unitX,0,2*unitX,2*unitX,0,true,false,1*unitX,0)){
		fmt.Println(p)
	}
	for p:=range(Divide(4).Arc(-1*unitX,0,2*unitX,2*unitX,0,true,true,1*unitX,0)){
		fmt.Println(p)
	}
	// Output:
	/* [-0.517 -0.199]
[0 -0.267]
[0.517 -0.199]
[-0.517 0.199]
[0 0.267]
[0.517 0.199]
[-1.931 -2.249]
[0 -3.732]
[1.931 -2.249]
[-1.931 2.249]
[0 3.732]
[1.931 2.249]
*/
}

func ExampleCurvesArcPrint(){
	for p:=range(Divide(4).Arc(-1*unitX,0,2*unitX,1*unitX,0,false,false,1*unitX,0)){
		fmt.Println(p)
	}
	for p:=range(Divide(4).Arc(-1*unitX,0,2*unitX,1*unitX,0,false,true,1*unitX,0)){
		fmt.Println(p)
	}
	for p:=range(Divide(4).Arc(-1*unitX,0,2*unitX,1*unitX,0,true,false,1*unitX,0)){
		fmt.Println(p)
	}
	for p:=range(Divide(4).Arc(-1*unitX,0,2*unitX,1*unitX,0,true,true,1*unitX,0)){
		fmt.Println(p)
	}
	// Output:
	/* [-0.517 -0.199]
[0 -0.267]
[0.517 -0.199]
[-0.517 0.199]
[0 0.267]
[0.517 0.199]
[-1.931 -2.249]
[0 -3.732]
[1.931 -2.249]
[-1.931 2.249]
[0 3.732]
[1.931 2.249]
*/
}

