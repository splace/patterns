package pattern

import "fmt"

//import "testing"

func ExampleSVGPathNib() {
	cbox := ChamferedRectangle(40*unitX, 30*unitX, 4*unitX)
	p := new(Path)
	b := NewBrush((*SimpleSvgPathNib)(p))
	cbox.Draw(b)
	fmt.Println(CompactStringer(*p))
	//Output:
	//M36,30
	//Q40,30 40,26
	//M40,26
	//L40,-26
	//M40,-26
	//Q40,-30 36,-30
	//M36,-30
	//L-36,-30
	//M-36,-30
	//Q-40,-30 -40,-26
	//M-40,-26
	//L-40,26
	//M-40,26
	//Q-40,30 -36,30
	//M-36,30
	//L36,30
}

func ExampleSVGPathNib2() {
	cbox := RoundedRectangle(40*unitX, 30*unitX, 4*unitX)
	p := new(Path)
	b := NewBrush((*SimpleSvgPathNib)(p))
	cbox.Draw(b)
	fmt.Println(*p)
	// Output:
	// L36,30 40,26 40,-26 36,-30 -36,-30 -40,-26 -40,26 -36,30
}

func ExampleSVGPathLinesOnlyFromRoundedBox() {
	cbox := RoundedRectangle(40*unitX, 30*unitX, 4*unitX)
	p := new(Path)
	b := NewBrush(FacettedNib{Nib: (*SimpleSvgPathNib)(p)})
	cbox.Draw(b)
	fmt.Println(*p)
	// Output:
	// L36,30 40,26 40,-26 36,-30 -36,-30 -40,-26 -40,26 -36,30
}

func ExampleSVGPathNib3() {
	cbox := ChamferedRectangle(40*unitX, 30*unitX, 4*unitX)
	p := new(Path)
	b := NewBrush(FacettedNib{Nib: (*SimpleSvgPathNib)(p), CurveDivision: 3})
	cbox.Draw(b)
	fmt.Println(CompactStringer(*p))
	//Output:
	//M36,30
	//Q40,30 40,26
	//M40,26
	//L40,-26
	//M40,-26
	//Q40,-30 36,-30
	//M36,-30
	//L-36,-30
	//M-36,-30
	//Q-40,-30 -40,-26
	//M-40,-26
	//L-40,26
	//M-40,26
	//Q-40,30 -36,30
	//M-36,30
	//L36,30
}
