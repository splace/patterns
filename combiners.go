package patterns

import (
	"encoding/gob"
)

func init() {
	gob.Register(Composite{})
}

// Pattern thats composed from layered Patterns
// without transparency this is that same as Stack,
type Composite []Pattern

func (c Composite) at(px, py x) (total y) {
	for _, p := range c {
		if lp, ok := p.(LimitedPattern); ok {
			m := lp.maxX()
			if px > m || py > m || px < -m || py < -m {
				continue
			}
		}
		total = composit(total, p.at(px, py))
		if total.isOpaque() {
			return
		}
	}
	return
}

// helper to enable generation from another slice.
func NewComposite(ps ...Pattern) Composite {
	return Composite(ps)
}
