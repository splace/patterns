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
	"fmt"
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

	png.Encode(file, Plan9PalettedImage{NewDepiction(NewFrame(199, 1, Filling{unitY}), 400, 400, color.Opaque, color.Transparent)})
}

func TestImageLines(t *testing.T) {
	file, err := os.Create("./test output/Lines.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := Brush{LineBrush:LineBrush{2*unitX, unitY}, Relative: true}
	png.Encode(file, Plan9PalettedImage{NewDepiction(b.Line(-100*unitX, 100*unitX, 100*unitX, -100*unitX), 400, 400, color.Opaque, color.Transparent)})
}

func TestImageBitCoin(t *testing.T) {
	p:=Path{}
	_,err:=fmt.Sscan("M217.021,167.042 c18.631,-9.483 30.288,-26.184 27.565,-54.007 c-3.667,-38.023 -36.526,-50.773 -78.006,-54.404 l-0.008,-52.741 h-32.139 l-0.009,51.354 c-8.456,0 -17.076,0.166 -25.657,0.338 L108.76,5.897 l-32.11,-0.003 l-0.006,52.728 c-6.959,0.142 -13.793,0.277 -20.466,0.277 v-0.156 l-44.33,-0.018 l0.006,34.282 c0,0 23.734,-0.446 23.343,-0.013 c13.013,0.009 17.262,7.559 18.484,14.076 l0.01,60.083 v84.397 c-0.573,4.09 -2.984,10.625 -12.083,10.637 c0.414,0.364 -23.379,-0.004 -23.379,-0.004 l-6.375,38.335 h41.817 c7.792,0.009 15.448,0.13 22.959,0.19 l0.028,53.338 l32.102,0.009 l-0.009,-52.779 c8.832,0.18 17.357,0.258 25.684,0.247 l-0.009,52.532 h32.138 l0.018,-53.249 c54.022,-3.1 91.842,-16.697 96.544,-67.385 C266.916,192.612 247.692,174.396 217.021,167.042 z  M109.535,95.321 c18.126,0 75.132,-5.767 75.14,32.064 c-0.008,36.269 -56.996,32.032 -75.14,32.032 V95.321 z  M109.521,262.447 l0.014,-70.672 c21.778,-0.006 90.085,-6.261 90.094,35.32 C199.638,266.971 131.313,262.431 109.521,262.447 z",&p)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("./test output/Bitcoin.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := Brush{LineBrush:LineBrush{20*unitX, unitY}, Relative: true}
	png.Encode(file, Plan9PalettedImage{NewDepiction(Limiter{UnlimitedShrunk{p.Draw(&b),6},50*unitX}, 400, 400, color.Opaque, color.Transparent)})
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
/*  Hal3 Thu 18 Aug 01:32:01 BST 2016 go version go1.6.2 linux/amd64
=== RUN   TestImageSquare
--- PASS: TestImageSquare (0.00s)
=== RUN   TestImageBox
--- PASS: TestImageBox (0.08s)
=== RUN   TestImageLines
--- PASS: TestImageLines (0.13s)
PASS
ok  	_/home/simon/Dropbox/github/working/patterns	0.231s
Thu 18 Aug 01:32:03 BST 2016 */

