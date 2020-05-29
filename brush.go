package patterns

import "math"
//import "fmt"


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
	for i := 1; i < len(s); i++ {
		s[i-1] = p.Line(coords[i-1][0], coords[i-1][1],coords[i][0], coords[i][1])
	}
	s[len(coords)-1] = p.Line(coords[len(coords)-1][0], coords[len(coords)-1][1],coords[0][0], coords[0][1])
	return NewComposite(s...)
}

type bezierResolution uint8
const bezierDivision bezierResolution= 1<<6   // divides into 4
const bezierMax = math.MaxUint8

func linearDivision(s,e x) (func (bezierResolution) x ){
		return func(t bezierResolution)x{return s+(e-s)*x(t)/bezierMax} 
	}
	
func doubleDivision(s,c,e x) (func (bezierResolution) x ){
		scfn:= linearDivision(s,c)
		cefn:= linearDivision(c,e)
		return func(t bezierResolution)x{
			return linearDivision(scfn(t),cefn(t))(t)
		}
	}

func tripleDivision(s,c1,c2,e x) (func (bezierResolution) x ){
		sc1fn:= linearDivision(s,c1)
		c1c2fn:= linearDivision(c1,c2)
		c2efn:= linearDivision(c2,e)
		return func(t bezierResolution)x{
			return doubleDivision(sc1fn(t),c1c2fn(t),c2efn(t))(t)
		}
	}

func quadroupleDivision(s,c1,c2,c3,e x) (func (bezierResolution) x ){
		sc1fn:= linearDivision(s,c1)
		c1c2fn:= linearDivision(c1,c2)
		c2c3fn:= linearDivision(c2,c3)
		c3efn:= linearDivision(c3,e)
		return func(t bezierResolution)x{
			return tripleDivision(sc1fn(t),c1c2fn(t),c2c3fn(t),c3efn(t))(t)
		}
	}

func (p LineBrush) QuadraticBezier(sx, sy, cx, cy, ex, ey x) LimitedPattern {
	xfn:=doubleDivision(sx, cx, ex)
	yfn:=doubleDivision(sy, cy, ey)
	var s []Pattern
	var li bezierResolution
	for i := bezierDivision-1; li<i ; li,i=i,i+bezierDivision {
		s= append(s,p.Line(xfn(li),yfn(li),xfn(i),yfn(i)))
	}
	return Limiter{NewComposite(s...),1000*unitX}
}


func (p LineBrush) CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey x) LimitedPattern {
	xfn:=tripleDivision(sx, c1x, c2x, ex)
	yfn:=tripleDivision(sy, c1y, c2y, ey)
	var s []Pattern
	var li bezierResolution
	for i := bezierDivision-1; li<i ; li,i=i,i+bezierDivision {
		s= append(s,p.Line(xfn(li),yfn(li),xfn(i),yfn(i)))
	}
	return Limiter{NewComposite(s...),1000*unitX}  
}


func (p LineBrush) QuinticBezier(sx, sy, c1x, c1y, c2x,c2y, c3x,c3y, ex, ey x) LimitedPattern {
	xfn:=quadroupleDivision(sx, c1x, c2x, c3x, ex)
	yfn:=quadroupleDivision(sy, c1y, c2y, c3y, ey)
	var s []Pattern
	var li bezierResolution
	for i := bezierDivision-1; li<i ; li,i=i,i+bezierDivision {
		s= append(s,p.Line(xfn(li),yfn(li),xfn(i),yfn(i)))
	}
	return Limiter{NewComposite(s...),1000*unitX}
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
