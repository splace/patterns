package patterns


// Pens have methods to create LimitedPatterns depending on, and maintaining, its current location.
// Optionally it includes a LimitedPattern at the end of each.
type Pen struct {
	Nib
	Marker LimitedPattern
	x, y     x
}


func (p *Pen) MoveTo(px, py x) {
	p.x, p.y= px, py
	return
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


// PenPath's have methods to create LimitedPatterns depending on a Pen and a start location.
// Optionally it includes a LimitedPattern's for the start and/or end.
type PenPath struct{
	Pen
	StartMarker,EndMarker LimitedPattern
	x,y x
}

func (p *PenPath) MoveTo(px, py x) LimitedPattern{
	p.x, p.y= px, py
	p.Pen.MoveTo(px,py)
	if p.StartMarker!=nil{
		return Translated{p.StartMarker,p.x,p.y}
	}
	if p.Marker!=nil{
		return Translated{p.Marker,p.x,p.y}
	}
	return nil
}

func (p *PenPath) LineClose() LimitedPattern {
	s := p.Straight(p.Pen.x,p.Pen.y,p.x, p.y)
	p.x, p.y = p.Pen.x,p.Pen.y
	return s
}

