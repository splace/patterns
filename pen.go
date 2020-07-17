package pattern

// Pens have methods to create Limited depending on, and maintaining, a current location.
// Optionally it adds a Limited at the joins between drawn, not just moved, segments.
type Pen struct {
	Nib
	x, y            x
	Joiner          Limited
	joinerNotNeeded bool
}

func (p *Pen) ComposeJoiner(l Limited) Limited {
	if p.joinerNotNeeded || p.Joiner == nil {
		p.joinerNotNeeded = false
		return l
	}
	if l == nil {
		return Translated{p.Joiner, p.x, p.y}
	}
	c:=Composite{l, Translated{p.Joiner, p.x, p.y}}
//	return NewCachedMaxX(Composite{l, Translated{p.Joiner, p.x, p.y}})
//	return Limiter{Composite{l, Translated{p.Joiner, p.x, p.y}}, l.MaxX() + p.Joiner.MaxX()}
//	return Composite{l, Translated{p.Joiner, p.x, p.y}}
//	c:=Composite{l, Translated{p.Joiner, p.x, p.y}}
	return Limiter{Unlimited(c),c.MaxX()}
}

func (p *Pen) MoveTo(x, y x) {
	p.x, p.y = x, y
	p.joinerNotNeeded = true
	return
}

func (p *Pen) LineTo(x, y x) (l Limited) {
	l = p.ComposeJoiner(p.Straight(p.x, p.y, x, y))
	p.x, p.y = x, y
	return
}

func (p *Pen) LineToVertical(y x) (l Limited) {
	l = p.ComposeJoiner(p.Straight(p.x, p.y, p.x, y))
	p.y = y
	return
}

func (p *Pen) LineToHorizontal(x x) (l Limited) {
	l = p.ComposeJoiner(p.Straight(p.x, p.y, x, p.y))
	p.x = x
	return
}

func (p *Pen) ArcTo(rx, ry x, a float64, large, sweep bool, x, y x) (l Limited) {
	l = p.ComposeJoiner(p.Conic(p.x, p.y, rx, ry, a, large, sweep, x, y))
	p.x, p.y = x, y
	return
}

func (p *Pen) QuadraticBezierTo(cx, cy, x, y x) (l Limited) {
	l = p.ComposeJoiner(p.Curved(p.x, p.y, cx, cy, cx, cy, x, y))
	p.x, p.y = x, y
	return
}

func (p *Pen) CubicBezierTo(c1x, c1y, c2x, c2y, x, y x) (l Limited) {
	l = p.ComposeJoiner(p.Curved(p.x, p.y, c1x, c1y, c2x, c2y, x, y))
	p.x, p.y = x, y
	return
}

// PenPath's have methods to create Limited depending on a Pen and if that pen is continuously drawing without gaps, or not.
type PenPath struct {
	Pen
	x, y x
}

func (p *PenPath) MoveTo(px, py x) (l Limited) {
	p.x, p.y = px, py
	p.Pen.MoveTo(px, py)
	return
}

func (p *PenPath) LineClose() (l Limited) {
	//if p.Pen.x==p.x && p.Pen.y==p.y {return}
	l = p.ComposeJoiner(p.LineTo(p.x, p.y))
	return
}
