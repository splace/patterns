package pattern

// a Path, implementing Nib, that when used appends to itself Drawers such that when, later, calling its Draw method returns the Limited.
type SimpleSvgPathNib Path

func (p *SimpleSvgPathNib) Straight(x1, y1, x2, y2 x) Limited {
	*p = append(*p, MoveTo([]x{x1, y1}), LineTo([]x{x2, y2}))
	return nil
}

func (f *SimpleSvgPathNib) Curved(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) Limited {
	return f.CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey)
}

func (f *SimpleSvgPathNib) SimpleCurved(sx, sy, c1x, c1y, ex, ey x) Limited {
	return f.QuadraticBezier(sx, sy, c1x, c1y, ex, ey)
}

func (p *SimpleSvgPathNib) QuadraticBezier(sx, sy, cx, cy, ex, ey x) Limited {
	*p = append(*p, MoveTo([]x{sx, sy}), QuadraticBezierTo([]x{cx, cy, ex, ey}))
	return nil
}

func (p *SimpleSvgPathNib) CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) Limited {
	*p = append(*p, MoveTo([]x{sx, sy}), CubicBezierTo([]x{c1x, c1y, c2x, c2y, ex, ey}))
	return nil
}

func (p *SimpleSvgPathNib) Conic(sx, sy, rx, ry x, a float64, large, sweep bool, ex, ey x) Limited {
	if large {
		if sweep {
			*p = append(*p, MoveTo([]x{sx, sy}), ArcTo([]x{rx, ry, x(a * unitX), 1, 1, ex, ey}))
		} else {
			*p = append(*p, MoveTo([]x{sx, sy}), ArcTo([]x{rx, ry, x(a * unitX), 1, 0, ex, ey}))
		}
	} else {
		if sweep {
			*p = append(*p, MoveTo([]x{sx, sy}), ArcTo([]x{rx, ry, x(a * unitX), 0, 1, ex, ey}))
		} else {
			*p = append(*p, MoveTo([]x{sx, sy}), ArcTo([]x{rx, ry, x(a * unitX), 0, 0, ex, ey}))
		}
	}
	return nil
}
