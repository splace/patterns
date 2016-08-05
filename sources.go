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
	Radius x
	Filling
}

func (p Disc) at(px, py x) (v y) {
	if px*px+py*py <= p.Radius*p.Radius {
		return p.Fill
	}
	return
}

func (p Disc) MaxX() x {
	return p.Radius
}

type Square struct {
	Extent x
	Filling
}

func (p Square) at(px, py x) (v y) {
	if py < p.Extent && py >= -p.Extent && px >= -p.Extent && px < p.Extent {
		return p.Fill
	}
	return
}

func (p Square) MaxX() x {
	return p.Extent
}

func NewBox(Extent, Width x, f Filling) LimitedPattern {
	return Limiter{Composite{Square{Extent - Width, f}, Inverted{Square{Extent + Width, f}}}, Extent + Width}
}
