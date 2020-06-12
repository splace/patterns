package patterns

import "fmt"
import "io"
import "strings"


func (p Path) String()string {
	b :=new(strings.Builder)
	var ls Drawer
	for _,s:=range(p){
		switch st:=s.(type){
		case MoveTo:
			fmt.Fprintf(b,"\nM%v,%v",st[0],st[1])
		case MoveToRelative:
			fmt.Fprintf(b,"\nm%v,%v",st[0],st[1])
		case LineTo:
			switch ls.(type){
			case LineTo,MoveTo:
				fmt.Fprintf(b," %v,%v",st[0],st[1])
			default:
				fmt.Fprintf(b,"\nL%v,%v",st[0],st[1])
			}
		case LineToRelative:
			switch ls.(type){
			case LineToRelative,MoveToRelative:
				fmt.Fprintf(b," %v,%v",st[0],st[1])
			default:
				fmt.Fprintf(b,"\nl%v,%v",st[0],st[1])
			}
		case VerticalLineTo:
			fmt.Fprintf(b,"\nV%v",st[0])
		case VerticalLineToRelative:
			fmt.Fprintf(b,"\nv%v",st[0])
		case HorizontalLineTo:
			fmt.Fprintf(b,"\nH%v",st[0])
		case HorizontalLineToRelative:
			fmt.Fprintf(b,"\nh%v",st[0])
		case CloseRelative:
			fmt.Fprintf(b,"\nz")
		case Close:
			fmt.Fprintf(b,"\nZ")
		case QuadraticBezierTo:
			switch ls.(type){
			case QuadraticBezierTo:
				fmt.Fprintf(b," %v,%v %v,%v",st[0],st[1],st[2],st[3])
			default:
				fmt.Fprintf(b,"\nQ%v,%v %v,%v",st[0],st[1],st[2],st[3])
			}
		case SmoothQuadraticBezierTo:
			switch ls.(type){
			case SmoothQuadraticBezierTo:
				fmt.Fprintf(b," %v,%v",st[0],st[1])
			default:
				fmt.Fprintf(b,"\nT%v,%v",st[0],st[1])
			}
		case QuadraticBezierToRelative:
			switch ls.(type){
			case QuadraticBezierToRelative:
				fmt.Fprintf(b," %v,%v %v,%v",st[0],st[1],st[2],st[3])
			default:
				fmt.Fprintf(b,"\nq%v,%v %v,%v",st[0],st[1],st[2],st[3])
			}
		case SmoothQuadraticBezierToRelative:
			switch ls.(type){
			case SmoothQuadraticBezierToRelative:
				fmt.Fprintf(b," %v,%v",st[0],st[1])
			default:
				fmt.Fprintf(b,"\nt%v,%v",st[0],st[1])
			}
		case CubicBezierTo:
			switch ls.(type){
			case CubicBezierTo:
				fmt.Fprintf(b," %v,%v %v,%v %v,%v",st[0],st[1],st[2],st[3],st[4],st[5])
			default:
				fmt.Fprintf(b,"\nC%v,%v %v,%v %v,%v",st[0],st[1],st[2],st[3],st[4],st[5])
			}
		case SmoothCubicBezierTo:
			switch ls.(type){
			case CubicBezierTo:
				fmt.Fprintf(b," %v,%v %v,%v",st[0],st[1],st[2],st[3])
			default:
				fmt.Fprintf(b,"\nS%v,%v %v,%v",st[0],st[1],st[2],st[3])
			}
		case CubicBezierToRelative:
			switch ls.(type){
			case CubicBezierToRelative:
				fmt.Fprintf(b," %v,%v %v,%v %v,%v",st[0],st[1],st[2],st[3],st[4],st[5])
			default:
				fmt.Fprintf(b,"\nc%v,%v %v,%v %v,%v",st[0],st[1],st[2],st[3],st[4],st[5])
			}
		case SmoothCubicBezierToRelative:
			switch ls.(type){
			case SmoothCubicBezierToRelative:
				fmt.Fprintf(b," %v,%v %v,%v",st[0],st[1],st[2],st[3])
			default:
				fmt.Fprintf(b,"\ns%v,%v %v,%v",st[0],st[1],st[2],st[3])
			}
		case ArcTo:
			switch ls.(type){
			case ArcTo:
				fmt.Fprintf(b," %v,%v %v %v %v %v,%v",st[0],st[1],st[2],st[3],st[4],st[5],st[6])
			default:
				fmt.Fprintf(b,"\nA%v,%v %v %v %v %v,%v",st[0],st[1],st[2],st[3],st[4],st[5],st[6])
			}
		case ArcToRelative:
			switch ls.(type){
			case ArcToRelative:
				fmt.Fprintf(b," %v,%v %v %v %v %v,%v",st[0],st[1],st[2],st[3],st[4],st[5],st[6])
			default:
				fmt.Fprintf(b,"\na%v,%v %v %v %v %v,%v",st[0],st[1],st[2],st[3],st[4],st[5],st[6])
			}
		}
		ls=s
	}
	return b.String()[1:]
}

// skip repeated commands and leading zero/space/comma when leading point or neg
type CompactPath Path

type presep x
func (cx presep) String()(s string){
	s=fmt.Sprint(x(cx))
	switch s[0]{
	case '0':
		if len(s)>1 {return ","+s[1:]}
		return ","+s
	case '1','2','3','4','5','6','7','8','9':
		return ","+s
	}	
	return 
}

type compactx x
func (cx compactx) String()(s string){
	s=fmt.Sprint(x(cx))
	if len(s)>1 && s[0]=='0' {return s[1:]}
	return 
}

func (p CompactPath) String()string {
	b :=new(strings.Builder)
	var ls Drawer
	for _,s:=range(p){
		switch st:=s.(type){
		case MoveTo:
			fmt.Fprintf(b,"M%v%v",compactx(st[0]),presep(st[1]))
		case MoveToRelative:
			fmt.Fprintf(b,"m%v%v",compactx(st[0]),presep(st[1]))
		case LineTo:
			switch ls.(type){
			case LineTo,MoveTo:
				fmt.Fprintf(b,"%v%v",presep(st[0]),presep(st[1]))
			default:
				fmt.Fprintf(b,"L%v%v",compactx(st[0]),presep(st[1]))
			}
		case LineToRelative:
			switch ls.(type){
			case LineToRelative,MoveToRelative:
				fmt.Fprintf(b,"%v%v",presep(st[0]),presep(st[1]))
			default:
				fmt.Fprintf(b,"l%v%v",compactx(st[0]),presep(st[1]))
			}
		case VerticalLineTo:
			fmt.Fprintf(b,"V%v",compactx(st[0]))
		case VerticalLineToRelative:
			fmt.Fprintf(b,"v%v",compactx(st[0]))
		case HorizontalLineTo:
			fmt.Fprintf(b,"H%v",compactx(st[0]))
		case HorizontalLineToRelative:
			fmt.Fprintf(b,"h%v",compactx(st[0]))
		case CloseRelative:
			fmt.Fprintf(b,"z")
		case Close:
			fmt.Fprintf(b,"Z")
		case QuadraticBezierTo:
			switch ls.(type){
			case QuadraticBezierTo:
				fmt.Fprintf(b,"%v%v%v%v",presep(st[0]),presep(st[1]),presep(st[2]),presep(st[3]))
			default:
				fmt.Fprintf(b,"Q%v%v%v%v",compactx(st[0]),presep(st[1]),presep(st[2]),presep(st[3]))
			}
		case SmoothQuadraticBezierTo:
			switch ls.(type){
			case SmoothQuadraticBezierTo:
				fmt.Fprintf(b,"%v%v",presep(st[0]),presep(st[1]))
			default:
				fmt.Fprintf(b,"T%v%v",compactx(st[0]),presep(st[1]))
			}
		case QuadraticBezierToRelative:
			switch ls.(type){
			case QuadraticBezierToRelative:
				fmt.Fprintf(b,"%v%v%v%v",presep(st[0]),presep(st[1]),presep(st[2]),presep(st[3]))
			default:
				fmt.Fprintf(b,"q%v%v%v%v",compactx(st[0]),presep(st[1]),presep(st[2]),presep(st[3]))
			}
		case SmoothQuadraticBezierToRelative:
			switch ls.(type){
			case SmoothQuadraticBezierToRelative:
				fmt.Fprintf(b,"%v%v",presep(st[0]),presep(st[1]))
			default:
				fmt.Fprintf(b,"t%v%v",compactx(st[0]),presep(st[1]))
			}
		case CubicBezierTo:
			switch ls.(type){
			case CubicBezierTo:
				fmt.Fprintf(b,"%v%v%v%v%v%v",presep(st[0]),presep(st[1]),presep(st[2]),presep(st[3]),presep(st[4]),presep(st[5]))
			default:
				fmt.Fprintf(b,"C%v%v%v%v%v%v",compactx(st[0]),presep(st[1]),presep(st[2]),presep(st[3]),presep(st[4]),presep(st[5]))
			}
		case SmoothCubicBezierTo:
			switch ls.(type){
			case CubicBezierTo:
				fmt.Fprintf(b,"%v%v%v%v",presep(st[0]),presep(st[1]),presep(st[2]),presep(st[3]))
			default:
				fmt.Fprintf(b,"S%v%v%v%v",compactx(st[0]),presep(st[1]),presep(st[2]),presep(st[3]))
			}
		case CubicBezierToRelative:
			switch ls.(type){
			case CubicBezierToRelative:
				fmt.Fprintf(b,"%v%v%v%v%v%v",presep(st[0]),presep(st[1]),presep(st[2]),presep(st[3]),presep(st[4]),presep(st[5]))
			default:
				fmt.Fprintf(b,"c%v%v%v%v%v%v",compactx(st[0]),presep(st[1]),presep(st[2]),presep(st[3]),presep(st[4]),presep(st[5]))
			}
		case SmoothCubicBezierToRelative:
			switch ls.(type){
			case SmoothCubicBezierToRelative:
				fmt.Fprintf(b,"%v%v%v%v",presep(st[0]),presep(st[1]),presep(st[2]),presep(st[3]))
			default:
				fmt.Fprintf(b,"s%v%v%v%v",compactx(st[0]),presep(st[1]),presep(st[2]),presep(st[3]))
			}
		case ArcTo:
			switch ls.(type){
			case ArcTo:
				fmt.Fprintf(b,"%v%v%v%v%v%v%v",presep(st[0]),presep(st[1]),presep(st[2]),presep(st[3]),presep(st[4]),presep(st[5]),presep(st[6]))
			default:
				fmt.Fprintf(b,"A%v%v%v%v%v%v%v",compactx(st[0]),presep(st[1]),presep(st[2]),presep(st[3]),presep(st[4]),presep(st[5]),presep(st[6]))
			}
		case ArcToRelative:
			switch ls.(type){
			case ArcToRelative:
				fmt.Fprintf(b,"%v%v%v%v%v%v%v",presep(st[0]),presep(st[1]),presep(st[2]),presep(st[3]),presep(st[4]),presep(st[5]),presep(st[6]))
			default:
				fmt.Fprintf(b,"a%v%v%v%v%v%v%v",compactx(st[0]),presep(st[1]),presep(st[2]),presep(st[3]),presep(st[4]),presep(st[5]),presep(st[6]))
			}
		}
		ls=s
	}
	return b.String()
}



// Scan a path into a slice of Drawers, all using slices of x's for the same array for their coordinates.
// the x's are a one-to-one mapping of the values in the provided texture path representation. (enabling regeneration of textual form exactly, not just equivalent)
// AND each Drawer, even smoothed, is independent (relatively) of any others, by vertue of overlapping slices from the x's.
// Note: above features prevent support for long sequences of smoothed quadratic Beziers, since they are dependent on a previous onn-smoothed Quadratic Bezier an unlimited number of segments previously, to find their control point.
// paths containing these should be pre-parsed to remove smoothed segments, really representing wrist friendly hand-editing compression.
// 
func (p *Path) Scan(state fmt.ScanState,r rune) (err error){
	var xs []x
	var lc,c rune
	for {
		state.SkipSpace()
		c,_,err=state.ReadRune()
		if err!=nil{
			if err==io.EOF {return nil}
			return err
		}
		if c==',' {
			state.SkipSpace()
			c,_,err=state.ReadRune()
			if err!=nil{
				if err==io.EOF {return nil}
				return err
			}
		}
		
		for {
			switch c{
			//case '?'
			//Segment-completing close path operations
			// close sub-path (so not lineend marker) 
			// if curve followed by close AND coords textually identical to start coords, then close path without line
			
			case 'M':
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,MoveTo(xs[len(xs)-2:]))
			case 'm':
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,MoveToRelative(xs[len(xs)-2:]))
			case 'L':
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,LineTo(xs[len(xs)-2:]))
			case 'l':
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,LineToRelative(xs[len(xs)-2:]))
			case 'Z':
				*p=append(*p,Close{})
			case 'z':
				*p=append(*p,CloseRelative{})
			case 'H':
				xs=append(xs,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,HorizontalLineTo(xs[len(xs)-1:]))
			case 'h':
				xs=append(xs,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,HorizontalLineToRelative(xs[len(xs)-1:]))
			case 'V':
				xs=append(xs,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,VerticalLineTo(xs[len(xs)-1:]))
			case 'v':
				xs=append(xs,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,VerticalLineToRelative(xs[len(xs)-1:]))
			
			// TODO  for curves create their own paths?	
			// or use transform to get quadratic, and create cubic from a series f these?
			// or both ways using rune selection?
			case 'Q': // quadratic Bézier curve
				xs=append(xs,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,QuadraticBezierTo(xs[len(xs)-4:]))
			case 'q': // quadratic Bézier curve relative
				xs=append(xs,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,QuadraticBezierToRelative(xs[len(xs)-4:]))
			case 'C': // cubic Bézier curve 
				xs=append(xs,0,0,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-6],&xs[len(xs)-5],&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,CubicBezierTo(xs[len(xs)-6:]))
			case 'c': // cubic Bézier curve relative
				xs=append(xs,0,0,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-6],&xs[len(xs)-5],&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,CubicBezierToRelative(xs[len(xs)-6:]))

			// smooth curves use back-referenced control points where possible..
			case 'T': // smooth quadratic Bézier curveto
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,SmoothQuadraticBezierTo(xs[len(xs)-2:])) 
			case 't': // smooth quadratic Bézier curveto relative
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,SmoothQuadraticBezierToRelative(xs[len(xs)-2:])) 

//			case 'T': // smooth quadratic Bézier curveto
//				// here,  unlike cubic, previous control point can come from previous to that, so propagate backward an unlimited number of segments.
//				xs=append(xs,0,0)
//				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
//				if err!=nil{return err}
//				switch (*p)[len(*p)-1].(type){
//				case QuadraticBezierTo,QuadraticBezierToRelative:
//					*p=append(*p,SmoothQuadraticBezierTo(xs[len(xs)-6:]))
//				case SmoothQuadraticBezierTo,SmoothQuadraticBezierToRelative:
//					// search back to find original control? relative/abs along the way.
//					// has to be QuadraticBezierTo,QuadraticBezierToRelative some where back.
//					switch (*p)[len(*p)-2].(type){
//					case SmoothQuadraticBezierTo,SmoothQuadraticBezierToRelative:
//						// TODO search backward to find origin of control point and insert into points.
//						return fmt.Errorf("Not supported: triple smooth quadratic sequence.")
//					default:
//						*p=append(*p,SmoothQuadraticBezierTo(xs[len(xs)-6:]))
//					}
//				default:
//					*p=append(*p,SmoothQuadraticBezierTo(xs[len(xs)-2:]))  // this the same as lineto?
//				}
//			case 't': // smooth quadratic Bézier curveto relative
//				xs=append(xs,0,0)
//				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
//				if err!=nil{return err}
//				switch (*p)[len(*p)-1].(type){
//				case QuadraticBezierTo,QuadraticBezierToRelative:
//					*p=append(*p,SmoothQuadraticBezierToRelative(xs[len(xs)-6:]))
//				case SmoothQuadraticBezierTo,SmoothQuadraticBezierToRelative:
//					switch (*p)[len(*p)-2].(type){
//					case SmoothQuadraticBezierTo,SmoothQuadraticBezierToRelative:
//						// TODO search backward to find origin of control point and insert new points.
//						return fmt.Errorf("Not supported: triple smooth quadratic sequence.")
//					default:
//						// wether quad, or not, same process due to fall back values.
//						*p=append(*p,SmoothQuadraticBezierToRelative(xs[len(xs)-6:]))
//					}
//				default:
//					*p=append(*p,SmoothQuadraticBezierToRelative(xs[len(xs)-2:])) 
//				}
				
//			case 'S': // smooth cubic Bézier curve
//				switch (*p)[len(*p)-1].(type){
//				case CubicBezierTo,SmoothCubicBezierTo,CubicBezierToRelative,SmoothCubicBezierToRelative:
//					xs=append(xs,xs[len(xs)-3]-xs[len(xs)-1],xs[len(xs)-4]-xs[len(xs)-2] , 0,0,0,0)
//				default:
//					xs=append(xs,xs[len(xs)-2],xs[len(xs)-1],0,0,0,0)
//					_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
//				}
//				if err!=nil{return err}
//				*p=append(*p,CubicBezierTo(xs[len(xs)-6:]))
//				err=Modified{}

			case 'S': // smooth cubic Bézier curve
				xs=append(xs,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,SmoothCubicBezierTo(xs[len(xs)-4:]))
			case 's': // smooth cubic Bézier curve relative
				xs=append(xs,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,SmoothCubicBezierToRelative(xs[len(xs)-4:]))

//			case 'S': // smooth cubic Bézier curve
//				xs=append(xs,0,0,0,0)
//				_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
//				if err!=nil{return err}
//				switch (*p)[len(*p)-1].(type){
//				case CubicBezierTo,SmoothCubicBezierTo,CubicBezierToRelative,SmoothCubicBezierToRelative:
//					*p=append(*p,SmoothCubicBezierTo(xs[len(xs)-8:]))  // include back-reference to last 4 paramters of these commands
//				//case MoveToRelative,LineToRelative,HorizontalLineToRelative,VeticalLineToRelative,QuadraticBezierToRelative,SmoothQuadraticBezierToRelative,CubicBezierToRelative,SmoothCubicBezierToRelative,ArcToRelative: 
//				default:
//					*p=append(*p,SmoothCubicBezierTo(xs[len(xs)-4:]))
//				}
//			case 's': // smooth cubic Bézier curve relative
//				xs=append(xs,0,0,0,0)
//				_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
//				if err!=nil{return err}
//				switch (*p)[len(*p)-1].(type){
//				case CubicBezierTo,SmoothCubicBezierTo,CubicBezierToRelative,SmoothCubicBezierToRelative:
//					*p=append(*p,SmoothCubicBezierToRelative(xs[len(xs)-8:]))
//				default:
//					*p=append(*p,SmoothCubicBezierToRelative(xs[len(xs)-4:]))
//				}
			case 'A': // elliptical Arc
				return fmt.Errorf("Not supported")
				xs=append(xs,0,0,0,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-7],&xs[len(xs)-6],&xs[len(xs)-5],&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if xs[len(xs)-3]!=0 && xs[len(xs)-3]!=1 || xs[len(xs)-4]!=0 && xs[len(xs)-4]!=1 {
					return fmt.Errorf("Arc flags not 0 or 1")
				}
				if err!=nil{return err}
				*p=append(*p,ArcTo(xs[len(xs)-7:]))
			case 'a': // elliptical Arc relative
				return fmt.Errorf("Not supported")
				xs=append(xs,0,0,0,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-7],&xs[len(xs)-6],&xs[len(xs)-5],&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				if xs[len(xs)-3]!=0 && xs[len(xs)-3]!=1 || xs[len(xs)-4]!=0 && xs[len(xs)-4]!=1 {
					return fmt.Errorf("Arc flags not 0 or 1")
				}
				*p=append(*p,ArcToRelative(xs[len(xs)-7:]))
		
			case '0','1','2','3','4','5','6','7','8','9','.','-','+':
				state.UnreadRune()
				switch lc {
				case 'M':
					c='L'
				case 'm':
					c='l'
				default:
					// numeric parameter so repeat switch using previous command
					c=lc
				}
				continue
			default:
				return fmt.Errorf("Unknown command:%v",string(c))
			}
			break
		}
		lc=c
	}
	return nil
}
