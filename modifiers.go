package patterns

import "math"

// a LimitedPattern translated
type Translated struct {
	LimitedPattern
	x, y x
}

func (p Translated) at(x, y x) y {
	return p.LimitedPattern.at(x-p.x, y-p.y)
}

func (p Translated) MaxX() x {
	return p.LimitedPattern.MaxX() + max2(p.x, p.y)
}

// a Pattern translated
type UnlimitedTranslated struct {
	Pattern
	x, y x
}

func (p UnlimitedTranslated) at(x, y x) y {
	return p.Pattern.at(x-p.x, y-p.y)
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
	x, y float32
}

func (p Reduced) at(px, py x) y {
	return p.LimitedPattern.at(x(float32(px)*p.x), x(float32(py)*p.y))
}

func (p Reduced) MaxX() x {
	if p.y > p.x {
		return x(float32(p.LimitedPattern.MaxX()) / p.x)
	}
	return x(float32(p.LimitedPattern.MaxX()) / p.y)
}

// a Pattern Scaled
type UnlimitedReduced struct {
	Pattern
	x, y float32
}

func (p UnlimitedReduced) at(px, py x) y {
	return p.Pattern.at(x(float32(px)*p.x), x(float32(py)*p.y))
}

// a LimitedPattern Zoomed
type Shrunk struct {
	LimitedPattern
	float32
}

func (p Shrunk) at(px, py x) y {
	return p.LimitedPattern.at(x(float32(px)*p.float32), x(float32(py)*p.float32))
}

func (p Shrunk) MaxX() x {
	return x(float32(p.LimitedPattern.MaxX()) / p.float32)
}

// a Pattern Zoomed
type UnlimitedShrunk struct {
	Pattern
	float32
}

func (p UnlimitedShrunk) at(px, py x) y {
	return p.Pattern.at(x(float32(px)*p.float32), x(float32(py)*p.float32))
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
