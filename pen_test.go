package patterns

import "fmt"

func ExamplePenLineZeroLength() {
	p := Pen{Nib: LineNib{unitX, unitY}}
	l:=p.Straight(0, 0, 0, 0)
	fmt.Printf("%v\n",l)
	// Output:
	// {{{true +Inf 2} 0 1} 0 0}
}


func ExamplePenLine() {
	p := Pen{Nib: LineNib{unitX, unitY}}
	Output(p.Straight(0, 0, 5*unitX, 5*unitX), unitX)
	/* Output:
	Graph
	       -7	---------------
	       -6	---------------
	       -5	---------------
	       -4	---------------
	       -3	---------------
	       -2	---------------
	       -1	---------------
	        0	---------------
	        0	--------X------
	        1	---------X-----
	        2	----------X----
	        3	-----------X---
	        4	------------X--
	        5	---------------
	        6	---------------
	*/
}

func ExamplePenLineNonZeroStart() {
	p := Pen{Nib: LineNib{unitX, unitY}}
	Output(p.Straight(5*unitX, -5*unitX, -10*unitX, 10*unitX), unitX)
	/* Output:
	Graph
	      -14	-----------------------------
	      -13	-----------------------------
	      -12	-----------------------------
	      -11	-----------------------------
	      -10	-----------------------------
	       -9	-----------------------------
	       -8	-----------------------------
	       -7	-----------------------------
	       -6	-----------------------------
	       -5	-------------------X---------
	       -4	------------------X----------
	       -3	-----------------X-----------
	       -2	----------------X------------
	       -1	---------------X-------------
	        0	--------------X--------------
	        0	-------------X---------------
	        1	------------X----------------
	        2	-----------X-----------------
	        3	----------X------------------
	        4	---------X-------------------
	        5	--------X--------------------
	        6	-------X---------------------
	        7	------X----------------------
	        8	-----X-----------------------
	        9	----X------------------------
	       10	-----------------------------
	       11	-----------------------------
	       12	-----------------------------
	       13	-----------------------------
	*/
}

func ExamplePenLineCross() {
	p := Pen{Nib: LineNib{unitX, unitY}}
	Output(
		Limiter{
			Composite{
				p.Straight(-5*unitX, 0, -10*unitX, 0),
				p.Straight(5*unitX, 0, 10*unitX, 0),
				p.Straight(0, -5*unitX, 0, -10*unitX),
				p.Straight(0, 5*unitX, 0, 10*unitX),
			},
			12 * unitX,
		},
		unitX,
	)
	/* Output:
	Graph
	      -13	---------------------------
	      -12	---------------------------
	      -11	---------------------------
	      -10	-------------X-------------
	       -9	-------------X-------------
	       -8	-------------X-------------
	       -7	-------------X-------------
	       -6	-------------X-------------
	       -5	-------------X-------------
	       -4	---------------------------
	       -3	---------------------------
	       -2	---------------------------
	       -1	---------------------------
	        0	---XXXXXX---------XXXXX----
	        1	---------------------------
	        2	---------------------------
	        3	---------------------------
	        4	---------------------------
	        5	-------------X-------------
	        6	-------------X-------------
	        7	-------------X-------------
	        8	-------------X-------------
	        9	-------------X-------------
	       10	---------------------------
	       11	---------------------------
	       12	---------------------------
	       13	---------------------------
	*/
}

func ExamplePenPolygon() {
	p := Pen{Nib: LineNib{unitX, unitY}}
	Output(Limiter{p.Nib.(LineNib).Polygon([2]x{0, 5 * unitX}, [2]x{7 * unitX, -7 * unitX}, [2]x{-10 * unitX, 5 * unitX}), 10 * unitX}, unitX)
	/* Output:
	Graph
	      -11	-----------------------
	      -10	-----------------------
	       -9	-----------------------
	       -8	-----------------------
	       -7	------------------X----
	       -6	----------------XX-----
	       -5	---------------XXX-----
	       -4	-------------XX-X------
	       -3	------------XX--X------
	       -2	-----------X---X-------
	       -1	---------XX---XX-------
	        0	--------X-----X--------
	        1	------XX-----X---------
	        2	-----XX------X---------
	        3	---XX-------X----------
	        4	--XX--------X----------
	        5	-XXXXXXXXXXX-----------
	        6	-----------------------
	        7	-----------------------
	        8	-----------------------
	        9	-----------------------
	       10	-----------------------
	       11	-----------------------
	*/
}

func ExamplePenJoinedPath() {
	p := Pen{Nib: LineNib{4 * unitX, unitY}, Joiner: Shrunk{Disc(unitY), .2}}

	p.MoveTo(-15*unitX, -5*unitX)
	Output(
		Limiter{
			Composite{
				p.LineTo(0,5*unitX),
				p.LineTo(15*unitX, -5*unitX),
			},
			unitX * 15,
		},
		unitX)
	/* Output:
	Graph
	      -11	-----------------------
	      -10	-----------------------
	       -9	-----------------------
	       -8	-----------------------
	       -7	------------------X----
	       -6	----------------XX-----
	       -5	---------------XXX-----
	       -4	-------------XX-X------
	       -3	------------XX--X------
	       -2	-----------X---X-------
	       -1	---------XX---XX-------
	        0	--------X-----X--------
	        1	------XX-----X---------
	        2	-----XX------X---------
	        3	---XX-------X----------
	        4	--XX--------X----------
	        5	-XXXXXXXXXXX-----------
	        6	-----------------------
	        7	-----------------------
	        8	-----------------------
	        9	-----------------------
	       10	-----------------------
	       11	-----------------------
	*/
}
