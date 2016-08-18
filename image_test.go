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

func TestImageLines(t *testing.T) {
	file, err := os.Create("./test output/Lines.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := Brush{Width: 2*unitX, In: unitY, Relative: true}
	png.Encode(file, Plan9PalettedImage{NewDepiction(b.Line(-100*unitX, 100*unitX, 100*unitX, -100*unitX), 400, 400, color.Opaque, color.Transparent)})
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
}/*  Hal3 Thu 18 Aug 01:29:45 BST 2016 go version go1.6.2 linux/amd64
=== RUN   TestImageSquare
--- PASS: TestImageSquare (0.02s)
=== RUN   TestImageBox
--- PASS: TestImageBox (0.08s)
=== RUN   TestImageLines
--- PASS: TestImageLines (0.10s)
=== RUN   ExampleBrushLine
--- PASS: ExampleBrushLine (0.00s)
=== RUN   ExampleComposite
--- PASS: ExampleComposite (0.00s)
=== RUN   ExampleShifted
--- PASS: ExampleShifted (0.00s)
=== RUN   ExampleTranslated
--- PASS: ExampleTranslated (0.00s)
=== RUN   ExampleZoomed
--- PASS: ExampleZoomed (0.00s)
=== RUN   ExampleScaled
--- PASS: ExampleScaled (0.00s)
=== RUN   ExampleRotated
--- PASS: ExampleRotated (0.00s)
=== RUN   ExampleInverted
--- PASS: ExampleInverted (0.00s)
=== RUN   ExampleConstant
--- PASS: ExampleConstant (0.00s)
=== RUN   ExampleDisc
--- PASS: ExampleDisc (0.00s)
=== RUN   ExampleSquare
--- PASS: ExampleSquare (0.00s)
PASS
ok  	_/home/simon/Dropbox/github/working/patterns	0.219s
Thu 18 Aug 01:29:46 BST 2016 */

