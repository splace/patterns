package patterns

import "math"
import "fmt"

// Facetted is a Nib using straight lines with a particular width.
// Curves are divided using CurveDivision:  power of 2 number of divisions.
// default 0 - no division, all curves a single straight line
// it uses direct definition of bezier curves, cascading linear division, to give more lines where more curvature.
// * bezier curves are limited to being within hull of control points.
// uses conic projection for arc, again more lines where more curvature.
type Facetted struct{
	Width    x
	In       y
	CurveDivision uint8
}

func (p Facetted) Line(x1, y1, x2, y2 x) LimitedPattern {
	ndx,dy:=float64(x1-x2),float64(y2-y1)
	// NewRotated actually returns a LimitedPattern (as a Pattern) because NewLine returns one, so assert can never fail.
	// TODO could reduce MaxX since we know better than worst case used by rotate.
	return Translated{NewRotated(Rectangle(x(math.Hypot(ndx,dy)),p.Width, Filling(p.In)),math.Atan2(dy,ndx)).(LimitedPattern),(x1+x2)>>1, (y1+y2)>>1}
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
	return Limiter{NewComposite(s...),max2(cx,cy)+p.Width}  
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
	return Limiter{NewComposite(s...),max4(c1x,c2x,c1y,c2y)+p.Width}  
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
	return Limiter{NewComposite(s...),max6(c1x,c2x,c3x,c1y,c2y,c3y)+p.Width}  
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

// find centre of circle given two points on rim and radius.
func circleCentre(sx, sy, r, ex, ey x) (x,y float64){
	// midpoint
	mx,my:=(ex+sx)>>1,(ey+sy)>>1
	// vector along midline times distance apart
	vmx,vmy:=sy-ey,ex-sx
	// distance apart squared
	d2:=vmx*vmx+vmy*vmy
	//
	r2:=r*r
	if d2>4*r2 {
		panic("circle too small")
	}
	// multiplying factor of centre along midline
	m:=math.Sqrt(float64(r2)/float64(d2)-0.25)
	// centre
	return float64(mx)+float64(vmx)*m,float64(my)+float64(vmy)*m
}

// functions that rotate clockwise and counterclockwise by the provided angle
func Rotaters (a float64) (func(float64,float64)(float64,float64),func(float64,float64)(float64,float64)){
	sa,ca:=math.Sincos(a)
	return func(x,y float64) (float64, float64) {
		return x*ca+y*sa, y*ca-x*sa
	},
	func(x,y float64) (float64, float64) {
		return x*ca-y*sa, y*ca+x*sa
	}
}

func OffsetRotaters (ox,oy,a float64) (func(float64,float64)(float64,float64),func(float64,float64)(float64,float64)){
	sa,ca:=math.Sincos(a)
	return func(x,y float64) (float64, float64) {
		x-=ox
		y-=oy
		return x*ca+y*sa+ox, y*ca-x*sa+oy
	},
	func(x,y float64) (float64, float64) {
		x-=ox
		y-=oy
		return x*ca-y*sa+ox, y*ca+x*sa+oy
	}
}


func (p Facetted) Arc(x1,y1,rx,ry x, a float64, large,sweep bool, x2,y2 x) LimitedPattern {
	// if ellipse too small expand to just fit, which will depend on angle, uggg.
	// TODO for rx!=ry translate/rotate and squash, then do below then reverse transform on every point. 
	if rx==ry{
		// much simpler, just a circle, angle redundant
		var cx,cy,a1,a2 float64
		if large == sweep {
			cx,cy= circleCentre(x1,y1,rx,x2,y2)
		}else{
			cx,cy= circleCentre(x2,y2,rx,x1,y1)
		}
		fmt.Println(cx,cy)
		a1,a2=math.Atan2(float64(x1)-cx,float64(y1)-cy),math.Atan2(float64(x2)-cx,float64(y2)-cy)
		fmt.Println(a1,a2)
		if !sweep {
			a1,a2=a2,a1+math.Pi*2
		}
		fmt.Println("Angles:",a1,a2)
		// scale divisions so you get, somewhat, consistent side angles
		halfDivisions:=int8(math.Abs(float64(uint8(1)<<p.CurveDivision)*(a2-a1)/math.Pi))+1
		ocwr,_:=OffsetRotaters(cx,cy,(a2-a1)*.5/float64(halfDivisions))
		s := make([]Pattern, halfDivisions*2) 
		maxx:=max2(max2(x1,y1),max2(x2,y2))
		dx,dy:= float64(x1),float64(y1)
		for i:=range(s[:len(s)-1]){
			ex,ey:=ocwr(dx,dy)
			fmt.Println(ex,ey)
			s[i]=p.Line(x(dx),x(dy),x(ex),x(ey))
			dx,dy=ex,ey
			maxx=max2(maxx,max2(x(ex),x(ey)))
		}
		s[len(s)-1]=p.Line(x(dx),x(dy),x2,y2)
	return Limiter{NewComposite(s...),maxx+p.Width}
	}
		
//	//fmt.Println(x,y,x1, y1, rx, ry,w, a , laf, psf)
//	// transform to put start and end points on a unit radius circle.
//	// * rotate to line ellipse up on axis
//	cwRter,ccwRter:=Rotaters(a*(180/math.Pi))
//	tx1, ty1 := cwRter(float64(x1)*scaleX,float64(y1)*scaleX)
//	tx2, ty2 := cwRter(float64(x2)*scaleX,float64(y2)*scaleX)
//	// scale to make unit radius
//	
//	tx1 /= float64(rx)*scaleX
//	ty1 /= float64(ry)*scaleX
//	tx2 /= float64(rx)*scaleX
//	ty2 /= float64(ry)*scaleX
//	// find distance between transformed start and end, (cord length)
//	tdx, tdy := tx1-tx2, ty1-ty2
//	td := math.Hypot(tdx,tdy)
//	// can't exceed 2, otherwise ellipse is too small to pass through both start and end.
//	if td > 2 {
//		panic(td)
//	}
//	// now find the centre of this unit circle
//	tdc := math.Sqrt(1-td*td/4) / td // this is, relative, to cord length, distance, at right angles, from mid-point, that is 1 unit from both start and end.
//	var tcx, tcy float64
//	// project to find actual center (transformed)
//	if large == sweep {
//		// on left for; large and cw, or small and cww
//		tcx = (tx2+tx1)/2 - tdc*tdy
//		tcy = (ty2+ty1)/2 + tdc*tdx
//	} else {
//		// on right
//		tcx = (tx2+tx1)/2 + tdc*tdy
//		tcy = (ty2+ty1)/2 - tdc*tdx
//	}
//	// reverse transform to find actual center
//	cxt, cyt := tcx*float64(rx)*scaleX, tcy*float64(ry)*scaleX
//	cx,cy := ccwRter(cxt,cyt)
//	// optimisation: pre-calc reciprocal of inner and outer radii squared
//	
//	fmt.Println(cx,cy)
//	
//	if rx==ry{
//		// much simpler, just a circle, angle redundent
//		return nil
//	}
//	// interesting solution
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
