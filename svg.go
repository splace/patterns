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

type CubicBezierToRelative []x

func (s CubicBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	return b.CubicBezierTo(s[0],s[1],s[2],s[3],s[4],s[5])
}

// TODO read though comma-> space filter

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
				// if lc=='z' or'Z' then new subpath for lineend style (non supported)
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				*p=append(*p,MoveTo(xs[len(xs)-2:]))
			case 'm':
				// if lc=='z' or'Z' then new subpath for lineend style (non supported)
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
			// smooth curves need back-referenced control point.
			case 'T': // smooth quadratic Bézier curveto
				return fmt.Errorf("Not supported")
				xs=append(xs,0,0)
				_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
				if err!=nil{return err}
				switch lc{
				case 't':
					fallthrough
				case 'q':
					//xs[len(xs)-2]
					//xs[len(xs)-1]
					//pass
				case 'T':
					fallthrough
				case 'Q':
					*p=append(*p,SmoothQuadraticBezierTo(xs[len(xs)-6:]))
				default:
					// duplicate point for control point
					xs=append(xs,xs[len(xs)-2],xs[len(xs)-1])
					*p=append(*p,QuadraticBezierTo(xs[len(xs)-4:]))
				}
			case 't': // smooth quadratic Bézier curveto relative
				return fmt.Errorf("Not supported")
				switch lc{
				case 'Q','T','q','t':
					xs=append(xs,0,0)
					_,err=fmt.Fscan(state,&xs[len(xs)-2],&xs[len(xs)-1])
					if err!=nil{return err}
					*p=append(*p,SmoothQuadraticBezierToRelative(xs[len(xs)-6:]))
				default:
				}
			case 'S': // smooth cubic Bézier curve
				return fmt.Errorf("Not supported")
				switch lc{
				case 'S','C':
				case 's','c':
				default:
				}
			case 's': // smooth cubic Bézier curve relative
				return fmt.Errorf("Not supported")
				switch lc{
				case 'S','C':
				case 's','c':
				default:
				}
			case 'A': // elliptical Arc
				return fmt.Errorf("Not supported")
			case 'a': // elliptical Arc relative
				return fmt.Errorf("Not supported")
		
			case '0','1','2','3','4','5','6','7','8','9','.','-','+':
				// numeric parameter so repeat switch using previous command
				state.UnreadRune()
				c=lc
				continue
			default:
				return fmt.Errorf("Unknown command:%v",c)
			}
			break
		}
		lc=c
	}
	return nil
}
