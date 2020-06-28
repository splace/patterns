package patterns

import "math"

type Drawer interface{
	Draw(*Brush)Pattern
}

// a Path is a collection of Drawers that uses the same Brush to Draw its contained Drawers in order.
// it is itself a Drawer so a Path can be an ordered collection of (sub) Paths.
// Notice: sub-Paths are drawn with the same Brush so relative Drawers carry on from the end of the previous sub-Path.
type Path []Drawer

// draw a path using the provided brush
func (p Path) Draw(b *Brush) Pattern {
	var c Composite
	for _,s:=range(p){
		if d:=s.Draw(b);d!=nil{
			c=append(c,d)
		}
	}
	return c
}

// a brush is a Pen that stores control points to allow generation of smoothed bezier segments
type Brush struct {
	PenPath
	dqcx, dqcy  x 
	dccx, dccy  x 
}

func NewBrush(n Nib) *Brush{
	return &Brush{PenPath:PenPath{Pen:Pen{Nib:n}}}
}

func NewFacettedBrush(width x,f filler, d uint8) *Brush{
	return &Brush{PenPath:PenPath{Pen:Pen{Nib:Facetted{LineNib:LineNib{width,f.fill()},CurveDivision:d},Marker:Shrunk{Disc(Filling(f.fill())),float32(2*unitX/width)}}}}
}


type MoveTo []x

func (s MoveTo) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return 	b.MoveTo(s[0],s[1])
}

type MoveToRelative []x

func (s MoveToRelative) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return 	b.MoveTo(b.PenPath.Pen.x+s[0],b.PenPath.Pen.y+s[1])
}

type LineTo []x

func (s LineTo) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineTo(s[0],s[1])
}

type LineToRelative []x

func (s LineToRelative) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineTo(b.PenPath.Pen.x+s[0],b.PenPath.Pen.y+s[1])
}

type VerticalLineTo []x

func (s VerticalLineTo) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineToVertical(s[0])
}

type VerticalLineToRelative []x

func (s VerticalLineToRelative) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineToVertical(b.PenPath.Pen.y+s[0])
}

type HorizontalLineTo []x

func (s HorizontalLineTo) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineToHorizontal(s[0])
}

type HorizontalLineToRelative []x

func (s HorizontalLineToRelative) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineToHorizontal(b.PenPath.Pen.x+s[0])
}


type Close struct{}

func (s Close) Draw(b *Brush)Pattern{
	//if b.x==b.sx && b.y==b.sy {return nil}
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = 0,0
	return b.LineClose()
}

type CloseRelative struct{
	Close
}

type QuadraticBezierTo []x

func (s QuadraticBezierTo) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = s[2]-s[0], s[3]-s[1]
	b.dccx, b.dccy = 0,0
	return b.QuadraticBezierTo(s[0],s[1],s[2],s[3])
}

type SmoothQuadraticBezierTo []x

func (s SmoothQuadraticBezierTo) Draw(b *Brush)Pattern{
	b.dccx, b.dccy = 0,0
	b.dqcx+=b.PenPath.Pen.x
	b.dqcy+=b.PenPath.Pen.y
	p:=b.QuadraticBezierTo(b.dqcx,b.dqcy,s[0],s[1])
	b.dqcx,b.dqcy=s[0]-b.dqcx,s[1]-b.dqcy
	return p
}

type QuadraticBezierToRelative []x

func (s QuadraticBezierToRelative) Draw(b *Brush)Pattern{
	b.dccx, b.dccy = 0,0
	b.dqcx, b.dqcy = s[2]-s[0], s[3]-s[1]
	return b.QuadraticBezierTo(b.PenPath.Pen.x+s[0],b.PenPath.Pen.y+s[1],b.PenPath.Pen.x+s[2],b.PenPath.Pen.y+s[3])
}

type SmoothQuadraticBezierToRelative []x

func (s SmoothQuadraticBezierToRelative) Draw(b *Brush)Pattern{
	b.dccx, b.dccy = 0,0
	p:=b.QuadraticBezierTo(b.PenPath.Pen.x+b.dqcx,b.PenPath.Pen.y+b.dqcy,b.PenPath.Pen.x+s[0],b.PenPath.Pen.y+s[1])
	b.dqcx, b.dqcy = s[0]-b.dqcx, s[1]-b.dqcy
	return p
}


type CubicBezierTo []x

func (s CubicBezierTo) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = s[4]-s[2], s[5]-s[3]
	return b.CubicBezierTo(s[0],s[1],s[2],s[3],s[4],s[5])
}

type SmoothCubicBezierTo []x

func (s SmoothCubicBezierTo) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	p:=b.CubicBezierTo(b.dccx+b.PenPath.Pen.x,b.dccy+b.PenPath.Pen.y,s[0],s[1],s[2],s[3])
	b.dccx, b.dccy = s[2]-s[0], s[3]-s[1]
	return p
}

type CubicBezierToRelative []x

func (s CubicBezierToRelative) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	b.dccx, b.dccy = s[4]-s[2], s[5]-s[3]
	return b.CubicBezierTo(b.PenPath.Pen.x+s[0],b.PenPath.Pen.y+s[1],b.PenPath.Pen.x+s[2],b.PenPath.Pen.y+s[3],b.PenPath.Pen.x+s[4],b.PenPath.Pen.y+s[5])
}


type SmoothCubicBezierToRelative []x

func (s SmoothCubicBezierToRelative) Draw(b *Brush)Pattern{
	b.dqcx, b.dqcy = 0,0
	p:= b.CubicBezierTo(b.dccx,b.dccy,b.PenPath.Pen.x+s[0],b.PenPath.Pen.y+s[1],b.PenPath.Pen.x+s[2],b.PenPath.Pen.y+s[3])
	b.dccx, b.dccy = s[2]-s[0], s[3]-s[1]
	return p
}

type ArcTo []x

func (s ArcTo) Draw(b *Brush)Pattern{
	return b.ArcTo(s[0],s[1],float64(s[2])/unitX*math.Pi/180,s[3]!=0,s[4]!=0,s[5],s[6])
}

type ArcToRelative []x

func (s ArcToRelative) Draw(b *Brush)Pattern{
	return b.ArcTo(s[0],s[1],float64(s[2])/unitX*math.Pi/180,s[3]!=0,s[4]!=0,b.PenPath.Pen.x+s[5],b.PenPath.Pen.y+s[6])
}
