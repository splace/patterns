package pattern

// Facetted is a Nib producing curves using a number of straight lines.
// curves are divided according to CurveDivision:  (power of 2 number of divisions.)
// lines are actually drawn by the Straight method of the embedded LineNib, except if a Nib is also embedded then its Straight method is used instead.
type FacettedNib struct {
	LineNib
	Nib
	CurveDivision uint8
}

func (n FacettedNib) Straight(sx, sy, ex, ey x) Limited {
	if n.Nib != nil {
		return n.Nib.Straight(sx, sy, ex, ey)
	}
	return n.LineNib.Straight(sx, sy, ex, ey)
}

func (n FacettedNib) Curved(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) Limited {
	return n.CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey)
}

func (n FacettedNib) SimpleCurved(sx, sy, c1x, c1y, ex, ey x) Limited {
	return n.QuadraticBezier(sx, sy, c1x, c1y, ex, ey)
}

func (n FacettedNib) QuadraticBezier(sx, sy, cx, cy, ex, ey x) Limited {
	return n.polygon(sx, sy, ex, ey, Divide(1<<n.CurveDivision).QuadraticBezier(sx, sy, cx, cy, ex, ey))
}

func (n FacettedNib) CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) Limited {
	return n.polygon(sx, sy, ex, ey, Divide(1<<n.CurveDivision).CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey))
}

func (n FacettedNib) QuinticBezier(sx, sy, c1x, c1y, c2x, c2y, c3x, c3y, ex, ey x) Limited {
	return n.polygon(sx, sy, ex, ey, Divide(1<<n.CurveDivision).QuinticBezier(sx, sy, c1x, c1y, c2x, c2y, c3x, c3y, ex, ey))
}

func (n FacettedNib) Conic(sx, sy, rx, ry x, a float64, large, sweep bool, ex, ey x) Limited {
	return n.polygon(sx, sy, ex, ey, Divide(1<<n.CurveDivision).Arc(sx, sy, rx, ry, a, large, sweep, ex, ey))
}

func (n FacettedNib) polygon(sx, sy, ex, ey x, pts <-chan [2]x) Limited {
	var c Composite
	joiner := Shrunk{Disc(Filling(n.In)), 2 * unitX / float32(n.Width)}
	l := Limits{sx, sy, sx, sy}
	for p := range pts {
		c = append(c, n.Straight(sx, sy, p[0], p[1]), Translated{&joiner, p[0], p[1]})
		sx, sy = p[0], p[1]
		l.Include(p)
	}
	c = append(c, n.Straight(sx, sy, ex, ey))
	l.Include([2]x{ex, ey})
	return Translated{Limiter{UnlimitedTranslated{c, (l.MaxX + l.MinX) >> 1, (l.MaxY + l.MinY) >> 1}, max((l.MaxX-l.MinX)>>1, (l.MaxY-l.MinY)>>1) + n.Width}, -((l.MaxX + l.MinX) >> 1), -((l.MaxY + l.MinY) >> 1)}
}
