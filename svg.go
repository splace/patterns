package patterns

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

// a brush is a Pen that also stores control points for smoothed bezier segments
type Brush struct {
	PenPath
	dqcx, dqcy  x 
	dccx, dccy  x 
}

func NewBrush(n Nib) *Brush{
	return &Brush{PenPath:PenPath{Pen:Pen{Nib:n}}}
}

type MoveTo []x

func (s MoveTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.MoveTo(s[0],s[1])
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return nil
}

type MoveToRelative []x

func (s MoveToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.MoveTo(s[0],s[1])
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return nil
}

type LineTo []x

func (s LineTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineTo(s[0],s[1])
}

type LineToRelative []x

func (s LineToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineTo(s[0],s[1])
}

type VerticalLineTo []x

func (s VerticalLineTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineToVertical(s[0])
}

type VerticalLineToRelative []x

func (s VerticalLineToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineToVertical(s[0])
}

type HorizontalLineTo []x

func (s HorizontalLineTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineToHorizontal(s[0])
}

type HorizontalLineToRelative []x

func (s HorizontalLineToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineToHorizontal(s[0])
}


type Close struct{}

func (s Close) Draw(b *Brush)Pattern{
	//if b.x==b.sx && b.y==b.sy {return nil}
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineClose()
}

type CloseRelative struct{}

func (s CloseRelative) Draw(b *Brush)Pattern{
	return Close(s).Draw(b)
}

//var quadraticControlx, quadraticControly x

type QuadraticBezierTo []x

func (s QuadraticBezierTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.dqcx, b.dqcy = s[2]-s[0], s[3]-s[1]
	b.dccx, b.dccy = 0,0
	return b.QuadraticBezierTo(s[0],s[1],s[2],s[3])
}

type SmoothQuadraticBezierTo []x

func (s SmoothQuadraticBezierTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.dccx, b.dccy = 0,0
	b.dqcx+=b.x
	b.dqcy+=b.y
	p:=b.QuadraticBezierTo(b.dqcx,b.dqcy,s[0],s[1])
	b.dqcx,b.dqcy=s[0]-b.dqcx,s[1]-b.dqcy
	return p
}

type QuadraticBezierToRelative []x

func (s QuadraticBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.dccx, b.dccy = 0,0
	b.dqcx, b.dqcy = s[2]-s[0], s[3]-s[1]
	return b.QuadraticBezierTo(s[0],s[1],s[2],s[3])
}

type SmoothQuadraticBezierToRelative []x

func (s SmoothQuadraticBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.dccx, b.dccy = 0,0
	p:=b.QuadraticBezierTo(b.dqcx,b.dqcy,s[0],s[1])
	b.dqcx, b.dqcy = s[0]-b.dqcx, s[1]-b.dqcy
	return p
}

//var cubicControlx, cubicControly x

type CubicBezierTo []x

func (s CubicBezierTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = s[4]-s[2], s[5]-s[3]
	return b.CubicBezierTo(s[0],s[1],s[2],s[3],s[4],s[5])
}

type SmoothCubicBezierTo []x

func (s SmoothCubicBezierTo) Draw(b *Brush)Pattern{
	b.Relative=false
	b.dqcx, b.dqcy = 0,0
	p:=b.CubicBezierTo(b.dccx+b.x,b.dccy+b.y,s[0],s[1],s[2],s[3])
	b.dccx, b.dccy = s[2]-s[0], s[3]-s[1]
	return p
}

type CubicBezierToRelative []x

func (s CubicBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = s[4]-s[2], s[5]-s[3]
	return b.CubicBezierTo(s[0],s[1],s[2],s[3],s[4],s[5])
}


type SmoothCubicBezierToRelative []x

func (s SmoothCubicBezierToRelative) Draw(b *Brush)Pattern{
	b.Relative=true
	b.dqcx, b.dqcy = 0,0
	p:= b.CubicBezierTo(b.dccx,b.dccy,s[0],s[1],s[2],s[3])
	b.dccx, b.dccy = s[2]-s[0], s[3]-s[1]
	return p
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

