package patterns

import (
	"image"
	"image/color"
	//"image/color/palette"
	//"image/draw"
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
	p := Pen{Nib:Facetted{2*unitX, unitY,3}, Relative: true}

	png.Encode(file, Plan9PalettedImage{NewDepiction(p.Line(-100*unitX, 100*unitX, 100*unitX, -100*unitX), 400, 400, color.Opaque, color.Transparent)})
}

func TestImageBitCoin(t *testing.T) {
	p:=Path{}
	_,err:=fmt.Sscan("M117.021,167.042c18.631,-9.483 30.288,-26.184 27.565,-54.007c-3.667,-38.023 -36.526,-50.773 -78.006,-54.404l-0.008,-52.741h-32.139l-0.009,51.354c-8.456,0 -17.076,0.166 -25.657,0.338L8.76,5.897l-32.11,-0.003l-0.006,52.728c-6.959,0.142 -13.793,0.277 -20.466,0.277v-0.156l-44.33,-0.018l0.006,34.282c0,0 23.734,-0.446 23.343,-0.013c13.013,0.009 17.262,7.559 18.484,14.076l0.01,60.083 v84.397c-0.573,4.09 -2.984,10.625 -12.083,10.637c0.414,0.364 -23.379,-0.004 -23.379,-0.004l-6.375,38.335h41.817c7.792,0.009 15.448,0.13 22.959,0.19l0.028,53.338l32.102,0.009l-0.009,-52.779c8.832,0.18 17.357,0.258 25.684,0.247l-0.009,52.532 h32.138l0.018,-53.249c54.022,-3.1 91.842,-16.697 96.544,-67.385C166.916,192.612 147.692,174.396 117.021,167.042zM9.535,95.321c18.126,0 75.132,-5.767 75.14,32.064c-0.008,36.269 -56.996,32.032 -75.14,32.032V95.321zM9.521,262.447l0.014,-70.672c21.778,-0.006 90.085,-6.261 90.094,35.32C99.638,266.971 31.313,262.431 9.521,262.447z",&p)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("./test output/Bitcoin.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := Brush{Pen:Pen{Nib:Facetted{20*unitX, unitY,3}, Relative: true}}
	png.Encode(file, Plan9PalettedImage{NewDepiction(Limiter{UnlimitedShrunk{p.Draw(&b),6},50*unitX}, 400, 400, color.Opaque, color.Transparent)})
}

func TestImageRings(t *testing.T) {
	f:="M-%[1]v,0C-%[1]v,-%[1]v %[1]v,-%[1]v %[1]v,0S-%[1]v,%[1]v -%[1]v,0z"
	
	radius:=10
	cpath:=fmt.Sprintf(f,radius)
	radius=7
	cpath+=fmt.Sprintf(f,radius)

	p:=Path{}
	_,err:=fmt.Sscan(cpath,&p)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("./test output/rings.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := Brush{Pen:Pen{Nib:Facetted{1*unitX, unitY,4}, Relative: true}}
	b1 := Brush{Pen:Pen{Nib:Facetted{2*unitX, unitY,3}, Relative: true}}
	c:=Composite{UnlimitedShrunk{p.Draw(&b),0.2},UnlimitedShrunk{p.Draw(&b1),0.5}}
	png.Encode(file, Plan9PalettedImage{NewDepiction(Limiter{c,50*unitX}, 1600, 1600, color.Opaque, color.Transparent)})
}





////ds := NewDepiction(NewBox(38,2,Filling{unitY}), 400, 400, color.Opaque, color.Transparent)

////m := &composable{image.NewPaletted(image.Rect(0, -150, 800, 150), palette.WebSafe)}
////m.draw(WebSafePalettedImage{ds})
////jpeg.Encode(file, m, nil)

//// composable is a draw.Image that comes with helper functions to simplify Draw function.
//type composable struct {
//	draw.Image
//}

//func (i *composable) draw(isrc image.Image) {
//	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min, draw.Src)
//}

//func (i *composable) drawAt(isrc image.Image, pt image.Point) {
//	draw.Draw(i, i.Bounds(), isrc, pt, draw.Src)
//}

//func (i *composable) drawOffset(isrc image.Image, pt image.Point) {
//	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min.Add(pt), draw.Src)
//}

//func (i *composable) drawOver(isrc image.Image) {
//	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min, draw.Over)
//}

//func (i *composable) drawOverAt(isrc image.Image, pt image.Point) {
//	draw.Draw(i, i.Bounds(), isrc, pt, draw.Over)
//}

//func (i *composable) drawOverOffset(isrc image.Image, pt image.Point) {
//	draw.Draw(i, i.Bounds(), isrc, isrc.Bounds().Min.Add(pt), draw.Over)
//}
