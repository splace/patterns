package pattern

import "fmt"
import "testing"

func ExampleSVGpath() {
	p := Path{MoveTo{0, 0}, LineTo{10, 10}}
	fmt.Printf("%#v\n", p)
	// Output:
	// patterns.Path{patterns.MoveTo{0, 0}, patterns.LineTo{10, 10}}
}

func ExampleSVGDrawPath() {
	p := Path{MoveTo{-5 * unitX, -5 * unitX}, LineTo{5 * unitX, 5 * unitX}}
	b := NewBrush(FacettedNib{LineNib: LineNib{2 * unitX, unitY}, CurveDivision: 2})
	PrintGraph(p.Draw(b), -10*unitX, 10*unitX, -10*unitX, 10*unitX, unitX)
	/* Output:
	Graph
	      -10	---------------------
	       -9	---------------------
	       -8	---------------------
	       -7	---------------------
	       -6	---------------------
	       -5	------X--------------
	       -4	-----XXX-------------
	       -3	------XXX------------
	       -2	-------XXX-----------
	       -1	--------XXX----------
	        0	---------XXX---------
	        1	----------XXX--------
	        2	-----------XXX-------
	        3	------------XXX------
	        4	-------------XXX-----
	        5	--------------XX-----
	        6	---------------------
	        7	---------------------
	        8	---------------------
	        9	---------------------
	       10	---------------------
	*/
}

func ExampleSVGpathScan() {
	p := Path{}
	_, err := fmt.Sscan("m 0,0 5,5 h -10 z", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", p)
	// Output:
	// patterns.Path{patterns.MoveToRelative{0, 0}, patterns.LineToRelative{5000, 5000}, patterns.HorizontalLineToRelative{-10000}, patterns.CloseRelative{}}
}

func ExampleSVGcompactPathScan() {
	p := Path{}
	_, err := fmt.Sscan("m0-1 5 5h-10z", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", p)
	// Output:
	// patterns.Path{patterns.MoveToRelative{0, -1000}, patterns.LineToRelative{5000, 5000}, patterns.HorizontalLineToRelative{-10000}, patterns.CloseRelative{}}
}

func ExampleSVGpathPrint() {
	p := Path{}
	_, err := fmt.Sscan("m 0 -1 5 0.5 0 0 h -10 l 0.5,2z", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	// Output:
	// m0,-1 5,0.5 0,0
	// h-10
	// l0.5,2
	// z
}

func ExampleSVGcompactPathPrint() {
	p := Path{}
	_, err := fmt.Sscan("m 0 -1 -5 0.5 0 0 h -10 l 0.5,2z", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", MaxCompactStringer(p))
	// Output:
	// m0-1-5,.5,0,0h-10l.5,2z
}

func TestSVGcompactPathTextRoundtrip(t *testing.T) {
	p := Path{}
	_, err := fmt.Sscan("m 0 -1 5 0.5 0 0 h -10 l 0.5,2z", &p)
	if err != nil {
		fmt.Println(err)
	}
	cpt := fmt.Sprint(MaxCompactStringer(p))
	p2 := Path{}
	_, err = fmt.Sscan(cpt, &p2)
	if err != nil {
		fmt.Println(err)
	}
	if cpt != fmt.Sprint(MaxCompactStringer(p2)) {
		t.Errorf("%q != %q (%v)", cpt, fmt.Sprint(MaxCompactStringer(p2)), p2)
	}
}

// XXX no right
func ExampleSVGcompactPathHard() {
	p := Path{}
	_, err := fmt.Sscan("M66-8c19-9 30-26 28-54-4-38-37-51-78-54l0-53h-32l0 51c-8 0-17 0-26 0L-42-169l-32 0 0 53c-7 0-14 0-20 0v0l-44 0 0 34c0 0 24 0 23 0 13 0 17 8 18 14l0 60v84c-1 4-3 11-12 11 0 0-23 0-23 0l-6 38h42c8 0 15 0 23 0l0 53 32 0 0-53c9 0 17 0 26 0l0 53h32l0-53c54-3 92-17 97-67C116 18 97-1 66-8zM-41-80c18 0 75-6 75 32 0 36-57 32-75 32V-80zM-41 87l0-71c22 0 90-6 90 35C49 92-20 87-41 87z", &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	// Output:

}

func ExampleSVGpathScanMissingCommands() {
	p := Path{}
	_, err := fmt.Sscan("m -5 -5 l 10 10 h -10", &p)
	if err != nil {
		fmt.Println(err)
	}
	b := NewBrush(FacettedNib{LineNib: LineNib{unitX, unitY}, CurveDivision: 2})
	Output(Limiter{p.Draw(b), 10 * unitX}, unitX)
	/* Output:
	Graph
	      -11	-----------------------
	      -10	-----------------------
	       -9	-----------------------
	       -8	-----------------------
	       -7	-----------------------
	       -6	-----------------------
	       -5	-----------------------
	       -4	-------X---------------
	       -3	--------X--------------
	       -2	---------X-------------
	       -1	----------X------------
	        0	-----------X-----------
	        1	------------X----------
	        2	-------------X---------
	        3	--------------X--------
	        4	---------------X-------
	        5	------XXXXXXXXXXX------
	        6	-----------------------
	        7	-----------------------
	        8	-----------------------
	        9	-----------------------
	       10	-----------------------
	       11	-----------------------
	*/
}

func ExampleSVGpathScanCubicOverlap() {
	p := Path{}
	_, err := fmt.Sscan("M20,15 C-30,-30 -30,30 20,-15", &p)
	if err != nil {
		fmt.Println(err)
	}
	b := NewBrush(FacettedNib{LineNib: LineNib{unitX, unitY}, CurveDivision: 2})
	Output(Limiter{p.Draw(b), 20 * unitX}, unitX)
	/* Output:
	Graph
	      -21	-------------------------------------------
	      -20	-------------------------------------------
	      -19	-------------------------------------------
	      -18	-------------------------------------------
	      -17	-------------------------------------------
	      -16	-------------------------------------------
	      -15	-----------------------------------------X-
	      -14	---------------------------------------XX--
	      -13	-------------------------------------XX----
	      -12	------------------------------------XX-----
	      -11	----------------------------------XX-------
	      -10	--------------------------------XX---------
	       -9	-------------------------------XX----------
	       -8	-----------------------------XX------------
	       -7	----------------------------X--------------
	       -6	--------------------------XX---------------
	       -5	------------------------XX-----------------
	       -4	-----------------------XX------------------
	       -3	---------------------XX--------------------
	       -2	----------XXXXX----XX----------------------
	       -1	------XXXX-----XX-XX-----------------------
	        0	----XX----------XX-------------------------
	        1	------XXXX-----X--XX-----------------------
	        2	----------XXXXX-----X----------------------
	        3	---------------------XX--------------------
	        4	-----------------------XX------------------
	        5	------------------------XX-----------------
	        6	--------------------------XX---------------
	        7	----------------------------XX-------------
	        8	-----------------------------XX------------
	        9	-------------------------------XX----------
	       10	--------------------------------XX---------
	       11	----------------------------------XX-------
	       12	------------------------------------XX-----
	       13	-------------------------------------XX----
	       14	---------------------------------------XX--
	       15	-----------------------------------------X-
	       16	-------------------------------------------
	       17	-------------------------------------------
	       18	-------------------------------------------
	       19	-------------------------------------------
	       20	-------------------------------------------
	       21	-------------------------------------------
	*/
}

func ExampleSVGpathScanSmoothQuadratic() {

	p := Path{}
	_, err := fmt.Sscan("M-50,30 Q-40,5 -20,30 T40,30", &p)
	if err != nil {
		fmt.Println(err)
	}
	b := NewBrush(FacettedNib{LineNib: LineNib{unitX, unitY}, CurveDivision: 2})
	Output(Limiter{p.Draw(b), 50 * unitX}, unitX)
	/* Output:
	Graph
	       -8	-----------------
	       -7	-----------------
	       -6	--XXXXXXXXXXXX---
	       -5	--XXXXXXXXXXXX---
	       -4	--XX--------XX---
	       -3	--XX--------XX---
	       -2	--XX--------XX---
	       -1	--XX--------XX---
	        0	--XX--------XX---
	        1	--XX--------XX---
	        2	--XX--------XX---
	        3	--XX--------XX---
	        4	--XXXXXXXXXXXX---
	        5	--XXXXXXXXXXXX---
	        6	-----------------
	        7	-----------------
	        8	-----------------
	*/
}

// INFO oplate circles
func ExampleSVGpathScanSmoothCubic() {

	bezierEllipseApprox := "M-%[1]v,0C-%[1]v,-%[1]v %[1]v,-%[1]v %[1]v,0S-%[1]v,%[1]v -%[1]v,0z"

	cpath := fmt.Sprintf(bezierEllipseApprox, 15)
	cpath += fmt.Sprintf(bezierEllipseApprox, 10)
	cpath += fmt.Sprintf(bezierEllipseApprox, 5)

	p := Path{}
	_, err := fmt.Sscan(cpath, &p)
	if err != nil {
		fmt.Println(err)
	}
	b := NewBrush(FacettedNib{LineNib: LineNib{unitX, unitY}, CurveDivision: 3})

	Output(Reduced{Limiter{p.Draw(b), 16 * unitX}, 1, 0.75}, unitX)
	/* Output:
	Graph
	       -22	---------------------------------------------
	      -21	---------------------------------------------
	      -20	---------------------------------------------
	      -19	---------------------------------------------
	      -18	---------------------------------------------
	      -17	---------------------------------------------
	      -16	---------------------------------------------
	      -15	---------------------XXXX--------------------
	      -14	-----------------XXXXXXXXXXXX----------------
	      -13	---------------XX-----------XXX--------------
	      -12	-------------XXX--------------XXX------------
	      -11	------------XX------------------XX-----------
	      -10	-----------XX--------XXXX--------X-----------
	       -9	----------XX------XXXXXXXXXX------X----------
	       -8	----------X-----XXX--------XXX-----X---------
	       -7	---------X-----XX------------XX----XX--------
	       -6	---------X----XX--------------X-----X--------
	       -5	--------X-----X------XXXX------X----X--------
	       -4	--------X----------XXXXXXX-----X----X--------
	       -3	--------X----X----XX------X-----X----X-------
	       -2	--------X----X----X--------X----X----X-------
	       -1	--------X----X----X--------X----X----X-------
	        0	-------X----X----X---------X----X----X-------
	        0	-------X----X----X---------X----X----X-------
	        1	--------X----X----X--------X----X----X-------
	        2	--------X----X----X-------XX----X----X-------
	        3	--------X----X-----XX----XX-----X----X-------
	        4	--------X----XX-----XXXXXX-----X----X--------
	        5	--------X-----X---------------XX----X--------
	        6	---------X-----X--------------X-----X--------
	        7	---------XX----XX-----------XX-----X---------
	        8	----------X------XX-------XXX-----XX---------
	        9	-----------X------XXXXXXXXX-------X----------
	       10	-----------XX--------------------X-----------
	       11	------------XX-----------------XX------------
	       12	--------------XX-------------XXX-------------
	       13	---------------XXXX--------XXX---------------
	       14	-----------------XXXXXXXXXXXX----------------
	       15	---------------------------------------------
	       16	---------------------------------------------
	       17	---------------------------------------------
	       18	---------------------------------------------
	       19	---------------------------------------------
	       20	---------------------------------------------
	       21	---------------------------------------------
	*/
}

func ExampleSVGpathScanMulti() {
	bitcoin := `M66-8c19-9 30-26 28-54-4-38-37-51-78-54l0-53h-32l0 51c-8 0-17 0-26 0L-42-169l-32 0 0 53c-7 0-14 0-20 0v0l-44 0 0 34c0 0 24 0 23 0 13 0 17 8 18 14l0 60v84c-1 4-3 11-12 11 0 0-23 0-23 0l-6 38h42c8 0 15 0 23 0l0 53 32 0 0-53c9 0 17 0 26 0l0 53h32l0-53c54-3 92-17 97-67C116 18 97-1 66-8zM-41-80c18 0 75-6 75 32 0 36-57 32-75 32V-80zM-41 87l0-71c22 0 90-6 90 35C49 92-20 87-41 87z`
	/*	bitcoin:=
		`M 117 167
		c 19 -9 30 -26 28 -54
		c -4 -38 -37 -51 -78 -54
		l 0 -53
		h -32
		l 0 51
		c -8 0 -17 0 -26 0
		L 9 6
		l -32 0
		l 0 53
		c -7 0 -14 0 -20 0
		v 0
		l -44 0
		l 0 34
		c 0 0 24 0 23 0
		c 13 0 17 8 18 14
		l 0 60
		v 84
		c -1 4 -3 11 -12 11
		c 0 0 -23 0 -23 0
		l -6 38
		h 42
		c 8 0 15 0 23 0
		l 0 53
		l 32 0
		l 0 -53
		c 9 0 17 0 26 0
		l 0 53
		h 32
		l 0 -53
		c 54 -3 92 -17 97 -67
		C 167 193 148 174 117 167
		z
		M 10 95
		c 18 0 75 -6 75 32
		c 0 36 -57 32 -75 32
		V 95
		z
		M 10 262
		l 0 -71
		c 22 0 90 -6 90 35
		C 100 267 31 262 10 262
		z
		`
	*/
	p := Path{}
	_, err := fmt.Sscan(bitcoin, &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", p)
	b := NewBrush(FacettedNib{LineNib: LineNib{16 * unitX, unitY}, CurveDivision: 2})
	Output(Limiter{UnlimitedReduced{p.Draw(b), 2.3, 4.4}, 60 * unitX}, unitX)
	/* Output:
	Graph
	       -8	-----------------
	       -7	-----------------
	       -6	--XXXXXXXXXXXX---
	       -5	--XXXXXXXXXXXX---
	       -4	--XX--------XX---
	       -3	--XX--------XX---
	       -2	--XX--------XX---
	       -1	--XX--------XX---
	        0	--XX--------XX---
	        1	--XX--------XX---
	        2	--XX--------XX---
	        3	--XX--------XX---
	        4	--XXXXXXXXXXXX---
	        5	--XXXXXXXXXXXX---
	        6	-----------------
	        7	-----------------
	        8	-----------------
	*/
}

func ExampleSVGPathMarker() {
	p := Path{
		MoveTo{-5 * unitX, -5 * unitX},
		LineTo{5 * unitX, 5 * unitX},
		LineTo{5 * unitX, 5 * unitX},
	}
	b := NewBrush(FacettedNib{LineNib: LineNib{2 * unitX, unitY}, CurveDivision: 2})
	b.StartMarker = Shrunk{Square(unitY), .33}
	b.EndMarker = Shrunk{Disc(unitY), .33}
	PrintGraph(p.Draw(b), -10*unitX, 10*unitX, -10*unitX, 10*unitX, unitX)
	/* Output:
	Graph
	      -10	---------------------
	       -9	---------------------
	       -8	---------------------
	       -7	---------------------
	       -6	---------------------
	       -5	------X--------------
	       -4	-----XXX-------------
	       -3	------XXX------------
	       -2	-------XXX-----------
	       -1	--------XXX----------
	        0	---------XXX---------
	        1	----------XXX--------
	        2	-----------XXX-------
	        3	------------XXX------
	        4	-------------XXX-----
	        5	--------------XX-----
	        6	---------------------
	        7	---------------------
	        8	---------------------
	        9	---------------------
	       10	---------------------
	*/
}

func ExampleSVGPathMarkerGaps() {
	p := Path{
		MoveTo{0 * unitX, -5 * unitX},
		MoveTo{-15 * unitX, 5 * unitX},
		LineTo{-5 * unitX, 5 * unitX},
		LineTo{-5 * unitX, -5 * unitX},
		MoveTo{5 * unitX, -5 * unitX},
		MoveTo{5 * unitX, 5 * unitX},
		LineTo{5 * unitX, -5 * unitX},
		LineTo{15 * unitX, 5 * unitX},
		MoveTo{0 * unitX, -5 * unitX},
		MoveTo{0 * unitX, -5 * unitX},
	}
	b := NewBrush(FacettedNib{LineNib: LineNib{2 * unitX, unitY}, CurveDivision: 2})
	b.StartMarker = Shrunk{Square(unitY), .33}
	b.EndMarker = Shrunk{Disc(unitY), .33}
	PrintGraph(p.Draw(b), -20*unitX, 20*unitX, -10*unitX, 10*unitX, unitX)
	/* Output:
	Graph
	      -10	---------------------
	       -9	---------------------
	       -8	---------------------
	       -7	---------------------
	       -6	---------------------
	       -5	------X--------------
	       -4	-----XXX-------------
	       -3	------XXX------------
	       -2	-------XXX-----------
	       -1	--------XXX----------
	        0	---------XXX---------
	        1	----------XXX--------
	        2	-----------XXX-------
	        3	------------XXX------
	        4	-------------XXX-----
	        5	--------------XX-----
	        6	---------------------
	        7	---------------------
	        8	---------------------
	        9	---------------------
	       10	---------------------
	*/
}
