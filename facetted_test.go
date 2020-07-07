package patterns

//import "fmt"
//import "testing"

func ExampleFacettedPolygon() {
	f := Facetted{LineNib: LineNib{unitX, unitY}}
	PrintGraph(f.Polygon([][2]x{{0, 0}, {5 * unitX, 5 * unitX}, {5 * unitX, -5 * unitX}}...), -10*unitX, 10*unitX, -10*unitX, 10*unitX, unitX)
	// Output:
	/*
	Graph
	      -10	---------------------
	       -9	---------------------
	       -8	---------------------
	       -7	---------------------
	       -6	---------------------
	       -5	---------------X-----
	       -4	--------------XX-----
	       -3	-------------X-X-----
	       -2	------------X--X-----
	       -1	-----------X---X-----
	        0	----------X----X-----
	        1	-----------X---X-----
	        2	------------X--X-----
	        3	-------------X-X-----
	        4	--------------XX-----
	        5	---------------X-----
	        6	---------------------
	        7	---------------------
	        8	---------------------
	        9	---------------------
	       10	---------------------
	*/
}

func ExampleFacettedQuadraticBezier() {
	f := Facetted{LineNib: LineNib{unitX, unitY}, CurveDivision: 2}
	p := f.QuadraticBezier(-10*unitX, -10*unitX, 0*unitX, 10*unitX, 10*unitX, -10*unitX)
	Output(p, unitX)
	// Output:
	/*
	Graph
	      -10	---------------------
	       -9	---------------------
	       -8	---------------------
	       -7	---------------------
	       -6	---------------------
	       -5	---------------X-----
	       -4	--------------XX-----
	       -3	-------------X-X-----
	       -2	------------X--X-----
	       -1	-----------X---X-----
	        0	----------X----X-----
	        1	-----------X---X-----
	        2	------------X--X-----
	        3	-------------X-X-----
	        4	--------------XX-----
	        5	---------------X-----
	        6	---------------------
	        7	---------------------
	        8	---------------------
	        9	---------------------
	       10	---------------------
	*/
}

func ExampleFacettedCircleArcPrint() {
	f := Facetted{LineNib: LineNib{unitX, unitY}, CurveDivision: 2}
	Output(
		LimitedComposite{
			f.Conic(-1*unitX, 0, 2*unitX, 2*unitX, 0, true, false, 1*unitX, 0),
			f.Conic(-1*unitX, 0, 2*unitX, 2*unitX, 0, false, false, 1*unitX, 0),
			f.Conic(-1*unitX, 0, 2*unitX, 2*unitX, 0, false, true, 1*unitX, 0),
			f.Conic(-1*unitX, 0, 2*unitX, 2*unitX, 0, true, true, 1*unitX, 0),
		},
		unitX*0.25,
	)
	// Output:
	/*
	Graph
	      -10	---------------------
	       -9	---------------------
	       -8	---------------------
	       -7	---------------------
	       -6	---------------------
	       -5	---------------X-----
	       -4	--------------XX-----
	       -3	-------------X-X-----
	       -2	------------X--X-----
	       -1	-----------X---X-----
	        0	----------X----X-----
	        1	-----------X---X-----
	        2	------------X--X-----
	        3	-------------X-X-----
	        4	--------------XX-----
	        5	---------------X-----
	        6	---------------------
	        7	---------------------
	        8	---------------------
	        9	---------------------
	       10	---------------------
	*/
}

func ExampleFacettedEllipseArcPrint() {
	f := Facetted{LineNib: LineNib{unitX, unitY}, CurveDivision: 2}
	Output(
		LimitedComposite{
			f.Conic(-1*unitX, 0, 2*unitX, 4*unitX, 0, true, false, 1*unitX, 0),
			f.Conic(-1*unitX, 0, 2*unitX, 4*unitX, 0, false, false, 1*unitX, 0),
			f.Conic(-1*unitX, 0, 2*unitX, 4*unitX, 0, false, true, 1*unitX, 0),
			f.Conic(-1*unitX, 0, 2*unitX, 4*unitX, 0, true, true, 1*unitX, 0),
		},
		unitX*0.25,
	)
	// Output:
	/*
	Graph
	      -10	---------------------
	       -9	---------------------
	       -8	---------------------
	       -7	---------------------
	       -6	---------------------
	       -5	---------------X-----
	       -4	--------------XX-----
	       -3	-------------X-X-----
	       -2	------------X--X-----
	       -1	-----------X---X-----
	        0	----------X----X-----
	        1	-----------X---X-----
	        2	------------X--X-----
	        3	-------------X-X-----
	        4	--------------XX-----
	        5	---------------X-----
	        6	---------------------
	        7	---------------------
	        8	---------------------
	        9	---------------------
	       10	---------------------
	*/
}

func ExampleFacettedEllipseRotatedArcPrint() {
	f := Facetted{LineNib: LineNib{unitX, unitY}, CurveDivision: 3}
	Output(
		LimitedComposite{
			f.Conic(-2*unitX, 0, 8*unitX, 2*unitX, 1, true, false, 2*unitX, 0),
			f.Conic(-2*unitX, 0, 8*unitX, 2*unitX, 1, false, false, 2*unitX, 0),
			f.Conic(-2*unitX, 0, 8*unitX, 2*unitX, 1, false, true, 2*unitX, 0),
			f.Conic(-2*unitX, 0, 8*unitX, 2*unitX, 1, true, true, 2*unitX, 0),
		},
		unitX/3,
	)
	// Output:
	/*
	Graph
	      -10	---------------------
	       -9	---------------------
	       -8	---------------------
	       -7	---------------------
	       -6	---------------------
	       -5	---------------X-----
	       -4	--------------XX-----
	       -3	-------------X-X-----
	       -2	------------X--X-----
	       -1	-----------X---X-----
	        0	----------X----X-----
	        1	-----------X---X-----
	        2	------------X--X-----
	        3	-------------X-X-----
	        4	--------------XX-----
	        5	---------------X-----
	        6	---------------------
	        7	---------------------
	        8	---------------------
	        9	---------------------
	       10	---------------------
	*/
}
