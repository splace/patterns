package pattern

import "math"

type DashedNib struct {
	FacettedNib
	Along, Repeat,Solid x
}

func (n DashedNib) Straight(x1, y1, x2, y2 x) Limited {
	if x1 == x2 && y1 == y2 {
		return nil
	}
	ndx, dy := float64(x1-x2), float64(y2-y1)
	l:=float32(math.Hypot(ndx, dy))/unitX
	// TODO internally using Reduced results in a smaller MaxX.
	return Translated{NewRotated(NewFitted(Square(Filling(n.In)), l, float32(n.Width)/unitX), math.Atan2(dy, ndx)), (x1 + x2) >> 1, (y1 + y2) >> 1}
}
