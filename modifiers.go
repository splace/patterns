package pattern

import "math"

type Modifier interface {
	modify(x, y x) (x, x)
}

// a Limited translated
type Translated struct {
	Limited
	X, Y x
}

func (p Translated) at(x, y x) y {
	return p.Limited.at(p.modify(x, y))
}

func (p Translated) modify(x, y x) (x, x) {
	return x - p.X, y - p.Y
}

func (p Translated) MaxX() x {
	return p.Limited.MaxX() + max2(p.X, p.Y)
}

// a Unlimited translated
type UnlimitedTranslated struct {
	Unlimited
	X, Y x
}

func (p UnlimitedTranslated) at(x, y x) y {
	return p.Unlimited.at(p.modify(x, y))
}

func (p UnlimitedTranslated) modify(x, y x) (x, x) {
	return x - p.X, y - p.Y
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

// a Limited Scaled
type Reduced struct {
	Limited
	X, Y float32
}

func (p Reduced) at(x, y x) y {
	return p.Limited.at(p.modify(x, y))
}

func (p Reduced) modify(px, py x) (x, x) {
	return x(float32(px) * p.X), x(float32(py) * p.Y)
}

func (p Reduced) MaxX() x {
	if p.Y > p.X {
		return x(float32(p.Limited.MaxX()) / p.X)
	}
	return x(float32(p.Limited.MaxX()) / p.Y)
}

func NewFitted(p Limited, h, w float32) Limited {
	return Reduced{p, 2 / h, 2 / w}
}

// a Unlimited Scaled
type UnlimitedReduced struct {
	Unlimited
	X, Y float32
}

func (p UnlimitedReduced) at(px, py x) y {
	return p.Unlimited.at(p.modify(px, py))
}

func (p UnlimitedReduced) modify(px, py x) (x, x) {
	return x(float32(px) * p.X), x(float32(py) * p.Y)
}

// a Limited Zoomed
type Shrunk struct {
	Limited
	Factor float32
}

func (p Shrunk) at(px, py x) y {
	return p.Limited.at(p.modify(px, py))
}

func (p Shrunk) modify(px, py x) (x, x) {
	return x(float32(px) * p.Factor), x(float32(py) * p.Factor)
}

func (p Shrunk) MaxX() x {
	return x(float32(p.Limited.MaxX()) / p.Factor)
}

// a Unlimited Zoomed
type UnlimitedShrunk struct {
	Unlimited
	Factor float32
}

func (p UnlimitedShrunk) at(px, py x) y {
	return p.Unlimited.at(p.modify(px, py))
}

func (p UnlimitedShrunk) modify(px, py x) (x, x) {
	return x(float32(px) * p.Factor), x(float32(py) * p.Factor)
}

// a Limited Rotated
type Rotated struct {
	Limited
	sinA, cosA float64
}

func NewRotated(p Limited, a float64) Limited {
	s, c := math.Sincos(a)
	return Rotated{p, s, c}
}

func (p Rotated) at(px, py x) y {
	return p.Limited.at(p.modify(px, py))
}

func (p Rotated) modify(px, py x) (x, x) {
	return x(float64(px)*p.cosA - float64(py)*p.sinA), x(float64(px)*p.sinA + float64(py)*p.cosA)
}

func (p Rotated) MaxX() x {
	return p.Limited.MaxX() * x(math.Abs(p.sinA)+math.Abs(p.cosA))
}

// a Unlimited Rotated
type UnlimitedRotated struct {
	Unlimited
	sinA, cosA float64
}

func (p UnlimitedRotated) at(px, py x) y {
	return p.Unlimited.at(p.modify(px, py))
}

func (p UnlimitedRotated) modify(px, py x) (x, x) {
	return x(float64(px)*p.cosA - float64(py)*p.sinA), x(float64(px)*p.sinA + float64(py)*p.cosA)
}

func NewUnlimitedRotated(p Unlimited, a float64) Unlimited {
	s, c := math.Sincos(a)
	return UnlimitedRotated{p, s, c}
}

// a Limited reversed
type Inverted struct {
	Limited
}

func (p Inverted) at(px, py x) (v y) {
	if p.Limited.at(px, py) == unitY {
		return
	}
	return unitY
}

// a Unlimited reversed
type UnlimitedInverted struct {
	Unlimited
}

func (p UnlimitedInverted) at(px, py x) (v y) {
	if p.Unlimited.at(px, py) == unitY {
		return
	}
	return unitY
}

type Limiter struct {
	Unlimited
	Max x
}

func (p Limiter) MaxX() x {
	return p.Max
}

// replaces Limiter{Unlimited(c),c.MaxX()} to enable simpler sniffing
type CachedMaxX struct{
	Limited
	max x
}

func (f CachedMaxX) MaxX() (max x) {
	return f.max
}

func NewCachedMaxX(l Limited) Limited{
	return CachedMaxX{l,l.MaxX()}
}



// replaces Limiter{Unlimited(c),c.MaxX()} to enable simpler sniffing
type CachedMaxXComposite struct{
	Composite
	max x
}

func (f CachedMaxXComposite) MaxX() (max x) {
	return f.max
}

func NewCachedMaxXComposite(c Composite) Limited{
	return CachedMaxXComposite{c,c.MaxX()}
}


