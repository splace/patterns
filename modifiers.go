package patterns

import (
	"encoding/gob"
	"math"
)

func init() {
	gob.Register(Translated{})
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

// a Pattern shifted
type UnlimitedTranslated struct {
	Pattern
	dx, dy x
}

func (p UnlimitedTranslated) at(px, py x) y {
	return p.Pattern.at(px-p.dx, py-p.dy)
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
	return
}

// a LimitedPattern Scaled
type Reduced struct {
	LimitedPattern
	sx, sy float32
}

func (p Reduced) at(px, py x) y {
	return p.LimitedPattern.at(x(float32(px)*p.sx), x(float32(py)*p.sy))
}

func (p Reduced) MaxX() x {
	if p.sx>p.sy{
		return x(float32(p.LimitedPattern.MaxX())/p.sy)
	}
	return x(float32(p.LimitedPattern.MaxX()) / p.sx)
}


// a Pattern Scaled
type UnlimitedReduced struct {
	Pattern
	sx, sy float32
}

func (p UnlimitedReduced) at(px, py x) y {
	return p.Pattern.at(x(float32(px)*p.sx), x(float32(py)*p.sy))
}

// a LimitedPattern Zoomed
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

// a Pattern Zoomed
type UnlimitedShrunk struct {
	Pattern
	zx float32
}

func (p UnlimitedShrunk) at(px, py x) y {
	return p.Pattern.at(x(float32(px)*p.zx), x(float32(py)*p.zx))
}

// a LimitedPattern Rotated
type Rotated struct {
	LimitedPattern
	sinA, cosA float64
}

func (p Rotated) at(px, py x) y {
	return p.LimitedPattern.at(x(float64(px)*p.cosA-float64(py)*p.sinA), x(float64(px)*p.sinA+float64(py)*p.cosA))
}

func max4float64(a, b, c, d float64) (max float64) {
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
	return
}

func (p Rotated) MaxX() x {
	return x(float64(p.LimitedPattern.MaxX())/max4float64(p.sinA,p.cosA,-p.sinA,-p.cosA))
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
	if lp,ok:=p.(LimitedPattern);ok{
		return Rotated{lp, math.Sin(a), math.Cos(a)}
	}
	return UnlimitedRotated{p, math.Sin(a), math.Cos(a)}
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
	Extent x
}

func (p Limiter) MaxX() x {
	return p.Extent
}



