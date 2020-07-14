package patterns

import (
	"image"
	"image/color"
	//"image/color/palette"
	//"image/draw"
	//"image/jpeg" // register de/encoding
	"fmt"
	"image/png" // register de/encoding
	"os"
	"testing"
)

func TestImageSquare(t *testing.T) {
	file, err := os.Create("./test output/Square.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, Plan9PalettedImage{Depiction{Shrunk{Square(unitY), .10}, image.Rect(-40, -40, 40, 40), color.Opaque, color.Transparent,0.5*unitX}})
}

func TestImageBox(t *testing.T) {
	file, err := os.Create("./test output/Box.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	png.Encode(file, Plan9PalettedImage{NewDepiction(NewFrame(199, 1, Filling(unitY)), 400, 400, color.Opaque, color.Transparent)})
}

func TestImageLines(t *testing.T) {
	file, err := os.Create("./test output/Lines.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := Pen{Nib: Facetted{LineNib: LineNib{2 * unitX, unitY}, CurveDivision: 3}}

	png.Encode(file, Plan9PalettedImage{NewDepiction(p.Straight(-100*unitX, 100*unitX, 100*unitX, -100*unitX), 400, 400, color.Opaque, color.Transparent)})
}

func TestImageBitCoin(t *testing.T) {
	bitcoin := "M117.021,167.042c18.631,-9.483 30.288,-26.184 27.565,-54.007c-3.667,-38.023 -36.526,-50.773 -78.006,-54.404l-0.008,-52.741h-32.139l-0.009,51.354c-8.456,0 -17.076,0.166 -25.657,0.338L8.76,5.897l-32.11,-0.003l-0.006,52.728c-6.959,0.142 -13.793,0.277 -20.466,0.277v-0.156l-44.33,-0.018l0.006,34.282c0,0 23.734,-0.446 23.343,-0.013c13.013,0.009 17.262,7.559 18.484,14.076l0.01,60.083 v84.397c-0.573,4.09 -2.984,10.625 -12.083,10.637c0.414,0.364 -23.379,-0.004 -23.379,-0.004l-6.375,38.335h41.817c7.792,0.009 15.448,0.13 22.959,0.19l0.028,53.338l32.102,0.009l-0.009,-52.779c8.832,0.18 17.357,0.258 25.684,0.247l-0.009,52.532 h32.138l0.018,-53.249c54.022,-3.1 91.842,-16.697 96.544,-67.385C166.916,192.612 147.692,174.396 117.021,167.042zM9.535,95.321c18.126,0 75.132,-5.767 75.14,32.064c-0.008,36.269 -56.996,32.032 -75.14,32.032V95.321zM9.521,262.447l0.014,-70.672c21.778,-0.006 90.085,-6.261 90.094,35.32C99.638,266.971 31.313,262.431 9.521,262.447z"
	p := Path{}
	_, err := fmt.Sscan(bitcoin, &p)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("./test output/Bitcoin.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := NewFacettedBrush(10*unitX, Filling(unitY), 1)
//	b.Joiner=nil
//	b.Nib=LineNib{b.Nib.(Facetted).Width,b.Nib.(Facetted).In}
	png.Encode(file, Plan9PalettedImage{NewCentredBelowDepiction(p.Draw(b), 600, 600, color.Opaque, color.Transparent)})
}

func TestImageRings(t *testing.T) {
	f := "M-%[1]v,0C-%[1]v,-%[1]v %[1]v,-%[1]v %[1]v,0S-%[1]v,%[1]v -%[1]v,0z"

	radius := 10
	cpath := fmt.Sprintf(f, radius)
	radius = 7
	cpath += fmt.Sprintf(f, radius)

	p := Path{}
	_, err := fmt.Sscan(cpath, &p)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(p)
	file, err := os.Create("./test output/rings.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := NewFacettedBrush(2*unitX, Filling(unitY), 4)
	b1 := NewFacettedBrush(2*unitX, Filling(unitY), 3)
	c := Composite{UnlimitedShrunk{p.Draw(b), 0.2}, UnlimitedShrunk{p.Draw(b1), 0.5}}
	png.Encode(file, Plan9PalettedImage{NewDepiction(Limiter{c, 50 * unitX}, 1600, 1600, color.Opaque, color.Transparent)})
}

func TestImageRoundedRectangle(t *testing.T) {
	file, err := os.Create("./test output/RoundedBox.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := RoundedRectangle(20*unitX, 15*unitX, 5*unitX)
	fmt.Println(p)
	b := NewFacettedBrush(unitX, Filling(unitY), 2)
	c := UnlimitedShrunk{p.Draw(b), 0.3}
	png.Encode(file, Plan9PalettedImage{NewDepiction(Limiter{c, 50 * unitX}, 2500, 1600, color.Opaque, color.Transparent)})
}

func TestImageSmoothCorneredTrapezoid(t *testing.T) {
	file, err := os.Create("./test output/SmoothCorneredTrapezoid.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := Path{
		SmoothCorneredTrapezoid(25*unitX, 1*unitX, 20*unitX, 8*unitX),
		SmoothCorneredTrapezoid(30*unitX, 1.5*unitX, 25*unitX, 12*unitX),
	}
	b := NewFacettedBrush(unitX, Filling(unitY), 2)
	c := UnlimitedShrunk{p.Draw(b), 0.5}
	png.Encode(file, Plan9PalettedImage{NewDepiction(Limiter{c, 50 * unitX}, 2500, 1600, color.Opaque, color.Transparent)})
}

func TestImageBallCorneredRectangle(t *testing.T) {
	file, err := os.Create("./test output/BallCorneredRectangle.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := BallCorneredRectangle(25*unitX, 20*unitX, 8*unitX)
	fmt.Println(p)
	b := NewFacettedBrush(unitX, Filling(unitY), 3)
	c := UnlimitedShrunk{p.Draw(b), 0.5}
	png.Encode(file, Plan9PalettedImage{NewDepiction(Limiter{c, 50 * unitX}, 2500, 1600, color.Opaque, color.Transparent)})
}

func TestImageArcCorneredTrapezoid(t *testing.T) {
	file, err := os.Create("./test output/ArcCorneredTrapezoid.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := Path{
		ArcCorneredTrapezoid(25*unitX, 5*unitX, 20*unitX, 8*unitX, 4*unitX, 0, 45*unitX, true, true),
		ArcCorneredTrapezoid(30*unitX, 5*unitX, 24*unitX, 8*unitX, 4*unitX, 0, 45*unitX, true, false),
	}
	b := NewFacettedBrush(2*unitX, Filling(unitY), 3)
	//c:=UnlimitedShrunk{p.Draw(b),0.5}
	png.Encode(file, OpaqueTransparentPalettedImage{NewDepiction(Limiter{p.Draw(b), 30 * unitX}, 5000, 3200, color.Opaque, color.Transparent)})
}
