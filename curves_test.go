package pattern

import "fmt"

//func ExampleCurvesPrint2() {
//	for p := range Divide(8).CubicBezier(0,0,19*unitX, -9*unitX, 30*unitX, -26*unitX, 20*unitX,-54*unitX) {
//		fmt.Println(p)
//	}
//	// Output:
//	//
//}

func ExampleCurvesPrint() {
	for p := range Divide(8).QuadraticBezier(0, 0, 1*unitX, 1*unitX, 2*unitX, 0) {
		fmt.Println(p)
	}
	fmt.Println()
	for p := range Divide(8).CubicBezier(0, 0, 1*unitX, 1*unitX, 1*unitX, 1*unitX, 2*unitX, 0) {
		fmt.Println(p)
	}
	fmt.Println()
	for p := range Divide(8).QuinticBezier(0, 0, 1*unitX, 1*unitX, 1*unitX, 1*unitX, 1*unitX, 1*unitX, 2*unitX, 0) {
		fmt.Println(p)
	}
	// increased order makes for pointier curve if cotrol points co-incident

	// Output:
	/*
		[0.234 0.207]
		[0.478 0.363]
		[0.72 0.461]
		[0.964 0.499]
		[1.206 0.479]
		[1.45 0.399]
		[1.694 0.26]
		[1.936 0.062]

		[0.313 0.31]
		[0.573 0.545]
		[0.784 0.691]
		[0.972 0.748]
		[1.157 0.718]
		[1.36 0.599]
		[1.603 0.389]
		[1.907 0.093]

		[0.392 0.391]
		[0.667 0.661]
		[0.848 0.815]
		[0.981 0.873]
		[1.107 0.843]
		[1.269 0.719]
		[1.513 0.486]
		[1.878 0.122]
	*/
}

func ExampleCurvesLowResArcPrint() {
	for p := range Divide(4).Arc(-2, 0, 0, 0, 0, false, false, 2, 0) {
		fmt.Println(p)
	}
	// Output:
	/*
		[-0.001 0.001]
		[0 0.002]
		[0.001 0.001]
	*/
}

func ExampleCurvesCirculerArcPrint() {
	for p := range Divide(4).Arc(-1*unitX, 0, 2*unitX, 2*unitX, 0, false, false, 1*unitX, 0) {
		fmt.Println(p)
	}
	for p := range Divide(4).Arc(-1*unitX, 0, 2*unitX, 2*unitX, 0, false, true, 1*unitX, 0) {
		fmt.Println(p)
	}
	for p := range Divide(4).Arc(-1*unitX, 0, 2*unitX, 2*unitX, 0, true, false, 1*unitX, 0) {
		fmt.Println(p)
	}
	for p := range Divide(4).Arc(-1*unitX, 0, 2*unitX, 2*unitX, 0, true, true, 1*unitX, 0) {
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

func ExampleCurvesArcPrint() {
	for p := range Divide(4).Arc(-1*unitX, 0, 2*unitX, 1*unitX, 0, false, false, 1*unitX, 0) {
		fmt.Println(p)
	}
	for p := range Divide(4).Arc(-1*unitX, 0, 2*unitX, 1*unitX, 0, false, true, 1*unitX, 0) {
		fmt.Println(p)
	}
	for p := range Divide(4).Arc(-1*unitX, 0, 2*unitX, 1*unitX, 0, true, false, 1*unitX, 0) {
		fmt.Println(p)
	}
	for p := range Divide(4).Arc(-1*unitX, 0, 2*unitX, 1*unitX, 0, true, true, 1*unitX, 0) {
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
