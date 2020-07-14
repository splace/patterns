package pattern

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
)

// a Depictor is an image.Image, missing a colormodel, it is thus more general.
// embedded in one of the helper wrappers gets you an image.Image.
type Depictor interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
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

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero centred, width fitted
func NewCentredDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(-dxX/2, -dxY/2, dxX/2, dxY/2), in, out, unitX / x(int(int64(dxX)*int64(unitX)/int64(s.MaxX())/4+1))}
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero top left corner, width fitted
func NewFlowDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(0, 0, dxX, dxY), in, out, unitX / x(int(int64(dxX)*int64(unitX)/int64(s.MaxX())/4+1))}
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero top left corner, width fitted
func NewCentredBelowDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(-dxX/2, 0, dxX/2, dxY), in, out, unitX / x(int(int64(dxX)*int64(unitX)/int64(s.MaxX())/4+1))}
}

// makes a Depiction of a Limited, scaled to dxX by dxY pixels and sets the colours for above and below the value.
// zero top left corner, width fitted
func NewCentredRightDepiction(s Limited, dxX, dxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(0, -dxY/2, dxX, dxY/2), in, out, unitX / x(int(int64(dxX)*int64(unitX)/int64(s.MaxX())/4+1))}
}

func (i Depiction) Bounds() image.Rectangle {
	return i.size
}

func (i Depiction) At(xp, yp int) color.Color {
	if i.at(x(xp)*i.xsperpixel, x(yp)*i.xsperpixel) == unitY {
		return i.in
	}
	return i.out
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
type OpaqueTransparentPalettedImage struct {
	Depictor
}

func (i OpaqueTransparentPalettedImage) ColorModel() color.Model {
	return color.Palette([]color.Color{color.Opaque, color.Transparent})
}

// composable simplifies draw.Draw for incremental composition of images
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
