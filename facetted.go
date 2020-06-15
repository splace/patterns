package patterns

import "math"

// Facetted is a Nib producing curves using a number of straight lines.
// curves are divided according to CurveDivision:  (power of 2 number of divisions.)
// default 0 - no division, all curves a single straight line
// if a Nib is provided its Line method is used to draw the straight lines.
type Facetted struct{
	Nib
	Width    x
	In       y
	CurveDivision uint8
//	Lwidth x // last width to make tappered lines
}

func (p Facetted) Line(x1, y1, x2, y2 x) LimitedPattern {
	if p.Nib==nil{
		ndx,dy:=float64(x1-x2),float64(y2-y1)
		// NewRotated actually returns a LimitedPattern (as a Pattern) because NewLine returns one, so assert can never fail.
		// TODO could reduce MaxX since we know better than worst case used by rotate.
		return Translated{NewRotated(Rectangle(x(math.Hypot(ndx,dy)),p.Width, Filling(p.In)),math.Atan2(dy,ndx)).(LimitedPattern),(x1+x2)>>1, (y1+y2)>>1}
	}
	return p.Nib.Line(x1, y1, x2, y2)
}

func (p Facetted) QuadraticBezier(sx, sy, cx, cy, ex, ey x) LimitedPattern {
	var s []Pattern
	maxx:=max2(max2(sx,sy),max2(ex,ey))
	for l:=range(Divide(1<<(8-p.CurveDivision)).QuadraticBezier(sx, sy, cx, cy, ex, ey)){
		s= append(s,p.Line(sx,sy,l[0],l[1]))
		sx,sy=l[0],l[1]
		maxx=max2(maxx,max2(sx,sy))
	}
	s= append(s,p.Line(sx,sy,ex,ey))
	return Limiter{NewComposite(s...),maxx+p.Width}  
}


func (p Facetted) CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey x) LimitedPattern {
	var s []Pattern
	maxx:=max2(max2(sx,sy),max2(ex,ey))
	for l:=range(Divide(1<<(8-p.CurveDivision)).CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey)){
		s= append(s,p.Line(sx,sy,l[0],l[1]))
		sx,sy=l[0],l[1]
		maxx=max2(maxx,max2(sx,sy))
	}
	s= append(s,p.Line(sx,sy,ex,ey))
	return Limiter{NewComposite(s...),maxx+p.Width}  
}


func (p Facetted) QuinticBezier(sx, sy, c1x, c1y, c2x,c2y, c3x,c3y, ex, ey x) LimitedPattern {
	var s []Pattern
	maxx:=max2(max2(sx,sy),max2(ex,ey))
	for l:=range(Divide(1<<(8-p.CurveDivision)).QuinticBezier(sx, sy, c1x, c1y, c2x,c2y, c3x, c3y,ex, ey)){
		s= append(s,p.Line(sx,sy,l[0],l[1]))
		sx,sy=l[0],l[1]
		maxx=max2(maxx,max2(sx,sy))
	}
	s= append(s,p.Line(sx,sy,ex,ey))
	return Limiter{NewComposite(s...),maxx+p.Width}  
}


func (p Facetted) Arc(sx,sy,rx,ry x, a float64, large,sweep bool, ex,ey x) LimitedPattern {
	var s []Pattern
	maxx:=max2(max2(sx,sy),max2(ex,ey))
	for l:=range(Divide(1<<(8-p.CurveDivision)).Arc(sx,sy,rx,ry, a, large,sweep, ex,ey)){
		s= append(s,p.Line(sx,sy,l[0],l[1]))
		sx,sy=l[0],l[1]
		maxx=max2(maxx,max2(sx,sy))
	}
	s= append(s,p.Line(sx,sy,ex,ey))
	return Limiter{NewComposite(s...),maxx+p.Width}  
}


func (p Facetted) Box(x,y x) LimitedPattern {
	return Limiter{Composite{p.Line(-x,y, x,y),p.Line(x,y,x,-y),p.Line(x,-y,-x,-y),p.Line(-x,-y,-x,y)},max2(x+p.Width,y+p.Width)}
}

func (p Facetted) Polygon(coords ...[2]x) LimitedPattern {
	s := make([]Pattern, len(coords)) 
	maxx:=max2(coords[0][0], coords[0][1])
	for i := 1; i < len(s); i++ {
		s[i-1] = p.Line(coords[i-1][0], coords[i-1][1],coords[i][0], coords[i][1])
		maxx=max2(maxx,max2(coords[i][0], coords[i][1]))
	}
	s[len(coords)-1] = p.Line(coords[len(coords)-1][0], coords[len(coords)-1][1],coords[0][0], coords[0][1])
	return Limiter{NewComposite(s...),maxx+p.Width}
}

