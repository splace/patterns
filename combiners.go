package patterns


// Pattern that is made from overlaying Patterns
type Composite []Pattern

func (c Composite) at(px, py x) (total y) {
	for _, p := range c {
		if p==nil{continue}
		if lp, ok := p.(LimitedPattern); ok {
			m := lp.MaxX()
			if px >= m || py >= m || px < -m || py < -m {
				continue
			}
		}
		total = compose(total, p.at(px, py))
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


type LimitedComposite Composite

func (c LimitedComposite) at(px, py x) (total y) {
	return Composite(c).at(px,py)
}


// panics if its first Pattern item istn't a LimitedPattern
func (c LimitedComposite) MaxX() x{
	return Composite(c)[0].(LimitedPattern).MaxX() 
}



