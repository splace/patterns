package pattern

// Draw(er)'s return a Limited
type Drawer interface {
	Draw(*Brush) Limited
}

// Paths are ordered collections of Drawers.
// they are also themselves a Drawer that is a Limiter of a Composite of its Drawers Draw(n) Limited's, in order, using the same Brush.
// a Path can so be an ordered collection of (sub) Paths.
type Path []Drawer

// draw a path using the provided brush
func (p Path) Draw(b *Brush) Limited {
	var c Composite
	if b.StartMarker == nil && b.EndMarker == nil {
		for _, s := range p {
			if d := s.Draw(b); d != nil {
//			if du,is:=d.(Limiter);is{
//				if duc,is:=du.Unlimited.(Composite);is && len(duc)<4{
//					c = append(c, duc...)
//				}else{
//					c = append(c, d)
//				}
//			}else
				if cd,is:=d.(Composite);is{
					c = append(c,cd...)   // XXX leave to end and check if this is good or bad?
				}else{
					c = append(c, d)
				}
			}
		}
		return c //Limiter{Unlimited(c), c.MaxX()}
	}
	// same as above but with markers before first non-nil Draw and/or after last non-nil Draw
	var sx, sy, ex, ey x
	for _, s := range p {
		if len(c) == 0 {
			sx, sy = b.PenPath.Pen.x, b.PenPath.Pen.y
		}
		if d := s.Draw(b); d != nil {
			if len(c) == 0 && b.StartMarker != nil {
				c = append(c, Translated{b.StartMarker, sx, sy})
			}
			// XXX assume sub-composite is locialised to flatten and then centre to regain the division. 
			if cd,is:=d.(Composite);is{
				c = append(c,cd...)   // XXX leave to end and check if this is good or bad?
			}else{
				c = append(c, d)
			}
//			if du,is:=d.(Limiter);is{
//				if duc,is:=du.Unlimited.(Composite);is && len(duc)<4{
//					c = append(c, duc...)
//				}else{
//					c = append(c, d)
//				}
//			}else{
//				c = append(c, d)
//			}
			ex, ey = b.PenPath.Pen.x, b.PenPath.Pen.y
		}
	}
	if b.EndMarker != nil {
		c = append(c, Translated{b.EndMarker, ex, ey})
	}
	return c //Limiter{Unlimited(c), c.MaxX()}
}
