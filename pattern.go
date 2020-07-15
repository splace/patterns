package pattern

import "fmt"

// satisfying the Unlimited interface means a type represents a 2D pattern, where a y property varies with two x parameters.
type Unlimited interface {
	at(x, x) y
}

type Limited interface {
	Unlimited
	MaxX() x
}

// the x represents a value from -infinity to +infinity, but is actually limited by its type.
type x int32

//const xBits = 64
const unitX = 1000

//  x scaled to unitX
//func (p x) String() string {
//	return fmt.Sprintf("%9.2f", float32(p)/float32(unitX))
//}

func (p x) String() string {
	return fmt.Sprint(float32(p) / float32(unitX))
}

// x is scaled as required on scan
// commas are separators by default
func (p *x) Scan(state fmt.ScanState, v rune) (err error) {
	var xscan float32
	state.SkipSpace()
	r, _, err := state.ReadRune()
	if err != nil {
		return
	}
	if r == ',' {
		state.SkipSpace()
	} else {
		state.UnreadRune()
	}
	_, err = fmt.Fscan(state, &xscan)
	if err != nil {
		return
	}
	*p = x(xscan * unitX)
	return
}

// the y type represents a value between +unitY and -unitY.
type y bool

const unitY y = true

// string representation of a y
func (v y) String() string {
	if v {
		return "X"
	}
	return "-"
}

func (v y) isOpaque() bool {
	return bool(v)
}

func compose(y1, y2 y) y {
	return y1 || y2
}

type filler interface {
	fill() y
}

type Filling y

func (f Filling) fill() y {
	return y(f)
}

// copy a Limited leaving out Limited's that are entirely outside a parents Composite's MaxX().
func Optimise(p Limited) Limited {

	return nil
}
