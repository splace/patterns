package pattern

import "image"
import "image/color"

// in order to render (parallel process image in sections) and use std.lib. image file encoding, need to draw to a memory (buffer/cache) image 

// ImageY is an in-memory image.Image (+ draw.Image) whose At method returns Y values.
type ImageY struct {
	Pix []y
	Stride int
	Rect image.Rectangle
}

func (p *ImageY) At(x, y int) color.Color {
	return p.YAt(x, y)
}

func (p *ImageY) YAt(x, y int) y {
	if !(image.Point{x, y}.In(p.Rect)) {
		return zeroY
	}
	return p.Pix[p.PixOffset(x, y)]
}

func (p *ImageY) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)
}

func (p *ImageY) Bounds() image.Rectangle { return p.Rect }

func (p *ImageY) ColorModel() color.Model { return YModel }

func (p *ImageY) Set(px, py int, c color.Color) {
	if !(image.Point{px, py}.In(p.Rect)) {
		return
	}
	p.Pix[p.PixOffset(px, py)] = YModel.Convert(c).(y)
	return
}

func (p *ImageY) SetY(x, y int, c y ) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	p.Pix[p.PixOffset(x, y)] = 	c
	return
}

func (p *ImageY) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	if r.Empty() {
		return &ImageY{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &ImageY{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// NewRGBA returns a new RGBA image with the given bounds.
func NewImageY(r image.Rectangle) *ImageY {
	w, h := r.Dx(), r.Dy()
	buf := make([]y, w*h)
	return &ImageY{buf, w, r}
}
