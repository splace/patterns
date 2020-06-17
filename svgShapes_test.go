package patterns

import "fmt"
//import "testing"

func ExampleSVGShapesChamferedRectanglePath() {
	p := ChamferedRectangle(20*unitX,15*unitX, 4*unitX)
	fmt.Println(p)
	/* Output:
M16,15
L20,11
V-11
L16,-15
H-16
L-20,-11
V11
L-16,15
Z
*/
}

func ExampleSVGShapesBallCorneredRectanglePath() {
	p := BallCorneredRectangle(20*unitX,15*unitX, 8*unitX)
	fmt.Println(p)
	/* Output:
M12,15
A8,8 0 1 0 20,7
V-7
A8,8 0 1 0 12,-15
H-12
A8,8 0 1 0 -20,-7
V7
A8,8 0 1 0 -12,15
Z
*/
}


func ExampleSVGShapesChamferedRectanglePathGraph() {
	p := ChamferedRectangle(20*unitX,15*unitX, 4*unitX)
	b := NewBrush(Facetted{Width: 2*unitX, In: unitY, CurveDivision:2})
	PrintGraph(p.Draw(b),-25*unitX,25*unitX,-20*unitX,20*unitX,unitX)
	/* Output:
Graph
      -20	---------------------------------------------------
      -19	---------------------------------------------------
      -18	---------------------------------------------------
      -17	---------------------------------------------------
      -16	---------------------------------------------------
      -15	---------XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX---------
      -14	--------X---------------------------------X--------
      -13	-------X-----------------------------------X-------
      -12	------X-------------------------------------X------
      -11	-----X---------------------------------------X-----
      -10	-----X---------------------------------------X-----
       -9	-----X---------------------------------------X-----
       -8	-----X---------------------------------------X-----
       -7	-----X---------------------------------------X-----
       -6	-----X---------------------------------------X-----
       -5	-----X---------------------------------------X-----
       -4	-----X---------------------------------------X-----
       -3	-----X---------------------------------------X-----
       -2	-----X---------------------------------------X-----
       -1	-----X---------------------------------------X-----
        0	-----X---------------------------------------X-----
        1	-----X---------------------------------------X-----
        2	-----X---------------------------------------X-----
        3	-----X---------------------------------------X-----
        4	-----X---------------------------------------X-----
        5	-----X---------------------------------------X-----
        6	-----X---------------------------------------X-----
        7	-----X---------------------------------------X-----
        8	-----X---------------------------------------X-----
        9	-----X---------------------------------------X-----
       10	-----X---------------------------------------X-----
       11	-----X---------------------------------------X-----
       12	------X-------------------------------------X------
       13	-------X-----------------------------------X-------
       14	--------X---------------------------------X--------
       15	---------XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX---------
       16	---------------------------------------------------
       17	---------------------------------------------------
       18	---------------------------------------------------
       19	---------------------------------------------------
       20	---------------------------------------------------
*/
}

func ExampleSVGShapesBallCorneredRectanglePathGraph() {
	p := BallCorneredRectangle(20*unitX,15*unitX, 8*unitX)
	b := NewBrush(Facetted{Width: 2*unitX, In: unitY, CurveDivision:2})
	PrintGraph(p.Draw(b),-35*unitX,35*unitX,-30*unitX,30*unitX,unitX)
	/* Output:
Graph
      -20	---------------------------------------------------
      -19	---------------------------------------------------
      -18	---------------------------------------------------
      -17	---------------------------------------------------
      -16	---------------------------------------------------
      -15	---------XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX---------
      -14	--------X---------------------------------X--------
      -13	-------X-----------------------------------X-------
      -12	------X-------------------------------------X------
      -11	-----X---------------------------------------X-----
      -10	-----X---------------------------------------X-----
       -9	-----X---------------------------------------X-----
       -8	-----X---------------------------------------X-----
       -7	-----X---------------------------------------X-----
       -6	-----X---------------------------------------X-----
       -5	-----X---------------------------------------X-----
       -4	-----X---------------------------------------X-----
       -3	-----X---------------------------------------X-----
       -2	-----X---------------------------------------X-----
       -1	-----X---------------------------------------X-----
        0	-----X---------------------------------------X-----
        1	-----X---------------------------------------X-----
        2	-----X---------------------------------------X-----
        3	-----X---------------------------------------X-----
        4	-----X---------------------------------------X-----
        5	-----X---------------------------------------X-----
        6	-----X---------------------------------------X-----
        7	-----X---------------------------------------X-----
        8	-----X---------------------------------------X-----
        9	-----X---------------------------------------X-----
       10	-----X---------------------------------------X-----
       11	-----X---------------------------------------X-----
       12	------X-------------------------------------X------
       13	-------X-----------------------------------X-------
       14	--------X---------------------------------X--------
       15	---------XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX---------
       16	---------------------------------------------------
       17	---------------------------------------------------
       18	---------------------------------------------------
       19	---------------------------------------------------
       20	---------------------------------------------------
*/
}

func ExampleSVGShapesArcCorneredTrapezoidPathGraph() {
	p := ArcCorneredTrapezoid(30*unitX,5*unitX,25*unitX,8*unitX,4*unitX,0,45*unitX,false,true)
	b := NewBrush(Facetted{Width: 2*unitX, In: unitY, CurveDivision:2})
	PrintGraph(p.Draw(b),-40*unitX,40*unitX,-30*unitX,30*unitX,unitX)
	/* Output:
Graph
      -20	---------------------------------------------------
      -19	---------------------------------------------------
      -18	---------------------------------------------------
      -17	---------------------------------------------------
      -16	---------------------------------------------------
      -15	---------XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX---------
      -14	--------X---------------------------------X--------
      -13	-------X-----------------------------------X-------
      -12	------X-------------------------------------X------
      -11	-----X---------------------------------------X-----
      -10	-----X---------------------------------------X-----
       -9	-----X---------------------------------------X-----
       -8	-----X---------------------------------------X-----
       -7	-----X---------------------------------------X-----
       -6	-----X---------------------------------------X-----
       -5	-----X---------------------------------------X-----
       -4	-----X---------------------------------------X-----
       -3	-----X---------------------------------------X-----
       -2	-----X---------------------------------------X-----
       -1	-----X---------------------------------------X-----
        0	-----X---------------------------------------X-----
        1	-----X---------------------------------------X-----
        2	-----X---------------------------------------X-----
        3	-----X---------------------------------------X-----
        4	-----X---------------------------------------X-----
        5	-----X---------------------------------------X-----
        6	-----X---------------------------------------X-----
        7	-----X---------------------------------------X-----
        8	-----X---------------------------------------X-----
        9	-----X---------------------------------------X-----
       10	-----X---------------------------------------X-----
       11	-----X---------------------------------------X-----
       12	------X-------------------------------------X------
       13	-------X-----------------------------------X-------
       14	--------X---------------------------------X--------
       15	---------XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX---------
       16	---------------------------------------------------
       17	---------------------------------------------------
       18	---------------------------------------------------
       19	---------------------------------------------------
       20	---------------------------------------------------
*/
}

