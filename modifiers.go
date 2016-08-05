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
	p      Pattern
	dx, dy x
}

func (p Shifted) at(px, py x) y {
	return p.p.at(px-p.dx, py-p.dy)
}

// a LimitedPattern translated
type Translated struct {
	p      LimitedPattern
	dx, dy x
}

func (p Translated) at(px, py x) y {
	return p.p.at(px-p.dx, py-p.dy)
}

func (p Translated) MaxX() x {
	return p.p.MaxX() + max4(p.dx, -p.dx, p.dy, -p.dy)
}

func max4(a, b, c, d x) (max x) {
	max=a
	switch {
	case b>max:
		max=b
		fallthrough
	case c>max:
		max=c
		fallthrough
	case d>max:
		return d
	}
	return max
}

// a Pattern Scaled
type Scaled struct {
	p      Pattern
	sx, sy float32
}

func (p Scaled) at(px, py x) y {
	return p.p.at(x(float32(px)*p.sx), x(float32(py)*p.sy))
}

// a Pattern Rotated
type Rotated struct {
	p          Pattern
	sinA, cosA float64
}

func (p Rotated) at(px, py x) y {
	return p.p.at(x(float64(px)*p.cosA-float64(py)*p.sinA), x(float64(px)*p.sinA+float64(py)*p.cosA))
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

func (p Limiter) at(px, py x) (v y) {
	if p.Pattern.at(px, py) == unitY {
		return
	}
	return unitY
}

func (p Limiter) MaxX() x {
	return p.Extent
}
