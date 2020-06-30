package patterns

type Drawer interface {
	Draw(*Brush) Pattern
}

// a Path is a collection of Drawers that uses the same Brush to Draw its contained Drawers in order.
// it is itself a Drawer so a Path can be an ordered collection of (sub) Paths.
// Notice: sub-Paths are drawn with the same Brush so relative Drawers carry on from the end of the previous sub-Path.
type Path []Drawer

// draw a path using the provided brush
func (p Path) Draw(b *Brush) Pattern {
	var c Composite
	if b.StartMarker == nil && b.EndMarker == nil  {
		for _, s := range p {
			if d := s.Draw(b); d != nil {
				c = append(c, d)
			}
		}
		return c
	}
	// add markers before first non-nil Draw and/or after last non-nil Draw
	var sx,sy,ex,ey x
	for _, s := range p {
		if len(c)==0 {
			sx,sy= b.PenPath.Pen.x, b.PenPath.Pen.y
		}
		if d := s.Draw(b); d != nil {
			if len(c)==0 && b.StartMarker!=nil{
				c = append(c, Translated{b.StartMarker, sx, sy})
			}
			c = append(c, d)
			ex,ey=b.PenPath.Pen.x, b.PenPath.Pen.y
		}
	}
	if b.EndMarker != nil {
		c = append(c, Translated{b.EndMarker, ex, ey})
	}
	return c
}
