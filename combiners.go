package pattern

// Limited that is made from composing Patterns
type Composite []Unlimited

func (c Composite) at(px, py x) (total y) {
	for i, p := range c {
		if p == nil {
			continue
		}
		if lp, ok := p.(Limited); ok {
			m := lp.MaxX()
			if px >= m || py >= m || px < -m || py < -m {
				continue
			}
		}
		total = compose(total, p.at(px, py))
		if total.isOpaque() {
			// XXX optimisation only for bool y: put this success as first in search for next time
			if i != 0 {
				c[0], c[i] = c[i], c[0]
			}
			return
		}
	}
	return
}

func (c Composite) MaxX() (max x) {
	for _, p := range c {
		if lp, is := p.(Limited); is {
			if m := lp.MaxX(); m > max {
				max = m
			}
		}
	}
	return
}

type UnlimitedComposite Composite

func (c UnlimitedComposite) at(px, py x) (total y) {
	return Composite(c).at(px,py)
}
