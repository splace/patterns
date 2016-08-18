package patterns

import (
	"math"
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
	length := float32(math.Sqrt(float64(px2-px1)*float64(px2-px1) + float64(py2-py1)*float64(py2-py1)))
	
	return Translated{NewRotated(Reduced{Square{Filling{p.In}}, float32(unitX*2)/length,float32(unitX*2)/float32(p.Width) },math.Atan2(float64(px1-px2),float64(py2-py1)) ).(LimitedPattern),(px1+px2)/2, (py1+py2)/2}
}



