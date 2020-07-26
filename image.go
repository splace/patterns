package pattern

import (
	"image"
	"image/color"
//	"image/color/palette"
	"image/draw"
)

// a Depictor is an image.Image, missing a colormodel, so more general.
// embedded in one of the helper wrappers gets you an image.Image.
type Depictor interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

// simple visual Depiction of a Limited, always square because it uses MaxX for bounds, implements Depictor
type LimitedDepiction struct {
	Limited
	xsperpixel x
}

func (d LimitedDepiction) Bounds() image.Rectangle {
	max:=int(d.Limited.MaxX()/d.xsperpixel)
	return image.Rectangle{image.Pt(-max,-max),image.Pt(max,max)}
}

func (d LimitedDepiction) At(xp, yp int) color.Color {
	return d.at(x(xp)*d.xsperpixel, x(yp)*d.xsperpixel)
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero centred, width fitted
func NewLimitedDepiction(s Limited, dxX, dxY int) LimitedDepiction {
	return LimitedDepiction{s, xspp(dxX,dxY,s.MaxX()) }
}

func NewHalfLimitedDepiction(s Limited, dxX, dxY int) LimitedDepiction {
	return LimitedDepiction{Limiter{s,s.MaxX()/2}, xspp(dxX,dxY,s.MaxX()/2) }
}

// simple visual Depiction of a Limited using MaxX for bounds, so always square, implements Depictor
type TrimmedLimitedDepiction struct {
	LimitedDepiction
	dx,dy int
}

func (d TrimmedLimitedDepiction) Bounds() image.Rectangle {
	db:=d.LimitedDepiction.Bounds()
	db.Min.X+=d.dx
	db.Min.Y+=d.dy
	db.Max.X-=d.dx
	db.Max.Y-=d.dy
	return db
}

//// visual Depiction of a Limited using adjected, trimmed, MaxX for bounds, implements Depictor
//func NewTrimmedLimitedDepiction(s Limited, dxX, dxY int) TrimmedLimitedDepiction {
//	return TrimmedLimitedDepiction{s, xspp(dxX,dxY,s.MaxX()) }
//}

// XXX find samller offset limits by dugging down
//func Bounds(l Limited) image.Rectangle{
//	switch c:=l.(type){
//	case Composite:
//		return bounds(c)
//	case UnlimitedComposite:
//		return bounds(Composite(c))
//	}
//	max:=int(d.Limited.MaxX()/d.xsperpixel)
//	return image.Rectangle{image.Pt(-max,-max),image.Pt(max,max)}
//}

//func bounds(c Composite) (r image.Rectangle){
//	for p:=range(c){
//		r.Add(image.Pt())
//	}
//}




// simple visual Depiction of a Unlimited, implements Depictor
type Depiction struct {
	Unlimited
	size       image.Rectangle
	in, out    color.Color
	xsperpixel x
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
func NewDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return NewCentredDepiction(s, dxX, dxY, in, out)
}

func (d Depiction) Bounds() image.Rectangle {
	return d.size
}

func (d Depiction) At(xp, yp int) color.Color {
	if d.at(x(xp)*d.xsperpixel, x(yp)*d.xsperpixel) == unitY {
		return d.in
	}
	return d.out
}

func xspp(dx,dy int,max x) x{
	if dx>dy {
		return 2 * max / x(dy)
	}		
	return 2 * max / x(dx)
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero centred, width fitted
func NewCentredDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(-dxX/2, -dxY/2, dxX/2, dxY/2), in, out, xspp(dxX,dxY,s.MaxX()) }
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero top left corner, width fitted
func NewFlowDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(0, 0, dxX, dxY), in, out, xspp(dxX,dxY,s.MaxX())}
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero top left corner, width fitted
func NewCentredBelowDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(-dxX/2, 0, dxX/2, dxY), in, out,xspp(dxX,dxY,s.MaxX())}
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero top left corner, width fitted
func NewCentredRightDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(0, -dxY/2, dxX, dxY/2), in, out,xspp(dxX,dxY,s.MaxX())}
}


type CachedImage struct{
	image.Image
	c color.Color
	x,y int
}

func (ci CachedImage) At(x, y int) color.Color {
	if ci.c==nil || x!=ci.x || y!=ci.y {
		ci.x,ci.y=x,y
		ci.c=ci.Image.At(x,y)
	}
	return ci.c
}


type OffsetImage struct {
	image.Image
	dx,dy int 
}

func (oi OffsetImage) At(x, y int) color.Color {
	return oi.Image.At(x-oi.dx,y-oi.dy)
}

// wrapper to add a dummy ColorModel to implement image.Image, used to avoid polution of depictor interface.
// (ColorModel is completely pointless/meaningless for non-editable images, that is for the interface, image.Image, its includied in! it is useful on draw.Image's so it seems basically to be in the wrong interface!!)
type Image struct {
	Depictor
}

type colormodel struct{}

func (colormodel) Convert(i color.Color) color.Color{
	return nil
}

var dummy colormodel

func (Image) ColorModel() color.Model { 
	return dummy
}

// Drawable simplifies draw.Draw for incremental composition of images
type Drawable struct {
	draw.Image
}

func (i Drawable) draw(isrc image.Image) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min, draw.Src)
}

func (i Drawable) drawAt(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, pt, draw.Src)
}

func (i Drawable) drawOffset(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min.Add(pt), draw.Src)
}

func (i Drawable) drawOver(isrc image.Image) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min, draw.Over)
}

func (i Drawable) drawOverAt(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, pt, draw.Over)
}

func (i Drawable) drawOverOffset(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min.Add(pt), draw.Over)
}
