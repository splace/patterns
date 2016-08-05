package patterns

import (
	//	"math"
	"encoding/gob"
)

func init() {
	gob.Register(Constant{})
}

type source interface {
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

func (p Disc) at(px, py x) (v y) {
	if px*px+py*py <= 1 {
		return p.Fill
	}
	return
}

func (p Disc) MaxX() x {
	return 1
}

type Square struct {
	Filling
}

func (p Square) at(px, py x) (v y) {
	if py < 1 && py >= -1 && px >= -1 && px < 1 {
		return p.Fill
	}
	return
}

func (p Square) MaxX() x {
	return 1
}

func NewBox(Extent, Width x, f Filling) LimitedPattern {
	return Limiter{Inverted{Composite{Shrunk{Square{f}, 1/float32(Extent - Width)}, Inverted{Shrunk{Square{f}, 1/float32(Extent + Width)}}}}, Extent + Width}
}
