package patterns

// Facetted is a Nib producing curves using a number of straight lines.
// curves are divided according to CurveDivision:  (power of 2 number of divisions.)
// lines are actually drawn by the Straight method of the embedded LineNib, except if a Nib is also embedded then its Straight method is used instead.
type Facetted struct {
	LineNib
	Nib
	CurveDivision uint8
}

func (f Facetted) Straight(sx, sy, ex, ey x) LimitedPattern {
	if f.Nib != nil {
		return f.Nib.Straight(sx, sy, ex, ey)
	}
	return f.LineNib.Straight(sx, sy, ex, ey)
}

func (f Facetted) Curved(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) LimitedPattern {
	return f.CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey)
}

func (f Facetted) SimpleCurved(sx, sy, c1x, c1y, ex, ey x) LimitedPattern {
		return f.QuadraticBezier(sx, sy, c1x, c1y, ex, ey)
}

func (f Facetted) QuadraticBezier(sx, sy, cx, cy, ex, ey x) LimitedPattern {
	return f.polygon(sx, sy, ex, ey, Divide(1<<f.CurveDivision).QuadraticBezier(sx, sy, cx, cy, ex, ey))
}

func (f Facetted) CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) LimitedPattern {
	return f.polygon(sx, sy, ex, ey, Divide(1<<f.CurveDivision).CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey))
}

func (f Facetted) QuinticBezier(sx, sy, c1x, c1y, c2x, c2y, c3x, c3y, ex, ey x) LimitedPattern {
	return f.polygon(sx, sy, ex, ey, Divide(1<<f.CurveDivision).QuinticBezier(sx, sy, c1x, c1y, c2x, c2y, c3x, c3y, ex, ey))
}

func (f Facetted) Conic(sx, sy, rx, ry x, a float64, large, sweep bool, ex, ey x) LimitedPattern {
	return f.polygon(sx, sy, ex, ey, Divide(1<<f.CurveDivision).Arc(sx, sy, rx, ry, a, large, sweep, ex, ey))
}

func (f Facetted) polygon(sx, sy, ex, ey x, pts <-chan [2]x) LimitedPattern {
	var c Composite
	joiner := Shrunk{Disc(Filling(f.In)), 2 * unitX / float32(f.Width)}
	l := Limits{sx, sy, sx, sy}
	for p := range pts {
		c = append(c, f.Straight(sx, sy, p[0], p[1]))
		c = append(c, Translated{&joiner, p[0], p[1]})
		sx, sy = p[0], p[1]
		l.Include(p)
	}
	c = append(c, f.Straight(sx, sy, ex, ey))
	l.Include([2]x{ex, ey})
	return Translated{Limiter{UnlimitedTranslated{c, (l.MaxX + l.MinX) >> 1, (l.MaxY + l.MinY) >> 1}, max((l.MaxX-l.MinX)>>1, (l.MaxY-l.MinY)>>1) + f.Width}, -((l.MaxX + l.MinX) >> 1), -((l.MaxY + l.MinY) >> 1)}
}
