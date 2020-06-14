package patterns

import "fmt"
//import "testing"


func ExampleSVGPathNib() {
	cbox := ChamferedBox(40*unitX,30*unitX, 4*unitX)
	p:=new(Path)
	b := NewBrush((*SimpleSvgPathNib)(p))
	cbox.Draw(b)
	fmt.Println(CompactStringer(*p))
	// Output:
	// L36,30 40,26 40,-26 36,-30 -36,-30 -40,-26 -40,26 -36,30
}

func ExampleSVGPathNib2() {
	cbox := RoundedBox(40*unitX,30*unitX, 4*unitX)
	p:=new(Path)
	b := NewBrush((*SimpleSvgPathNib)(p))
	cbox.Draw(b)
	fmt.Println(*p)
	// Output:
	// L36,30 40,26 40,-26 36,-30 -36,-30 -40,-26 -40,26 -36,30
}

