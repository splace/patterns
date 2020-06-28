package patterns


type Drawer interface{
	Draw(*Brush)Pattern
}

// a Path is a collection of Drawers that uses the same Brush to Draw its contained Drawers in order.
// it is itself a Drawer so a Path can be an ordered collection of (sub) Paths.
// Notice: sub-Paths are drawn with the same Brush so relative Drawers carry on from the end of the previous sub-Path.
type Path []Drawer

// draw a path using the provided brush
func (p Path) Draw(b *Brush) Pattern {
	var c Composite
	if b.StartMarker!=nil{c=append(c,Translated{b.StartMarker,b.PenPath.Pen.x,b.PenPath.Pen.y})}
	for _,s:=range(p){
		if d:=s.Draw(b);d!=nil{
			c=append(c,d)
		}
	}
	if b.EndMarker!=nil{c=append(c,Translated{b.EndMarker,b.PenPath.Pen.x,b.PenPath.Pen.y})}
	return c
}
