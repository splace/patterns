package pattern

import (
	"image"
	"image/color"
	//"image/color/palette"
	"image/draw"
	//"image/jpeg" // register de/encoding
	"fmt"
	"image/png" // register de/encoding
	"os"
	"testing"
)

func ExampleRenderSquare() {
	file, err := os.Create("./test output/Square(rendered).png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	i:=NewImageY(image.Rect(-40, -40, 40, 40))
	s:=Image{Depiction{Shrunk{Square(unitY), .10}, i.Bounds(), color.Opaque, color.Transparent, 0.5 * unitX}}
	
	draw.Draw(i,s.Bounds(), s, s.Bounds().Min, draw.Over)
	
	fmt.Println(png.Encode(file,i ))
	
	// Output:
	// 

}

func TestRenderBitCoin(t *testing.T) {
	bitcoin := "M117.021,167.042c18.631,-9.483 30.288,-26.184 27.565,-54.007c-3.667,-38.023 -36.526,-50.773 -78.006,-54.404l-0.008,-52.741h-32.139l-0.009,51.354c-8.456,0 -17.076,0.166 -25.657,0.338L8.76,5.897l-32.11,-0.003l-0.006,52.728c-6.959,0.142 -13.793,0.277 -20.466,0.277v-0.156l-44.33,-0.018l0.006,34.282c0,0 23.734,-0.446 23.343,-0.013c13.013,0.009 17.262,7.559 18.484,14.076l0.01,60.083 v84.397c-0.573,4.09 -2.984,10.625 -12.083,10.637c0.414,0.364 -23.379,-0.004 -23.379,-0.004l-6.375,38.335h41.817c7.792,0.009 15.448,0.13 22.959,0.19l0.028,53.338l32.102,0.009l-0.009,-52.779c8.832,0.18 17.357,0.258 25.684,0.247l-0.009,52.532 h32.138l0.018,-53.249c54.022,-3.1 91.842,-16.697 96.544,-67.385C166.916,192.612 147.692,174.396 117.021,167.042zM9.535,95.321c18.126,0 75.132,-5.767 75.14,32.064c-0.008,36.269 -56.996,32.032 -75.14,32.032V95.321zM9.521,262.447l0.014,-70.672c21.778,-0.006 90.085,-6.261 90.094,35.32C99.638,266.971 31.313,262.431 9.521,262.447z"
	p := Path{}
	_, err := fmt.Sscan(bitcoin, &p)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Create("./test output/Bitcoin(rendered).png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	b := NewFacettedBrush(10*unitX, Filling(unitY), 1)
	//	b.Joiner=nil
	//	b.Nib=LineNib{b.Nib.(Facetted).Width,b.Nib.(Facetted).In}
	
	
	pat:=NewLimitedDepiction(p.Draw(b), 600, 600)
//	i:=NewImageY(pat.Bounds(),color.Opaque,color.Transparent)
	s:=Image{pat}
	
//	draw.Draw(i,s.Bounds(), s, s.Bounds().Min, draw.Over)
//	draw.Draw(i,s.Bounds(), s, image.ZP, draw.Over)
//	draw.Draw(i,s.Bounds(), s, s.Bounds().Max, draw.Over)
	
	png.Encode(file,s)

}


