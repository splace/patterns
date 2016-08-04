package patterns

import (
	"image"
	"image/color"
	"image/color/palette"
)

// simple visual Depiction of a Pattern, implements Depictor
type Depiction struct {
	Pattern
	size           image.Rectangle
	pixelsPerUnitX int
	in,out   color.Color
}

// makes a Depiction of a LimitedPattern, scaled to pxMaxx by pxMaxy pixels and sets the colours for above and below the value.
func NewDepiction(s LimitedPattern, pxMaxX, pxMaxY int, in,out color.Color) Depiction {
	return Depiction{s, image.Rect(0, -pxMaxY/2, pxMaxX, pxMaxY/2), int(int64(pxMaxX) * int64(unitX) / int64(s.MaxX())), in,out}
}

func (i Depiction) Bounds() image.Rectangle {
	return i.size
}

func (i Depiction) At(xp, yp int) color.Color {
	if i.at(x(xp)*unitX/x(i.pixelsPerUnitX),x(yp)*unitX/x(i.pixelsPerUnitX)) == unitY {
		return i.in
	}
	return i.out
}

// a Depictor is an image.Image without a colormodel, so is more general.
// embedded in one of the helper wrappers gets you an image.Image.
// (this and the wrappers might be moved, in a later version of this package, to their own package.)
type Depictor interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
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
