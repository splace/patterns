package patterns


// Nibs can create LimitedPatterns from lines and curves.
type Nib interface{
	Line(x,x,x,x) LimitedPattern
	QuadraticBezier(x,x,x,x,x,x) LimitedPattern
	CubicBezier(x,x,x,x,x,x,x,x) LimitedPattern
	Arc(x,x,x,x,float64,bool,bool,x,x) LimitedPattern
}

// Pens have methods to create LimitedPatterns using their current position.
type Pen struct {
	Nib
	Relative bool
	x, y     x
}

// PenPath have methods to create LimitedPatterns using a Pen and the loops start.
type PenPath struct{
	Pen
	x,y x
}

func (p *PenPath) MoveTo(px, py x) {
	if p.Relative {
		px += p.Pen.x
		py += p.Pen.y
	}
	p.Pen.x, p.Pen.y, p.x, p.y = px, py, px, py
	return
}

func (p *PenPath) LineTo(px, py x) LimitedPattern {
	if p.Relative {
		px += p.Pen.x
		py += p.Pen.y
	}
	s := p.Line(p.Pen.x, p.Pen.y, px, py)
	p.Pen.x, p.Pen.y = px, py
	return s
}

func (p *PenPath) LineToVertical(py x) LimitedPattern {
	if p.Relative {
		py += p.Pen.y
	}
	s := p.Line(p.Pen.x, p.Pen.y, p.Pen.x, py)
	p.Pen.y = py
	return s
}

func (p *PenPath) LineToHorizontal(px x) LimitedPattern {
	if p.Relative {
		px += p.Pen.x
	}
	s := p.Line(p.Pen.x, p.Pen.y, px, p.Pen.y)
	p.Pen.x = px
	return s
}

func (p *PenPath) StartLine(px1, py1, px2, py2 x) LimitedPattern {
	p.MoveTo(px1,py1)
	return p.LineTo(px2, py2)
}

func (p *PenPath) LineClose() LimitedPattern {
	s := p.Line(p.Pen.x, p.Pen.y, p.x, p.y)
	p.x, p.y = p.Pen.x,p.Pen.y
	return s
}

func (p *PenPath) ArcTo(rx,ry x, a float64, large,sweep bool,x,y x) LimitedPattern {
	if p.Relative {
		x += p.Pen.x
		y += p.Pen.y
	}
	s := p.Arc(p.Pen.x,p.Pen.y,rx,ry,a,large,sweep,x,y)
	p.Pen.x, p.Pen.y=x,y
	return s
}

func (p *PenPath) QuadraticBezierTo(cx,cy,px,py x) LimitedPattern {
	if p.Relative {
		px += p.Pen.x
		py += p.Pen.y
		cx += p.Pen.x
		cy += p.Pen.y
	}
	s := p.QuadraticBezier(p.Pen.x,p.Pen.y,cx,cy,px, py)
	p.Pen.x, p.Pen.y=px,py
	return s
}

func (p *PenPath) CubicBezierTo(c1x,c1y,c2x,c2y,px,py x) LimitedPattern {
	if p.Relative {
		px += p.Pen.x
		py += p.Pen.y
		c1x += p.Pen.x
		c1y += p.Pen.y
		c2x += p.Pen.x
		c2y += p.Pen.y
	}
	s := p.CubicBezier(p.Pen.x,p.Pen.y,c1x,c1y,c2x,c2y,px, py)
	p.Pen.x, p.Pen.y=px,py
	return s
}

