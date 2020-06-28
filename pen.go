package patterns

// Pens have methods to create LimitedPatterns depending on, and maintaining, its current location.
// Optionally it adds a LimitedPattern at the joins between segments.
type Pen struct {
	Nib
	x, y     x
	Joiner LimitedPattern
	previousWasMove bool
}

func (p *Pen) AddMark(l LimitedPattern) LimitedPattern {
	if p.previousWasMove || p.Joiner==nil{
		p.previousWasMove=false
		return l
	}
	return LimitedComposite{l,Translated{p.Joiner,p.x,p.y}}
}

func (p *Pen) MoveTo(px, py x) {
	p.x, p.y= px, py
	p.previousWasMove=true
	return
}

func (p *Pen) LineTo(px, py x) (l LimitedPattern) {
	l= p.AddMark(p.Straight(p.x, p.y, px, py))
	p.x, p.y = px, py
	return
}

func (p *Pen) LineToVertical(py x) (l LimitedPattern) {
	l=p.AddMark(p.Straight(p.x, p.y, p.x, py))
	p.y = py
	return
}

func (p *Pen) LineToHorizontal(px x) (l LimitedPattern) {
	l =p.AddMark(p.Straight(p.x, p.y, px, p.y))
	p.x = px
	return
}

func (p *Pen) ArcTo(rx,ry x, a float64, large,sweep bool,x,y x) (l LimitedPattern) {
	l = p.AddMark(p.Conic(p.x,p.y,rx,ry,a,large,sweep,x,y))
	p.x, p.y=x,y
	return
}

func (p *Pen) QuadraticBezierTo(cx,cy,px,py x) (l LimitedPattern) {
	l = p.AddMark(p.Curved(p.x,p.y,cx,cy,cx,cy,px, py))
	p.x, p.y=px,py
	return
}

func (p *Pen) CubicBezierTo(c1x,c1y,c2x,c2y,px,py x) (l LimitedPattern) {
	l= p.AddMark(p.Curved(p.x,p.y,c1x,c1y,c2x,c2y,px, py))
	p.x, p.y=px,py
	return
}

// PenPath's have methods to create LimitedPatterns depending on a Pen and if that pen is continuously drawing without moves, or not.
type PenPath struct{
	Pen
	x,y x
}

func (p *PenPath) MoveTo(px, py x) (l LimitedPattern){
	p.x, p.y= px, py
	p.Pen.MoveTo(px,py)
	return
}

func (p *PenPath) LineClose() (l LimitedPattern) {
	l = p.AddMark(p.LineTo(p.x, p.y))
	return
}

