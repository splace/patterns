package patterns

import (
	"math"
	//"fmt"
)

// brush holds state, style/position , for line based patterns
type Brush struct {
	Width    x
	In       y
	Relative bool
	x, y     x
	sx, sy   x
}

func (p Brush) Line(px1, py1, px2, py2 x) LimitedPattern {
	length := x(math.Sqrt(float64((px2-px1)*(px2-px1) + (py2-py1)*(py2-py1))))
	return Translated{Shrunk{Square{Filling{p.In}}, 1/float32(px2 - px1)}, length, p.Width}
}
