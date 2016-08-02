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

func (p Constant) property(px,py x) y {
	return p.Value
}

type Disc struct {
	Radius x
	Value y
}

func (p Disc) property(px,py x) (v y) {
	if px*px+py*py <= p.Radius*p.Radius {
		return p.Value
	}
	return
}

type Ring struct {
	Radius,Width x
	Value y
}

func (p Ring) property(px,py x) (v y) {
	r:= px*px+py*py
	if r <= (p.Radius+p.Width)*(p.Radius+p.Width) && r>= (p.Radius-p.Width)*(p.Radius-p.Width) {
		return p.Value
	}
	return
}


type Square struct {
	Height,Breadth x
	Value y
}

func (p Square) property(px,py x) (v y) {
	if py <= p.Height && py>= -p.Height && px >= -p.Breadth && px <= p.Breadth {
		return p.Value
	}
	return
}

