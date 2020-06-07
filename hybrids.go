package patterns

func Rectangle(h, w x, f Filling) LimitedPattern {
	return Reduced{Square{f}, float32(unitX*2)/float32(h),float32(unitX*2)/float32(w)}
}

// nice for testing shows region not in pattern AND inside extended limits
func NewBorderedInverse(p LimitedPattern, border x) LimitedPattern {
	return Inverted{Limiter{Inverted{p},p.MaxX()+border}}
}


func NewFrame(Extent, Width float32, f Filling) LimitedPattern {
	return Limiter{UnlimitedInverted{Composite{Shrunk{Square{f}, 1/(Extent - Width/2)}, UnlimitedInverted{Shrunk{Square{f}, 1/(Extent + Width/2)}}}}, X(Extent + Width)}
}

