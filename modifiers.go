package patterns

import "math"

// a LimitedPattern translated
type Translated struct {
	LimitedPattern
	X, Y x
}

func (p Translated) at(x, y x) y {
	return p.LimitedPattern.at(x-p.X, y-p.Y)
}

func (p Translated) MaxX() x {
	return p.LimitedPattern.MaxX() + max2(p.X, p.Y)
}

// a Pattern translated
type UnlimitedTranslated struct {
	Pattern
	X, Y x
}

func (p UnlimitedTranslated) at(x, y x) y {
	return p.Pattern.at(x-p.X, y-p.Y)
}

func abs(a x) x {
	if a < 0 {
		return -a
	}
	return a
}
func max(a, b x) x {
	if b > a {
		a = b
	}
	return a
}

func max2(a, b x) x {
	return max(abs(a), abs(b))
}

func max4(a, b, c, d x) x {
	return max2(max2(a, b), max2(c, d))
}

func max6(a, b, c, d, e, f x) x {
	return max2(max4(a, b, c, d), max2(e, f))
}

func max8(a, b, c, d, e, f, g, h x) x {
	return max2(max4(a, b, c, d), max4(e, f, g, h))
}

func max10(a, b, c, d, e, f, g, h, i, j x) x {
	return max2(max8(a, b, c, d, e, f, g, h), max2(i, j))
}

// a LimitedPattern Scaled
type Reduced struct {
	LimitedPattern
	X, Y float32
}

func (p Reduced) at(px, py x) y {
	return p.LimitedPattern.at(x(float32(px)*p.X), x(float32(py)*p.Y))
}

func (p Reduced) MaxX() x {
	if p.Y > p.X {
		return x(float32(p.LimitedPattern.MaxX()) / p.X)
	}
	return x(float32(p.LimitedPattern.MaxX()) / p.Y)
}

// a Pattern Scaled
type UnlimitedReduced struct {
	Pattern
	X, Y float32
}

func (p UnlimitedReduced) at(px, py x) y {
	return p.Pattern.at(x(float32(px)*p.X), x(float32(py)*p.Y))
}

// a LimitedPattern Zoomed
type Shrunk struct {
	LimitedPattern
	Factor float32
}

func (p Shrunk) at(px, py x) y {
	return p.LimitedPattern.at(x(float32(px)*p.Factor), x(float32(py)*p.Factor))
}

func (p Shrunk) MaxX() x {
	return x(float32(p.LimitedPattern.MaxX()) / p.Factor)
}

// a Pattern Zoomed
type UnlimitedShrunk struct {
	Pattern
	Factor float32
}

func (p UnlimitedShrunk) at(px, py x) y {
	return p.Pattern.at(x(float32(px)*p.Factor), x(float32(py)*p.Factor))
}

// a LimitedPattern Rotated
type Rotated struct {
	LimitedPattern
	sinA, cosA float64
}

func (p Rotated) at(px, py x) y {
	return p.LimitedPattern.at(x(float64(px)*p.cosA-float64(py)*p.sinA), x(float64(px)*p.sinA+float64(py)*p.cosA))
}

func (p Rotated) MaxX() x {
	return p.LimitedPattern.MaxX() * x(math.Abs(p.sinA)+math.Abs(p.cosA))
}

// a Pattern Rotated
type UnlimitedRotated struct {
	Pattern
	sinA, cosA float64
}

func (p UnlimitedRotated) at(px, py x) y {
	return p.Pattern.at(x(float64(px)*p.cosA-float64(py)*p.sinA), x(float64(px)*p.sinA+float64(py)*p.cosA))
}

func NewRotated(p Pattern, a float64) Pattern {
	s, c := math.Sincos(a)
	if lp, ok := p.(LimitedPattern); ok {
		return Rotated{lp, s, c}
	}
	return UnlimitedRotated{p, s, c}
}

// a LimitedPattern reversed
type Inverted struct {
	LimitedPattern
}

func (p Inverted) at(px, py x) (v y) {
	if p.LimitedPattern.at(px, py) == unitY {
		return
	}
	return unitY
}

// a Pattern reversed
type UnlimitedInverted struct {
	Pattern
}

func (p UnlimitedInverted) at(px, py x) (v y) {
	if p.Pattern.at(px, py) == unitY {
		return
	}
	return unitY
}

type Limiter struct {
	Pattern
	Max x
}

func (p Limiter) MaxX() x {
	return p.Max
}
/* run: args="" Fri 3 Jul 00:26:46 BST 2020 go version go1.14.3 linux/amd64
Fri 3 Jul 00:26:46 BST 2020
*/
