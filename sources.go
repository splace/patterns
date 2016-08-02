package patterns

import (
	//	"math"
	"encoding/gob"
)

func init() {
	gob.Register(Constant{})
}

// a Pattern with constant value
type Constant struct {
	Value y
}

func (p Constant) at(px, py x) y {
	return p.Value
}

type Disc struct {
	Radius x
	Value  y
}

func (p Disc) at(px, py x) (v y) {
	if px*px+py*py <= p.Radius*p.Radius {
		return p.Value
	}
	return
}

func (p Disc) maxX() x {
	return p.Radius
}

type Square struct {
	Extent x
	Value  y
}

func (p Square) at(px, py x) (v y) {
	if py <= p.Extent && py >= -p.Extent && px >= -p.Extent && px <= p.Extent {
		return p.Value
	}
	return
}

func (p Square) maxX() x {
	return p.Extent
}
