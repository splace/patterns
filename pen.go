package patterns


// Pens have methods to create LimitedPatterns depending on, and maintaining, its current location.
// Optionally it includes a LimitedPattern at the end of each.
type Pen struct {
	Nib
	Marker LimitedPattern
	x, y     x
}


func (p *Pen) MoveTo(px, py x) LimitedPattern{
	p.x, p.y= px, py
	return Translated{p.Marker,p.x,p.y}
}

func (p *Pen) LineTo(px, py x) LimitedPattern {
	s := p.Straight(p.x, p.y, px, py)
	p.x, p.y = px, py
	if p.Marker==nil{
		return s
	}
	return LimitedComposite{s,Translated{p.Marker,p.x,p.y}}
}


func (p *Pen) LineToVertical(py x) LimitedPattern {
	s := p.Straight(p.x, p.y, p.x, py)
	p.y = py
	if p.Marker==nil{
		return s
	}
	return LimitedComposite{s,Translated{p.Marker,p.x,p.y}}
}

func (p *Pen) LineToHorizontal(px x) LimitedPattern {
	s := p.Straight(p.x, p.y, px, p.y)
	p.x = px
	if p.Marker==nil{
		return s
	}
	return LimitedComposite{s,Translated{p.Marker,p.x,p.y}}
}


func (p *Pen) ArcTo(rx,ry x, a float64, large,sweep bool,x,y x) LimitedPattern {
	s := p.Conic(p.x,p.y,rx,ry,a,large,sweep,x,y)
	p.x, p.y=x,y
	if p.Marker==nil{
		return s
	}
	return LimitedComposite{s,Translated{p.Marker,p.x,p.y}}
}

func (p *Pen) QuadraticBezierTo(cx,cy,px,py x) LimitedPattern {
	s := p.Curved(p.x,p.y,cx,cy,cx,cy,px, py)
	p.x, p.y=px,py
	if p.Marker==nil{
		return s
	}
	return LimitedComposite{s,Translated{p.Marker,p.x,p.y}}
}

func (p *Pen) CubicBezierTo(c1x,c1y,c2x,c2y,px,py x) LimitedPattern {
	s := p.Curved(p.x,p.y,c1x,c1y,c2x,c2y,px, py)
	p.x, p.y=px,py
	if p.Marker==nil{
		return s
	}
	return LimitedComposite{s,Translated{p.Marker,p.x,p.y}}
}



// PenPath have methods to create LimitedPatterns using a Pen and an origin point.
type PenPath struct{
	Pen
	x,y x
}

func (p *PenPath) MoveTo(px, py x) LimitedPattern{
	s:=p.Pen.MoveTo(px,py)
	p.x, p.y= px, py
	return s
}

func (p *PenPath) StartLine(x1, y1, x2, y2 x) LimitedPattern {
	return LimitedComposite{p.MoveTo(x1,y1),p.LineTo(x2, y2)}
}

func (p *PenPath) LineClose() LimitedPattern {
	s := p.Straight(p.Pen.x,p.Pen.y,p.x, p.y)
	p.x, p.y = p.Pen.x,p.Pen.y
	return s
}

