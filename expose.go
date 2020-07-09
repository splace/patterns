package patterns

import "math"
import "fmt"

// with internal representation/scale of x values hidden, need to be able to convert to it, 1 -> UnitX
func X(d interface{}) x {
	return MultiplyX(d, unitX)
}

// multiply anything by an x quantity
func MultiplyX(m interface{}, d x) x {
	switch mt := m.(type) {
	case int:
		return d * x(mt)
	case uint:
		return d * x(mt)
	case int8:
		return d * x(mt)
	case uint8:
		return d * x(mt)
	case int16:
		return d * x(mt)
	case uint16:
		return d * x(mt)
	case int32:
		return d * x(mt)
	case uint32:
		return d * x(mt)
	case int64:
		return d * x(mt)
	case uint64:
		return d * x(mt)
	case float32:
		if math.IsNaN(float64(mt)) || math.IsInf(float64(mt),0) {
			panic("Unable to convert:"+fmt.Sprint(mt))
		}
		return x(float32(d)*mt + .5)
	case float64:
		if math.IsNaN(mt) || math.IsInf(mt,0) {
			panic("Unable to convert:"+fmt.Sprint(mt))
		}
		return x(float64(d)*mt + .5)
	default:
		return d
	}
}
