package patterns


// Pens have methods to create LimitedPatterns relative to the previously drawn position.
type Pen struct {
	Nib
	x, y     x
}


func (p *Pen) MoveTo(px, py x) {
	p.x, p.y= px, py
	return
}

func (p *Pen) LineTo(px, py x) LimitedPattern {
	s := p.Line(p.x, p.y, px, py)
	p.x, p.y = px, py
	return s
}


func (p *Pen) LineToVertical(py x) LimitedPattern {
	s := p.Line(p.x, p.y, p.x, py)
	p.y = py
	return s
}

func (p *Pen) LineToHorizontal(px x) LimitedPattern {
	s := p.Line(p.x, p.y, px, p.y)
	p.x = px
	return s
}


func (p *Pen) ArcTo(rx,ry x, a float64, large,sweep bool,x,y x) LimitedPattern {
	s := p.Arc(p.x,p.y,rx,ry,a,large,sweep,x,y)
	p.x, p.y=x,y
	return s
}

func (p *Pen) QuadraticBezierTo(cx,cy,px,py x) LimitedPattern {
	s := p.QuadraticBezier(p.x,p.y,cx,cy,px, py)
	p.x, p.y=px,py
	return s
}

func (p *Pen) CubicBezierTo(c1x,c1y,c2x,c2y,px,py x) LimitedPattern {
	s := p.CubicBezier(p.x,p.y,c1x,c1y,c2x,c2y,px, py)
	p.x, p.y=px,py
	return s
}



// PenPath have methods to create LimitedPatterns using a Pen and an origin point.
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

