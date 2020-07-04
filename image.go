package patterns

import (
	"image"
	"image/color"
	"image/color/palette"
)

// a Depictor is like an image.Image but without a colormodel, so is more general.
// embedded in one of the helper wrappers gets you an image.Image.
type Depictor interface {
	Bounds() image.Rectangle
	At(x, y int) color.Color
}

// simple visual Depiction of a Pattern, implements Depictor
type Depiction struct {
	Pattern
	size           image.Rectangle
	in, out        color.Color
	xsperpixel x
}

// makes a Depiction of a LimitedPattern, scaled to pxMaxx by pxMaxy pixels and sets the colours for above and below the value.
func NewDepiction(s LimitedPattern, pxMaxX, pxMaxY int, in, out color.Color) Depiction {
	return Depiction{s, image.Rect(-pxMaxX/2, -pxMaxY/2, pxMaxX/2, pxMaxY/2), in, out, unitX/x(int(int64(pxMaxX)*int64(unitX)/int64(s.MaxX())/4 + 1))}
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
