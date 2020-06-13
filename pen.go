package patterns


// Nibs can create LimitedPatterns from lines and curves.
type Nib interface{
	Line(x,x,x,x) LimitedPattern
	QuadraticBezier(x,x,x,x,x,x) LimitedPattern
	CubicBezier(x,x,x,x,x,x,x,x) LimitedPattern
	Arc(x,x,x,x,float64,bool,bool,x,x) LimitedPattern
}

// Pens have methods to create LimitedPatterns relative, or not, to their current position.
type Pen struct {
	Nib
	Relative bool
	x, y     x
}


func (p *Pen) MoveTo(px, py x) {
	if p.Relative {
		p.x += px
		p.y += py
		return
	}
	p.x, p.y= px, py
	return
}

func (p *Pen) LineTo(px, py x) LimitedPattern {
	if p.Relative {
		px += p.x
		py += p.y
	}
	s := p.Line(p.x, p.y, px, py)
	p.x, p.y = px, py
	return s
}


func (p *Pen) LineToVertical(py x) LimitedPattern {
	if p.Relative {
		py += p.y
	}
	s := p.Line(p.x, p.y, p.x, py)
	p.y = py
	return s
}

func (p *Pen) LineToHorizontal(px x) LimitedPattern {
	if p.Relative {
		px += p.x
	}
	s := p.Line(p.x, p.y, px, p.y)
	p.x = px
	return s
}


func (p *Pen) ArcTo(rx,ry x, a float64, large,sweep bool,x,y x) LimitedPattern {
	if p.Relative {
		x += p.x
		y += p.y
	}
	s := p.Arc(p.x,p.y,rx,ry,a,large,sweep,x,y)
	p.x, p.y=x,y
	return s
}

func (p *Pen) QuadraticBezierTo(cx,cy,px,py x) LimitedPattern {
	if p.Relative {
		px += p.x
		py += p.y
		cx += p.x
		cy += p.y
	}
	s := p.QuadraticBezier(p.x,p.y,cx,cy,px, py)
	p.x, p.y=px,py
	return s
}

func (p *Pen) CubicBezierTo(c1x,c1y,c2x,c2y,px,py x) LimitedPattern {
	if p.Relative {
		px += p.x
		py += p.y
		c1x += p.x
		c1y += p.y
		c2x += p.x
		c2y += p.y
	}
	s := p.CubicBezier(p.x,p.y,c1x,c1y,c2x,c2y,px, py)
	p.x, p.y=px,py
	return s
}



// PenPath have methods to create LimitedPatterns using a Pen and the loops start.
type PenPath struct{
	Pen
	x,y x
}

func (p *PenPath) MoveTo(px, py x) {
	p.Pen.MoveTo(px,py)
	p.x, p.y= px, py
	return
}

func (p *PenPath) StartLine(x1, y1, x2, y2 x) LimitedPattern {
	p.MoveTo(x1,y1)
	return p.LineTo(x2, y2)
}

func (p *PenPath) LineClose() LimitedPattern {
	s := p.Line(p.Pen.x,p.Pen.y,p.x, p.y)
	p.x, p.y = p.Pen.x,p.Pen.y
	return s
}

