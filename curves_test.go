package patterns

import "fmt"

func ExampleCurvesPrint(){
	for p:=range(Divider(8).QuadraticBezier(0,0,1*unitX,1*unitX,2*unitX,0)){
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
