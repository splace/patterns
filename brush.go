package patterns

import (
	"math"
)

type LineBrush struct{
	Width    x
	In       y

}

func (p LineBrush) Line(px1, py1, px2, py2 x) LimitedPattern {
	return Translated{NewRotated(Reduced{Square{Filling{p.In}}, float32(unitX*2)/float32(math.Sqrt(float64(px2-px1)*float64(px2-px1) + float64(py2-py1)*float64(py2-py1))),float32(unitX*4)/float32(p.Width) },math.Atan2(float64(py1-py2),float64(px2-px1)) ).(LimitedPattern),(px1+px2)/2, (py1+py2)/2}
}


func (p LineBrush) Box(x,y x) LimitedPattern {
	return Limiter{Composite{p.Line(-x,y, x,y),p.Line(x,y,x,-y),p.Line(x,-y,-x,-y),p.Line(-x,-y,-x,y)},max4(x+p.Width,p.Width-x,p.Width-y,y-p.Width)}
}

func (p LineBrush) Polygon(coords ...[2]x) Pattern {
	// TODO calc limits
	s := make([]Pattern, len(coords)) 
	for i := 1; i < len(coords); i++ {
		s[i-1] = p.Line(coords[i-1][0], coords[i-1][1],coords[i][0], coords[i][1])
	}
	s[len(coords)-1] = p.Line(coords[len(coords)-1][0], coords[len(coords)-1][1],coords[0][0], coords[0][1])
	return NewComposite(s...)
}


func (p LineBrush) QuadraticBezier(sx, sy, cx, cy, ex, ey x) LimitedPattern {
	ComponentQuadBezier:= func(s,c,e x) (func (uint16) x ){
		f:=s+(e-s)
		return func(t uint16)x{return f*x(t/math.MaxUint16)}   // FIXME linear
	}
	xfn:=ComponentQuadBezier(sx, cx, ex)
	yfn:=ComponentQuadBezier(sy, cy, ey)
	l:= p.Line(xfn(0),yfn(0),xfn(math.MaxUint16),yfn(uint16(math.MaxUint16)))
	return Limiter{Composite{l},l.MaxX()}
}



func (p LineBrush) CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey x) LimitedPattern {
	ComponentCubicBezier:= func(s,c1,c2,e x) (func (uint16) x ){
		f:=s+(e-s)
		return func(t uint16)x{return f*x(t/math.MaxUint16)}   // FIXME linear
	}
	xfn:=ComponentCubicBezier(sx, c1x, c2x, ex)
	yfn:=ComponentCubicBezier(sy, c1y, c2y, ey)
	l:= p.Line(xfn(0),yfn(0),xfn(math.MaxUint16),yfn(uint16(math.MaxUint16)))
	return Limiter{Composite{l},l.MaxX()}
}


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
