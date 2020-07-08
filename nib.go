package patterns

// Nibs return LimitedPatterns when asked for lines and/or curves.
type Nib interface {
	Straight(x, x, x, x) LimitedPattern
	Curved(x, x, x, x, x, x, x, x) LimitedPattern
	Conic(x, x, x, x, float64, bool, bool, x, x) LimitedPattern
}
