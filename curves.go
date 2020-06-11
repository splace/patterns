package patterns

import "math"
import "fmt"


type Divide uint8
const dividerMax = math.MaxUint8

func (d Divide) Curve(xfn, yfn func(Divide)x)  <-chan [2]x {
	step:=Divide(dividerMax/d)
	ch:=make(chan [2]x,d)
	var li Divide
	go func(){
		for i := step-1; li<i ; li,i=i,i+step {
			ch <- [2]x{xfn(i),yfn(i)}
		}
		close(ch)
	}()
	return ch
}


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
		panic("circle too small")
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

func  (d Divide) Arc(x1,y1,rx,ry x, a float64, large,sweep bool, x2,y2 x)  <-chan [2]x{
	// if ellipse too small expand to just fit, which will depend on angle, uggg.
	// TODO for rx!=ry translate/rotate and squash, then do below then reverse transform on every point. 
	if rx==ry{
		// much simpler, just a circle, angle redundant
		var cx,cy,a1,a2 float64
		if large == sweep {
			cx,cy= centreOfCircle(x1,y1,rx,x2,y2)
		}else{
			cx,cy= centreOfCircle(x2,y2,rx,x1,y1)
		}
		fmt.Println(cx,cy)
		a1,a2=math.Atan2(float64(x1)-cx,float64(y1)-cy),math.Atan2(float64(x2)-cx,float64(y2)-cy)
		fmt.Println(a1,a2)
		if !sweep {
			a1,a2=a2,a1+math.Pi*2
		}
		fmt.Println("Angles:",a1,a2)
		// scale divisions so you get, somewhat, consistent side angles
		halfDivisions:=int8(math.Abs(float64(uint8(1)<<d)*(a2-a1)/math.Pi))+1
		ocwr,_:=offsetRotaters(cx,cy,(a2-a1)*.5/float64(halfDivisions))
		dx,dy:= float64(x1),float64(y1)
		
		step:=Divide(dividerMax>>d)
		ch:=make(chan [2]x,halfDivisions<<1)
		var li Divide
		go func(){
			for i := step-1; li<i ; li,i=i,i+step {
				dx,dy:=ocwr(dx,dy)
				fmt.Println(dx,dy)
				//s[i]=p.Line(x(dx),x(dy),x(ex),x(ey))
				ch <- [2]x{x(dx),x(dy)}
			}
			close(ch)
		}()
		//s[len(s)-1]=p.Line(x(dx),x(dy),x2,y2)
		return ch
	}

	return nil
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

/* run: args="" Thu 11 Jun 23:57:19 BST 2020 go version go1.14.3 linux/amd64
Thu 11 Jun 23:57:20 BST 2020
*/
