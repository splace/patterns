package pattern

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
)

// a Image is an image.Image, missing a colormodel, so more general.
// embedded in one of the helper wrappers gets you an image.Image.
type Depictor interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
}



// simple visual Depiction of a Unlimited, implements Depictor
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

// RGBA depiction wrapper
type RGBAImage struct {
	Depictor
}

func (i RGBAImage) ColorModel() color.Model { return color.RGBAModel }

// gray depiction wrapper.
type GrayImage struct {
	Depictor
}

func (i GrayImage) ColorModel() color.Model { return color.GrayModel }

// plan9 paletted, depiction wrapper.
type Plan9PalettedImage struct {
	Depictor
}

func (i Plan9PalettedImage) ColorModel() color.Model { return color.Palette(palette.Plan9) }

// WebSafe paletted, depiction wrapper.
type WebSafePalettedImage struct {
	Depictor
}

func (i WebSafePalettedImage) ColorModel() color.Model { return color.Palette(palette.WebSafe) }

// black/white paletted, depiction wrapper.
type BlackAndWhitePalettedImage struct {
	Depictor
}

func (i BlackAndWhitePalettedImage) ColorModel() color.Model {
	return color.Palette([]color.Color{color.Black, color.White})
}


// black/white paletted, depiction wrapper.
type PalettedImage struct {
	Depictor
}

func (i PalettedImage) ColorModel() color.Model {
	return color.Palette([]color.Color{zeroY, unitY})
}

// black/white paletted, depiction wrapper.
type OpaqueTransparentPalettedImage struct {
	Depictor
}

func (i OpaqueTransparentPalettedImage) ColorModel() color.Model {
	return color.Palette([]color.Color{color.Opaque, color.Transparent})
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
