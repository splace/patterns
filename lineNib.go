package patterns

import "math"

// Nib where curves are returned as straight lines
type LineNib struct {
	Width x
	In    y
}

func (p LineNib) Straight(x1, y1, x2, y2 x) LimitedPattern {
	if x1==x2 && y1==y2 {return nil}
	ndx, dy := float64(x1-x2), float64(y2-y1)
	// NewRotated actually returns a LimitedPattern (as a Pattern) because NewLine returns one, so assert can never fail.
	// TODO internally using Reduced results in a smaller MaxX.
	return Translated{NewRotated(Reduced{Square(Filling(p.In)), float32(unitX * 2 / math.Hypot(ndx, dy)), float32(unitX * 2 / float64(p.Width))}, math.Atan2(dy, ndx)).(LimitedPattern), (x1 + x2) >> 1, (y1 + y2) >> 1}

	//	return Translated{NewRotated(Rectangle(x(math.Hypot(ndx,dy)),p.Width, Filling(p.In)),math.Atan2(dy,ndx)).(LimitedPattern),(x1+x2)>>1, (y1+y2)>>1}
}

func (p LineNib) Curved(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) LimitedPattern {
	return p.Straight(sx, sy, ex, ey)
}

func (p LineNib) SimpleCurved(sx, sy, c1x, c1y, ex, ey x) LimitedPattern {
	return p.Straight(sx, sy, ex, ey)
}

func (p LineNib) Conic(sx, sy, rx, ry x, a float64, large, sweep bool, ex, ey x) LimitedPattern {
	return p.Straight(sx, sy, ex, ey)
}


func (f LineNib) Box(x, y x) LimitedPattern {
	return Limiter{Composite{f.Straight(-x, y, x, y), f.Straight(x, y, x, -y), f.Straight(x, -y, -x, -y), f.Straight(-x, -y, -x, y)}, max(x+f.Width, y+f.Width)}
}

func (f LineNib) Polygon(coords ...[2]x) LimitedPattern {
	s := make([]Pattern, len(coords))
	l := Limits{coords[0][0], coords[0][1], coords[0][0], coords[0][1]}
	for i := 1; i < len(s); i++ {
		s[i-1] = f.Straight(coords[i-1][0], coords[i-1][1], coords[i][0], coords[i][1])
		l.Include([2]x{coords[i][0], coords[i][1]})
	}
	s[len(coords)-1] = f.Straight(coords[len(coords)-1][0], coords[len(coords)-1][1], coords[0][0], coords[0][1])
	return Translated{Limiter{UnlimitedTranslated{NewComposite(s...), (l.MaxX + l.MinX) >> 1, (l.MaxY + l.MinY) >> 1}, max((l.MaxX-l.MinX)>>1, (l.MaxY-l.MinY)>>1) + f.Width}, -((l.MaxX + l.MinX) >> 1), -((l.MaxY + l.MinY) >> 1)}
}

// Limits hold max and min points
type Limits struct {
	MinX, MinY, MaxX, MaxY x
}

//  expanded Limits to include new points.
func (d *Limits) Include(p [2]x) {
	if p[0] < d.MinX {
		d.MinX = p[0]
	} else {
		if p[0] > d.MaxX {
			d.MaxX = p[0]
		}
	}
	if p[1] < d.MinY {
		d.MinY = p[1]
	} else {
		if p[1] > d.MaxY {
			d.MaxY = p[1]
		}
	}
}
