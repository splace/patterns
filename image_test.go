package patterns

import (
	"image"
	"image/color"
	//"image/color/palette"
	"image/draw"
	//"image/jpeg" // register de/encoding
	"image/png" // register de/encoding
	"os"
	"testing"
	//"fmt"
)

func TestImageSquare(t *testing.T) {
	file, err := os.Create("./test output/Square.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, Plan9PalettedImage{Depiction{Shrunk{Square{Filling{unitY}}, .10}, image.Rect(-40, -40, 40, 40), 2, color.Opaque, color.Transparent}})
}

func TestImageBox(t *testing.T) {
	file, err := os.Create("./test output/Box.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, Plan9PalettedImage{NewDepiction(NewBox(199, 1, Filling{unitY}), 400, 400, color.Opaque, color.Transparent)})
}

//ds := NewDepiction(NewBox(38,2,Filling{unitY}), 400, 400, color.Opaque, color.Transparent)

//m := &composable{image.NewPaletted(image.Rect(0, -150, 800, 150), palette.WebSafe)}
//m.draw(WebSafePalettedImage{ds})
//jpeg.Encode(file, m, nil)

// composable is a draw.Image that comes with helper functions to simplify Draw function.
type composable struct {
	draw.Image
}

func (i *composable) draw(isrc image.Image) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min, draw.Src)
}

func (i *composable) drawAt(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, pt, draw.Src)
}

func (i *composable) drawOffset(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min.Add(pt), draw.Src)
}

func (i *composable) drawOver(isrc image.Image) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min, draw.Over)
}

func (i *composable) drawOverAt(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, pt, draw.Over)
}

func (i *composable) drawOverOffset(isrc image.Image, pt image.Point) {
	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min.Add(pt), draw.Over)
}
