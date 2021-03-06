package pattern

// nice for testing shows region not in pattern AND inside extended limits
func NewBorderedInverse(p Limited, border x) Limited {
	return Inverted{Limiter{Inverted{p}, p.MaxX() + border}}
}

func NewFrame(Extent, Width float32, f Filling) Limited {
	return Limiter{UnlimitedInverted{Composite{Shrunk{Square(f), 1 / (Extent - Width/2)}, UnlimitedInverted{Shrunk{Square(f), 1 / (Extent + Width/2)}}}}, X(Extent + Width)}
}
