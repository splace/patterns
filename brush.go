package patterns


// brush holds state, style/position , for line based patterns
// lines produced are independent.
type Brush struct {
	LineBrush
	Relative bool
	x, y     x
	sx, sy   x
}


func (p *Brush) LineTo(px, py x) LimitedPattern {
	if p.Relative {
		px += p.x
		py += p.y
	}
	s := p.Line(p.x, p.y, px, py)
	p.x, p.y = px, py
	return s
}

func (p *Brush) LineToVertical(py x) LimitedPattern {
	if p.Relative {
		py += p.y
	}
	s := p.Line(p.x, p.y, p.x, py)
	p.y = py
	return s
}

func (p *Brush) LineToHorizontal(px x) LimitedPattern {
	if p.Relative {
		px += p.x
	}
	s := p.Line(p.x, p.y, px, p.y)
	p.x = px
	return s
}

func (p *Brush) StartLine(px1, py1, px2, py2 x) LimitedPattern {
	p.MoveTo(px1,py1)
	return p.LineTo(px2, py2)
}

func (p *Brush) MoveTo(px, py x) {
	if p.Relative {
		px += p.x
		py += p.y
	}
	p.x, p.y, p.sx, p.sy = px, py, px, py
	return
}

func (p *Brush) LineClose() LimitedPattern {
	s := p.Line(p.x, p.y, p.sx, p.sy)
	p.x, p.y = p.sx, p.sy
	return s
}

func (p *Brush) ArcTo(rx,ry x, a float64, large,sweep bool,x,y x) LimitedPattern {
	if p.Relative {
		x += p.x
		y += p.y
	}
	s := p.Arc(p.x,p.y,rx,ry,a,large,sweep,x,y)
	p.x, p.y=x,y
	return s
}

func (p *Brush) QuadraticBezierTo(cx,cy,px,py x) LimitedPattern {
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

func (p *Brush) CubicBezierTo(c1x,c1y,c2x,c2y,px,py x) LimitedPattern {
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

