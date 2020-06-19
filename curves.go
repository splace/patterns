package patterns

import "math"
//import "fmt"
// Divide provides channels of a number of points along various curves (not including ends)
type Divide uint8
const dividerMax = math.MaxUint8

// Curve provides variable numbers of intermediate points from two independent axis functions.
func (d Divide) Curve(xfn, yfn func(Divide)x)  <-chan [2]x {
	ch:=make(chan [2]x,d)
	step:=dividerMax/d
	var li Divide
	go func(){
		for i := step-1; li<i ; li,i=i,i+step {
			ch <- [2]x{xfn(i),yfn(i)}
		}
		close(ch)
	}()
	return ch
}

// bezier curves direct from definition, that is; hierarchical linear division, this gives more points where more curvature.

func (d Divide) QuadraticBezier(sx, sy, cx, cy, ex, ey x)  <-chan [2]x {
	return  d.Curve(doubleDivision(sx, cx, ex),doubleDivision(sy, cy, ey))
}

func (d Divide) CubicBezier(sx, sy, c1x, c1y, c2x, c2y, ex, ey x)  <-chan [2]x {
	return  d.Curve(tripleDivision(sx, c1x, c2x, ex),tripleDivision(sy, c1y, c2y, ey))
}

func (d Divide) QuinticBezier(sx, sy, c1x, c1y, c2x,c2y, c3x,c3y, ex, ey x)  <-chan [2]x {
	return  d.Curve(quadroupleDivision(sx, c1x, c2x, c3x, ex),quadroupleDivision(sy, c1y, c2y, c3y, ey))
}




// find centre of a circle given two points on rim and radius.
func centreOfCircle(sx, sy, r, ex, ey x) (x,y float64){
	// midpoint
	mx,my:=(ex+sx)>>1,(ey+sy)>>1
	// vector along midline times distance apart
	vmx,vmy:=sy-ey,ex-sx
	// distance apart squared
	d2:=vmx*vmx+vmy*vmy
	//
	r2:=r*r
	if d2>4*r2 {
		return float64(mx),float64(my)
	}
	// multiplying factor of centre along midline
	m:=math.Sqrt(float64(r2)/float64(d2)-0.25)
	// centre
	return float64(mx)+float64(vmx)*m,float64(my)+float64(vmy)*m
}

// functions that rotate clockwise and counterclockwise by the provided angle
func rotaters (a float64) (func(float64,float64)(float64,float64),func(float64,float64)(float64,float64)){
	sa,ca:=math.Sincos(a)
	return func(x,y float64) (float64, float64) {
		return x*ca+y*sa, y*ca-x*sa
	},
	func(x,y float64) (float64, float64) {
		return x*ca-y*sa, y*ca+x*sa
	}
}

func offsetRotaters (ox,oy,a float64) (func(float64,float64)(float64,float64),func(float64,float64)(float64,float64)){
	cwr,ccwr:=rotaters(a)
	return offsetter(ox,oy,cwr),offsetter(ox,oy,ccwr)
}

func offsetter (ox,oy float64, t func(float64,float64)(float64,float64)) (func(float64,float64)(float64,float64)){
	return func(x,y float64) (float64, float64) {
		x,y=t(x-ox,y-oy)
		return x+ox, y+oy
	}
}

func xsquashers (s float64) (func(float64,float64)(float64,float64),func(float64,float64)(float64,float64)){
	rs:=1/s
	return func(x,y float64) (float64, float64) {
		return x*rs,y
	},
	func(x,y float64) (float64, float64) {
		return x*s,y
	}
}



func  (d Divide) Arc(x1,y1,rx,ry x, a float64, large,sweep bool, x2,y2 x)  <-chan [2]x{
	if rx!=ry{
		// for rx!=ry use squash and/or rotate transforms, then do Circle Sector and reverse transform on points returned. 
		if a!=0 {
			cwa,ccwa:=rotaters(a)
			scwa,usccwa:=xsquashers(float64(rx)/float64(ry))
			tx1,ty1:=scwa(cwa(float64(x1),float64(y1)))
			tx2,ty2:=scwa(cwa(float64(x2),float64(y2)))
			ch:=make(chan [2]x,d-1)
			go func(){
				for l:=range(d.Sector(x(tx1),x(ty1),x(ry),large,sweep,x(tx2),x(ty2))){
					utlx,utly:=ccwa(usccwa(float64(l[0]),float64(l[1])))
					ch <- [2]x{x(utlx),x(utly)}
				}
				close(ch)
			}()
			return ch
		}else{
			sa,usa:=xsquashers(float64(rx)/float64(ry))
			tx1,ty1:=sa(float64(x1),float64(y1))
			tx2,ty2:=sa(float64(x2),float64(y2))
			ch:=make(chan [2]x,d-1)
			go func(){
				for l:=range(d.Sector(x(tx1),x(ty1),x(ry),large,sweep,x(tx2),x(ty2))){
					utlx,utly:=usa(float64(l[0]),float64(l[1]))
					ch <- [2]x{x(utlx),x(utly)}
				}
				close(ch)
			}()
			return ch
		
		}
	}
	return d.Sector(x1,y1,rx,large,sweep,x2,y2)
}

func  (d Divide) Sector(x1,y1,r x, large,sweep bool, x2,y2 x)  <-chan [2]x{
	var cx,cy float64
	// two possible centres for short sweep
	if large != sweep {
		cx,cy= centreOfCircle(x1,y1,r,x2,y2)
	}else{
		cx,cy= centreOfCircle(x2,y2,r,x1,y1)
	}
	// find angles from centre to start and end points 
	a1,a2:=math.Atan2(float64(y1)-cy,float64(x1)-cx),math.Atan2(float64(y2)-cy,float64(x2)-cx)
	// delta angle from start to end, calculation issues due to float64 type at odds with modula of angles
	da:=a2-a1
	// a1 and a2 can only be +-Pi, but da can be +-2Pi
	if large {
		if da<0 && da>= -math.Pi {
			da+=2*math.Pi
		}else{
			if da>0 && da<math.Pi {da-=2*math.Pi}
		}
	}else{
		if da<= -math.Pi {
			da+=2*math.Pi
		}else{
			if da>math.Pi {da-=2*math.Pi}
		}
	}
	// make rotation direction as required for sweep 
	if sweep == (da<0) {da=-da}
	// atan2 produces angles counter-clockwise from +ve x-axis
	_,occwr:=offsetRotaters(cx,cy,da/float64(d))
	ch:=make(chan [2]x,d-1)
	dx,dy:= float64(x1),float64(y1)
	go func(){
		for li:=Divide(1); li<d ; li++ {
			dx,dy=occwr(dx,dy)
			ch <- [2]x{x(dx),x(dy)}
		}
		close(ch)
	}()
	return ch
}

func linearDivision(f x) (func (Divide) x ){
	return func(t Divide)x{return f*x(t)/dividerMax} 
}
	
func doubleDivision(s,c,e x) (func (Divide) x ){
	scfn:= linearDivision(c-s)
	cefn:= linearDivision(e-c)
	return func(t Divide)x{
		return s+scfn(t)+linearDivision(-s-scfn(t)+c+cefn(t))(t)
	}
}

func tripleDivision(s,c1,c2,e x) (func (Divide) x ){
	sc1fn:= linearDivision(c1-s)
	c1c2fn:= linearDivision(c2-c1)
	c2efn:= linearDivision(e-c2)
	return func(t Divide)x{
		return doubleDivision(s+sc1fn(t),c1+c1c2fn(t),c2+c2efn(t))(t)
	}
}

func quadroupleDivision(s,c1,c2,c3,e x) (func (Divide) x ){
	sc1fn:= linearDivision(c1-s)
	c1c2fn:= linearDivision(c2-c1)
	c2c3fn:= linearDivision(c3-c2)
	c3efn:= linearDivision(e-c3)
	return func(t Divide)x{
		return tripleDivision(s+sc1fn(t),c1+c1c2fn(t),c2+c2c3fn(t),c3+c3efn(t))(t)
	}
}
