package patterns

import (
	//	"math"
	"encoding/gob"
)

func init() {
	gob.Register(Constant{})
}

type filler interface {
	fill() y
}

type Filling struct {
	Fill y
}

func (s Filling) fill() y {
	return s.Fill
}

// a Pattern with constant value
type Constant struct {
	Filling
}

func (p Constant) at(px, py x) y {
	return p.Fill
}

type Disc struct {
	Filling
}

const unitX2 = int64(unitX)*int64(unitX)

func (p Disc) at(px, py x) (v y) {
	x2:=int64(px)
	y2:=int64(py)
	if x2*x2+y2*y2 <= unitX2 {
		return p.Fill
	}
	return
}

func (p Disc) MaxX() x {
	return unitX
}

type Square struct {
	Filling
}

const unitXm = -unitX

func (p Square) at(px, py x) (v y) {
	if py < unitX && py >= unitXm && px >= unitXm && px < unitX {
		return p.Fill
	}
	return
}

func (p Square) MaxX() x {
	return unitX
}

func NewBox(Extent, Width float32, f Filling) LimitedPattern {
	return Limiter{UnlimitedInverted{Composite{Shrunk{Square{f}, 1/(Extent - Width/2)}, UnlimitedInverted{Shrunk{Square{f}, 1/(Extent + Width/2)}}}}, X(Extent + Width)}
}


