package pattern

import "math"

// Nib where curves are returned as straight lines
type LineNib struct {
	Width x
	In    y
}

func (n LineNib) Straight(x1, y1, x2, y2 x) Limited {
	if x1 == x2 && y1 == y2 {
		return nil
	}
	ndx, dy := float64(x1-x2), float64(y2-y1)
	// TODO internally using Reduced results in a smaller MaxX.
	return Translated{NewRotated(NewFitted(Square(Filling(n.In)), float32(math.Hypot(ndx, dy))/unitX, float32(n.Width)/unitX), math.Atan2(dy, ndx)), (x1 + x2) >> 1, (y1 + y2) >> 1}
}

func (n LineNib) Curved(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) Limited {
	return n.Straight(sx, sy, ex, ey)
}

func (n LineNib) SimpleCurved(sx, sy, c1x, c1y, ex, ey x) Limited {
	return n.Straight(sx, sy, ex, ey)
}

func (n LineNib) Conic(sx, sy, rx, ry x, a float64, large, sweep bool, ex, ey x) Limited {
	return n.Straight(sx, sy, ex, ey)
}

func (l LineNib) Box(x, y x) Limited {
	return Limiter{Composite{l.Straight(-x, y, x, y), l.Straight(x, y, x, -y), l.Straight(x, -y, -x, -y), l.Straight(-x, -y, -x, y)}, max(x+l.Width, y+l.Width)}
}

func (l LineNib) Polygon(coords ...[2]x) Composite {
	c := make(Composite, len(coords))
	//m := Limits{coords[0][0], coords[0][1], coords[0][0], coords[0][1]}
	for i := 1; i < len(c); i++ {
		c[i-1] = l.Straight(coords[i-1][0], coords[i-1][1], coords[i][0], coords[i][1])
		//m.Include([2]x{coords[i][0], coords[i][1]})
	}
	c[len(coords)-1] = l.Straight(coords[len(coords)-1][0], coords[len(coords)-1][1], coords[0][0], coords[0][1])
	// translate - limit - untranslate
	//return Translated{Limiter{UnlimitedTranslated{c, (m.MaxX + m.MinX) >> 1, (m.MaxY + m.MinY) >> 1}, max((m.MaxX-m.MinX)>>1, (m.MaxY-m.MinY)>>1) + l.Width}, -((m.MaxX + m.MinX) >> 1), -((m.MaxY + m.MinY) >> 1)}
	return c
}

//// Limits hold max and min points
//type Limits struct {
//	MinX, MinY, MaxX, MaxY x
//}

////  expanded Limits to include new points.
//func (d *Limits) Include(p [2]x) {
//	if p[0] < d.MinX {
//		d.MinX = p[0]
//	} else {
//		if p[0] > d.MaxX {
//			d.MaxX = p[0]
//		}
//	}
//	if p[1] < d.MinY {
//		d.MinY = p[1]
//	} else {
//		if p[1] > d.MaxY {
//			d.MaxY = p[1]
//		}
//	}
//}
