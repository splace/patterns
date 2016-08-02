package patterns
import (
//	"math"
	"encoding/gob"
)

func init() {
	gob.Register(Shifted{})
}

// a Pattern shifted 
type Shifted struct {
	p Pattern
	dx,dy x
}

func (p Shifted) property(px,py x) y {
	return p.p.property(px-p.dx,py-p.dy)
}

// a Pattern shifted 
type Scaled struct {
	p Pattern
	sx,sy float32
}

func (p Scaled) property(px,py x) y {
	return p.p.property(x(float32(px)*p.sx),x(float32(py)*p.sy))
}




