package pattern

type DashedNib struct {
	Width x
	In    y
	Repeat,Solid x 
}

func (n DashedNib) Straight(x1, y1, x2, y2 x) Limited {
	return LineNib{}.Straight(x1, y1, x2, y2)
}

func (n DashedNib) Curved(sx, sy, c1x, c1y, c2x, c2y, ex, ey x) Limited {
	return n.Straight(sx, sy, ex, ey)
}

func (n DashedNib) SimpleCurved(sx, sy, c1x, c1y, ex, ey x) Limited {
	return n.Straight(sx, sy, ex, ey)
}

func (n DashedNib) Conic(sx, sy, rx, ry x, a float64, large, sweep bool, ex, ey x) Limited {
	return n.Straight(sx, sy, ex, ey)
}
