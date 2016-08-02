package patterns

import (
	"fmt"
)

// satisfying the Pattern interface means a type represents a 2D pattern, where a y property varies with two x parameters.
type Pattern interface {
	property(x,x) y
}


// the x represents a value from -infinity to +infinity, but is actually limited by its current underlying representation.
type x int64 // current underlying representation
const xBits=64
const unitX = 1

// string representation of an x scaled to unitX
func (p x) String() string {
	return fmt.Sprintf("%9.2f", float32(p)/float32(unitX))
}

// the y type represents a value between +unitY and -unitY.
type y bool

const unitY y = true
const yBits = 64
const halfyBits = yBits / 2

// string representation of a y, scaled to unitY%
func (v y) String() string {
	if v {return "X"}
	return "-"
}

// transparency
func (v y) Transparency() (t y) {
	if v {return}
	return unitY
}

type LimitedPattern interface{
	Pattern
	MaxX() x
}


