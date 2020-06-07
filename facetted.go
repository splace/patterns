package patterns

import "math"
//import "fmt"

// TODO brush interface for alternative ways to draw stuff

// Facetted is a Nib that produces patterns with curves sub-divided into straight lines.
// CurveDivision:  power of 2 number of divisions.
// default 0 - no division, all curves a single straight line
type Facetted struct{
	Width    x
	In       y
	CurveDivision uint8
}



func (p Facetted) Line(x1, y1, x2, y2 x) LimitedPattern {
	ndx,dy:=float64(x1-x2),float64(y2-y1)
	// NewRotated actually returns a LimitedPattern (as a Pattern) because NewLine returns one, so assert can never fail.
	// TODO here we know MaxX better than z.
	return Translated{NewRotated(Rectangle(x(math.Hypot(ndx,dy)),p.Width, Filling{p.In}),math.Atan2(dy,ndx)).(LimitedPattern),(x1+x2)>>1, (y1+y2)>>1}
}



func (p Facetted) Arc(x1,y1,rx,ry x, a float64, large,sweep bool, x2,y2 x) LimitedPattern {
	// if ellipse too small expand to just fit, which will depend on angle, uggg.
	if rx==ry{
		// much simpler, just a circle, angle redundent
		return nil
	}
	// interesting solution
	// using conic projection (so need to go to 3D) rather than squashing a circle, equal angle separation gives more points at tighter curvature. 
	// except for far half, so simply reuse use near half reflected.
	// for this the CurveDivision parameter divides each half, pro-rata to angle needed.
	// cone is tipped along major axis by factor from radius's.
	//  
	// in order for ends to align, cone needs to be positioned.
	// and section plane needs to pass through start and end points and be angled to give radius ratio
	
	//  centre of cone is equi-distance from ends
	
	// centre line of cone goes through ellipse foci
	//foci2:=rx*rx-ry*ry
	// apex angle of cone is a
	//cca:=0.707   // cosine cone apex angle
	// angle of cone to plane from axis lengths
	//t:=math.Acos(math.Sqrt(1-float64(rx*rx)/float64(ry*ry))*cca) // if rx<ry

	//t:=math.Acos(math.Sqrt(ry*ry-rx*rx)/ry*cca) // if rx<ry
   	
	
////	Cone apex and centre intercept points
////Two intercepts selected by sweep
////Relative apex from angle and rx and ry
////Intercept, cone midline, ellipse foci, with start and end on edge, so, foci along points midline right angle.
////A lot already coded.
////	

	// roatate line ofset in 3d for cone

// 
	
	return nil
}

func (p Facetted) Box(x,y x) LimitedPattern {
	return Limiter{Composite{p.Line(-x,y, x,y),p.Line(x,y,x,-y),p.Line(x,-y,-x,-y),p.Line(-x,-y,-x,y)},max(x+p.Width,y+p.Width)}
}

func (p Facetted) Polygon(coords ...[2]x) Pattern {
	// TODO calc limits
	s := make([]Pattern, len(coords)) 
	for i := 1; i < len(s); i++ {
		s[i-1] = p.Line(coords[i-1][0], coords[i-1][1],coords[i][0], coords[i][1])
	}
	s[len(coords)-1] = p.Line(coords[len(coords)-1][0], coords[len(coords)-1][1],coords[0][0], coords[0][1])
	return NewComposite(s...)
}

type divider uint8
const dividerMax = math.MaxUint8

func linearDivision(f x) (func (divider) x ){
		return func(t divider)x{return f*x(t)/dividerMax} 
	}
	
func doubleDivision(s,c,e x) (func (divider) x ){
		scfn:= linearDivision(c-s)
		cefn:= linearDivision(e-c)
		return func(t divider)x{
			return s+scfn(t)+linearDivision(-s-scfn(t)+c+cefn(t))(t)
		}
	}

func tripleDivision(s,c1,c2,e x) (func (divider) x ){
		sc1fn:= linearDivision(c1-s)
		c1c2fn:= linearDivision(c2-c1)
		c2efn:= linearDivision(e-c2)
		return func(t divider)x{
			return doubleDivision(s+sc1fn(t),c1+c1c2fn(t),c2+c2efn(t))(t)
		}
	}

func quadroupleDivision(s,c1,c2,c3,e x) (func (divider) x ){
		sc1fn:= linearDivision(c1-s)
		c1c2fn:= linearDivision(c2-c1)
		c2c3fn:= linearDivision(c3-c2)
		c3efn:= linearDivision(e-c3)
		return func(t divider)x{
			return tripleDivision(s+sc1fn(t),c1+c1c2fn(t),c2+c2c3fn(t),c3+c3efn(t))(t)
		}
	}


//func linearDivision(s,e x) (func (divider) x ){
//		return func(t divider)x{return s+(e-s)*x(t)/dividerMax} 
//	}
//	
//func doubleDivision(s,c,e x) (func (divider) x ){
//		scfn:= linearDivision(s,c)
//		cefn:= linearDivision(c,e)
//		return func(t divider)x{
//			return linearDivision(scfn(t),cefn(t))(t)
//		}
//	}

//func tripleDivision(s,c1,c2,e x) (func (divider) x ){
//		sc1fn:= linearDivision(s,c1)
//		c1c2fn:= linearDivision(c1,c2)
//		c2efn:= linearDivision(c2,e)
//		return func(t divider)x{
//			return doubleDivision(sc1fn(t),c1c2fn(t),c2efn(t))(t)
//		}
//	}

//func quadroupleDivision(s,c1,c2,c3,e x) (func (divider) x ){
//		sc1fn:= linearDivision(s,c1)
//		c1c2fn:= linearDivision(c1,c2)
//		c2c3fn:= linearDivision(c2,c3)
//		c3efn:= linearDivision(c3,e)
//		return func(t divider)x{
//			return tripleDivision(sc1fn(t),c1c2fn(t),c2c3fn(t),c3efn(t))(t)
//		}
//	}

func (p Facetted) QuadraticBezier(sx, sy, cx, cy, ex, ey x) LimitedPattern {
	xfn:=doubleDivision(sx, cx, ex)
	yfn:=doubleDivision(sy, cy, ey)
	var s []Pattern
	var li divider
	step:=divider(1<<(8-p.CurveDivision))
	for i := step-1; li<i ; li,i=i,i+step {
		s= append(s,p.Line(xfn(li),yfn(li),xfn(i),yfn(i)))
	}
	return Limiter{NewComposite(s...),max(max(max(sx,ex),cx),max(max(sy,ey),cy))+p.Width}
}


func (p Facetted) CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey x) LimitedPattern {
	xfn:=tripleDivision(sx, c1x, c2x, ex)
	yfn:=tripleDivision(sy, c1y, c2y, ey)
	var s []Pattern
	var li divider
	step:=divider(1<<(8-p.CurveDivision))
	for i := step-1; li<i ; li,i=i,i+step {
		s= append(s,p.Line(xfn(li),yfn(li),xfn(i),yfn(i)))
	}
	return Limiter{NewComposite(s...),max(max(max(sx,ex),max(c1x,c2x)),max(max(sy,ey),max(c1y,c2y)))+p.Width}  
}


func (p Facetted) QuinticBezier(sx, sy, c1x, c1y, c2x,c2y, c3x,c3y, ex, ey x) LimitedPattern {
	xfn:=quadroupleDivision(sx, c1x, c2x, c3x, ex)
	yfn:=quadroupleDivision(sy, c1y, c2y, c3y, ey)
	var s []Pattern
	var li divider
	step:=divider(1<<(8-p.CurveDivision))
	for i := step-1; li<i ; li,i=i,i+step {
		s= append(s,p.Line(xfn(li),yfn(li),xfn(i),yfn(i)))
	}
	return Limiter{NewComposite(s...),max(max(max(max(sx,ex),max(c1x,c2x)),c3x),max(max(max(sy,ey),max(c1y,c2y)),c3y))+p.Width }
}

