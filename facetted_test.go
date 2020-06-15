package patterns

//import "fmt"
//import "testing"

func ExampleFacettedPolygon() {
	f:=Facetted{Width:unitX, In:unitY}
	PrintGraph(f.Polygon([][2]x{{0,0},{5*unitX,5*unitX},{5*unitX,-5*unitX}}...),-10*unitX,10*unitX,-10*unitX,10*unitX,unitX)
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


func ExampleFacettedArc() {
	new(Facetted).Arc(-10,0,1,21,90,false,false,10,0)
	// Output:
	//
}

func ExampleFacettedCircle() {
	new(Facetted).Arc(-3,0,5,5,0,false,false,3,0)
	new(Facetted).Arc(-3,0,5,5,0,false,true,3,0)
	new(Facetted).Arc(-3,0,5,5,0,true,false,3,0)
	new(Facetted).Arc(-3,0,5,5,0,true,true,3,0)
	// Output:
	//
}

func ExampleFacettedArcPrint() {
	f:=Facetted{Width:.2*unitX, In:unitY,CurveDivision:2}
	Output(
		LimitedComposite{
			f.Arc(-1*unitX,0,2*unitX,2*unitX,0,true,false,1*unitX,0),
			f.Arc(-1*unitX,0,2*unitX,2*unitX,0,false,false,1*unitX,0),
			f.Arc(-1*unitX,0,2*unitX,2*unitX,0,false,true,1*unitX,0),
			f.Arc(-1*unitX,0,2*unitX,2*unitX,0,true,true,1*unitX,0),
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
