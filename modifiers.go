package patterns

import (
	"encoding/gob"
	"math"
)

func init() {
	gob.Register(Shifted{})
}

// a Pattern shifted
type Shifted struct {
	Pattern
	dx, dy x
}

func (p Shifted) at(px, py x) y {
	return p.Pattern.at(px-p.dx, py-p.dy)
}

// a LimitedPattern translated
type Translated struct {
	LimitedPattern
	dx, dy x
}

func (p Translated) at(px, py x) y {
	return p.LimitedPattern.at(px-p.dx, py-p.dy)
}

func (p Translated) MaxX() x {
	return p.LimitedPattern.MaxX() + max4(p.dx, -p.dx, p.dy, -p.dy)
}

func max4(a, b, c, d x) (max x) {
	max = a
	switch {
	case b > max:
		max = b
		fallthrough
	case c > max:
		max = c
		fallthrough
	case d > max:
		return d
	}
	return max
}

// a Pattern Scaled
type Reduced struct {
	Pattern
	sx, sy float32
}

func (p Reduced) at(px, py x) y {
	return p.Pattern.at(x(float32(px)*p.sx), x(float32(py)*p.sy))
}

// a Pattern Zoomed
type Shrunk struct {
	LimitedPattern
	zx float32
}

func (p Shrunk) at(px, py x) y {
	return p.LimitedPattern.at(x(float32(px)*p.zx), x(float32(py)*p.zx))
}

func (p Shrunk) MaxX() x {
	return x(float32(p.LimitedPattern.MaxX()) / p.zx)
}

// a Pattern Rotated
type Rotated struct {
	Pattern
	sinA, cosA float64
}

func (p Rotated) at(px, py x) y {
	return p.Pattern.at(x(float64(px)*p.cosA-float64(py)*p.sinA), x(float64(px)*p.sinA+float64(py)*p.cosA))
}

func NewRotated(p Pattern, a float64) Pattern {
	return Rotated{p, math.Sin(a), math.Cos(a)}
}

// a Pattern reversed
type Inverted struct {
	Pattern
}

func (p Inverted) at(px, py x) (v y) {
	if p.Pattern.at(px, py) == unitY {
		return
	}
	return unitY
}

// a Pattern reversed
type LimitedInverted struct {
	LimitedPattern
}

func (p LimitedInverted) at(px, py x) (v y) {
	if p.LimitedPattern.at(px, py) == unitY {
		return
	}
	return unitY
}

type Limiter struct {
	Pattern
	Extent x
}

func (p Limiter) MaxX() x {
	return p.Extent
}
