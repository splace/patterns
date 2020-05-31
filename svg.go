package patterns

import "fmt"
import "io"

type Drawer interface{
	Draw(*Brush)Pattern
}

type Path []Drawer

// draw a path using the provided brush
func (p Path) Draw(b *Brush)(c Composite) {
	for _,s:=range(p){
		d:=s.Draw(b)
		// draw can modify the brush without producing a pattern, no need to add these to the pattern
		if d==nil{continue}
		c=append(c,d)
	}
	return
}

type MoveTo []x

func (s MoveTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.MoveTo(s[0],s[1])
	return nil
}

type MoveToRelative []x

func (s MoveToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.MoveTo(s[0],s[1])
	return nil
}

type LineTo []x

func (s LineTo) Draw(b *Brush)Pattern{
	b.Relative=false
	return b.LineTo(s[0],s[1])
}

type LineToRelative []x

func (s LineToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	return b.LineTo(s[0],s[1])
}

type VeticalLineTo []x

func (s VeticalLineTo) Draw(b *Brush)Pattern{
	b.Relative=false
	return b.LineToVertical(s[0])
}

type VeticalLineToRelative []x

func (s VeticalLineToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	return b.LineToVertical(s[0])
}

type HorizontalLineTo []x

func (s HorizontalLineTo) Draw(b *Brush)Pattern{
	b.Relative=false
	return b.LineToHorizontal(s[0])
}

type HorizontalLineToRelative []x

func (s HorizontalLineToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	return b.LineToHorizontal(s[0])
}


type Close struct{}

func (s Close) Draw(b *Brush)Pattern{
	if b.x==b.sx && b.y==b.sy {return nil}
	return b.LineClose()
}

type QuadraticBezierTo []x

func (s QuadraticBezierTo) Draw(b *Brush)Pattern{
	b.Relative=false
	return b.QuadraticBezierTo(s[0],s[1],s[2],s[3])
}

type SmoothQuadraticBezierTo []x

func (s SmoothQuadraticBezierTo) Draw(b *Brush)Pattern{
	b.Relative=false
	return b.QuadraticBezierTo(b.x+(s[0]-s[2]),b.y+(s[1]-s[3]),s[4],s[5])
}

type QuadraticBezierToRelative []x

func (s QuadraticBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	return b.QuadraticBezierTo(s[0],s[1],s[2],s[3])
}

type SmoothQuadraticBezierToRelative []x

func (s SmoothQuadraticBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	return b.QuadraticBezierTo((s[0]-s[2]),(s[1]-s[3]),s[4],s[5])
}


type CubicBezierTo []x

func (s CubicBezierTo) Draw(b *Brush)Pattern{
	b.Relative=false
	return b.CubicBezierTo(s[0],s[1],s[2],s[3],s[4],s[5])
}

type SmoothCubicBezierTo []x

func (s SmoothCubicBezierTo) Draw(b *Brush)Pattern{
	b.Relative=false
	if len(s)==4{
		return b.CubicBezierTo(s[0],s[1],s[0],s[1],s[2],s[3])
	}
	return b.CubicBezierTo(b.x+(s[2]-s[0]),b.y+(s[3]-s[1]),s[4],s[5],s[6],s[7])
}

type CubicBezierToRelative []x

func (s CubicBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	return b.CubicBezierTo(s[0],s[1],s[2],s[3],s[4],s[5])
}


type SmoothCubicBezierToRelative []x

func (s SmoothCubicBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	if len(s)==4{
		return b.CubicBezierTo(s[0],s[1],s[0],s[1],s[2],s[3])
	}
	return b.CubicBezierTo((s[2]-s[0]),(s[3]-s[1]),s[4],s[5],s[6],s[7])
}

type ArcTo []x

func (s ArcTo) Draw(b *Brush)Pattern{
	b.Relative=false
	return b.ArcTo(s[0],s[1],float64(s[2])/unitX,s[3]!=0,s[4]!=0,s[5],s[6])
}

type ArcToRelative []x

func (s ArcToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	return b.ArcTo(s[0],s[1],float64(s[2])/unitX,s[3]!=0,s[4]!=0,s[5],s[6])
}



// Scan a path into a slice of Drawers, all using the same slice of x's for their coordinates.
// the x's are a one-to-one mapping of the values in the provided texture path representation. (enabling regeneration of textual form exactly, not just equivalent)
// AND each Drawer, even smoothed, is independent (relatively) of any others, by vertue of overlapping slices from the x's.
// Notice: above features cause an odd edge case; long sequences of smoothed quadratic Beziers need to look all the way back to a non-smoothed Quadratic Bezier,to find their control point, which then needs to be reflected back through all points!.(with potential relative/absolute switching along the way)
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
			case 'z','Z':
				*p=append(*p,Close{})
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
				*p=append(*p,VeticalLineTo(xs[len(xs)-1:]))
			case 'v':
				xs=append(xs,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,VeticalLineToRelative(xs[len(xs)-1:]))
			
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
				return fmt.Errorf("Not supported")
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				switch (*p)[len(*p)-1].(type){
				case QuadraticBezierTo:
					*p=append(*p,QuadraticBezierTo(xs[len(xs)-4:]))

				case QuadraticBezierToRelative:
					//xs[len(xs)-2]
					//xs[len(xs)-1]
					//pass
				case SmoothQuadraticBezierTo:
				case SmoothQuadraticBezierToRelative:
					*p=append(*p,SmoothQuadraticBezierTo(xs[len(xs)-6:]))
				default:
					// add duplicate point for control point
					xs=append(xs,xs[len(xs)-2],xs[len(xs)-1])
					*p=append(*p,QuadraticBezierTo(xs[len(xs)-4:]))
				}
			case 't': // smooth quadratic Bézier curveto relative
				return fmt.Errorf("Not supported")
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				switch (*p)[len(*p)-1].(type){
				case QuadraticBezierTo:
				case QuadraticBezierToRelative:
					//xs[len(xs)-2]
					//xs[len(xs)-1]
					//pass
				case SmoothQuadraticBezierTo:
				case SmoothQuadraticBezierToRelative:
					*p=append(*p,SmoothQuadraticBezierTo(xs[len(xs)-6:]))
				default:
				}
			case 'S': // smooth cubic Bézier curve
				xs=append(xs,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				switch (*p)[len(*p)-1].(type){
				case CubicBezierTo,SmoothCubicBezierTo,CubicBezierToRelative,SmoothCubicBezierToRelative:
					*p=append(*p,SmoothCubicBezierTo(xs[len(xs)-8:]))  // include back-reference to last 4 paramters of these commands
				//case MoveToRelative,LineToRelative,HorizontalLineToRelative,VeticalLineToRelative,QuadraticBezierToRelative,SmoothQuadraticBezierToRelative,CubicBezierToRelative,SmoothCubicBezierToRelative,ArcToRelative: 
				default:
					// no first control point to back reference duplicate first point.
					*p=append(*p,SmoothCubicBezierTo(xs[len(xs)-4:]))
				}
			case 's': // smooth cubic Bézier curve relative
				xs=append(xs,0,0,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-4],&xs[len(xs)-3],&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				switch (*p)[len(*p)-1].(type){
				case CubicBezierTo,SmoothCubicBezierTo,CubicBezierToRelative,SmoothCubicBezierToRelative:
					*p=append(*p,SmoothCubicBezierToRelative(xs[len(xs)-8:]))
				default:
					*p=append(*p,SmoothCubicBezierToRelative(xs[len(xs)-4:]))
				}
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
				// numeric parameter so repeat switch using previous command
				state.UnreadRune()
				c=lc
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
