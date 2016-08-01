package patterns

import (
	"fmt"
	"image/color"
)

// satisfying the Signal interface means a type represents an analogue signal, where a y property varies with an x parameter.
type Pattern interface {
	property(x,x) color.Color
}

// the x represents a value from -infinity to +infinity, but is actually limited by its current underlying representation.
type x int64 // current underlying representation
const xBits=64
const unitX = x(1)

// string representation of an x scaled to unitX
func (p x) String() string {
	return fmt.Sprintf("%9.2f", float32(p)/float32(unitX))
}

