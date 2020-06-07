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


//func TestFacettedBox(t *testing.T) {
//	if cpt!=fmt.Sprint(CompactPath(p2)){
//		t.Errorf("%q != %q (%v)",cpt,fmt.Sprint(CompactPath(p2)),p2)
//	}
//}
