package patterns

// Draw(er)'s return a Pattern because a Brush can have different Nibs which can effect the Limits of the Pattern
type Drawer interface {
	Draw(*Brush) Pattern
}

// Paths are ordered collections of Drawers.
// they are also themselves a Drawer whose Pattern is a Composite of its Drawers Draw(n) Patterns in order using the same Brush.
// a Path can so be an ordered collection of (sub) Paths.
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
	// same as above but with markers before first non-nil Draw and/or after last non-nil Draw
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
