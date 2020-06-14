package patterns

import "fmt"
//import "testing"


func ExampleSVGShapesChamferedBoxPath() {
	p := ChamferedBox(40*unitX,30*unitX, 4*unitX)
	fmt.Println(p)
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

func ExampleSVGShapesRoundedBoxPath() {
	p := RoundedBox(40*unitX,30*unitX, 15*unitX)
	b := NewBrush(Facetted{Width: unitX, In: unitY, CurveDivision:2})
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

func ExampleSVGShapesRoundedBoxPathWithHole() {
	p := RoundedBox(20*unitX,15*unitX, 15*unitX)
	b := NewBrush(Facetted{Width: 5*unitX, In: unitY, CurveDivision:2})
	Output(Limiter{p.Draw(b),22.5*unitX},unitX)
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