package pattern

// Nibs return Limited when asked for lines and/or curves.
type Nib interface {
	Straight(x, x, x, x) Limited
	SimpleCurved(x, x, x, x, x, x) Limited
	Curved(x, x, x, x, x, x, x, x) Limited
	Conic(x, x, x, x, float64, bool, bool, x, x) Limited
}
