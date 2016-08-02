package patterns
import (
	//"math"
	//"fmt"
)

// brush holds state, style/position , for line based, possibly region outlinng
type Brush struct {
	Width                float32
	in               	y
	x, y                 x
	sx, sy               x
	Relative             bool
}

func (p Brush) Line(px1, py1, px2, py2 x) LimitedPattern {
	return Translated{Square{px2-px1,true},px2-px1, py2-py1}
}
