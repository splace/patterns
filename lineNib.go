package patterns

import "math"

// Nib where curves are returned as straight lines
type LineNib struct {
	Width x
	In    y
}

func (p LineNib) Straight(x1, y1, x2, y2 x) LimitedPattern {
	ndx, dy := float64(x1-x2), float64(y2-y1)
	// NewRotated actually returns a LimitedPattern (as a Pattern) because NewLine returns one, so assert can never fail.
	// TODO internally using Reduced results in a smaller MaxX.
//	if float32(unitX * 2 / math.Hypot(ndx, dy))==float32(math.Inf(+1)){
//		return Limiter{Constant(false),x(0)}
//	}
	return Translated{NewRotated(Reduced{Square(Filling(p.In)), float32(unitX * 2 / math.Hypot(ndx, dy)), float32(unitX * 2 / float64(p.Width))}, math.Atan2(dy, ndx)).(LimitedPattern), (x1 + x2) >> 1, (y1 + y2) >> 1}

	//	return Translated{NewRotated(Rectangle(x(math.Hypot(ndx,dy)),p.Width, Filling(p.In)),math.Atan2(dy,ndx)).(LimitedPattern),(x1+x2)>>1, (y1+y2)>>1}
}

func (p LineNib) Curved(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) LimitedPattern {
	return p.Straight(sx, sy, ex, ey)
}

func (p LineNib) Conic(sx, sy, rx, ry x, a float64, large, sweep bool, ex, ey x) LimitedPattern {
	return p.Straight(sx, sy, ex, ey)
}
