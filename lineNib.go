package patterns

import "math"

// Nib with unit width where curves are just drawn as straight lines.
type LineNib struct{}

func (p LineNib) Line(x1, y1, x2, y2 x) LimitedPattern {
	ndx,dy:=float64(x1-x2),float64(y2-y1)
	// NewRotated actually returns a LimitedPattern (as a Pattern) because NewLine returns one, so assert can never fail.
	// TODO could reduce MaxX since we know better than worst case used by rotate.
	return Translated{NewRotated(Rectangle(x(math.Hypot(ndx,dy)),unitX, Filling(unitY)),math.Atan2(dy,ndx)).(LimitedPattern),(x1+x2)>>1, (y1+y2)>>1}
}

func (p LineNib)	QuadraticBezier(sx, sy, cx, cy, ex, ey x) LimitedPattern{
	return p.Line(sx,sy,ex,ey)
}

func (p LineNib)	CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey x) LimitedPattern{
	return p.Line(sx,sy,ex,ey)
}

func (p LineNib)	Arc(sx,sy,rx,ry x, a float64, large,sweep bool, ex,ey x) LimitedPattern{
	return p.Line(sx,sy,ex,ey)
}
