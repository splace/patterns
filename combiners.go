package pattern

//import "fmt"

type unlimiteds []Unlimited

func (c unlimiteds) at(px, py x) (total y) {
	for _, p := range c {
		if p == nil {
			continue
		}
		if lp, ok := p.(Limited); ok {
			m := lp.MaxX()
			if px >= m || py >= m || px < -m || py < -m {
				continue
			}
		}
		total = compose(total, p.at(px, py))
		if total.isOpaque() {
			// XXX optimisation only for bool y: put this success as first in search for next time
//			if i != 0 {
//				c[0], c[i] = c[i], c[0]
//			}
			return
		}
	}
	return
}

func (c unlimiteds) MaxX() (max x) {
	for _, p := range c {
		if lp, is := p.(Limited); is {
			if m := lp.MaxX(); m > max {
				max = m
			}
		}
	}
	return
}


// Limited that is made from composing Patterns
type Composite unlimiteds

func (c Composite) at(px, py x) y {
	return unlimiteds(c).at(px,py)
}	

func (c Composite) MaxX() (max x) {
	return unlimiteds(c).MaxX()
}

type UnlimitedComposite unlimiteds

func (c UnlimitedComposite) at(px, py x) (total y) {
	return Composite(c).at(px,py)
}

// a composite of Translated's that are all 'one side' of the origin can be embedded in a Limited to a smaller region. 
func Recentre (c Composite) Limited {
	if len(c)>0 {
		if tp, is := c[0].(Translated); is {
			m := tp.Limited.MaxX()
			maxx,minx,maxy,miny:=tp.X+m,tp.X-m,tp.Y+m,tp.Y-m
			//fmt.Println(m,maxx,minx,maxy,miny)
			for _, p := range c[1:] {
				if tp, is = p.(Translated); !is {
					return c
				}
				m = tp.Limited.MaxX()
				if m+tp.X>maxx {
					maxx=m+tp.X
				}else{
					if tp.X-m<minx{
						minx=tp.X-m
					}
				}
				if m+tp.Y>maxy {
					maxy=m+tp.Y
				}else{
					if tp.Y-m<miny{
						miny=tp.Y-m
					}
				}
			}
			m=max2(maxx-minx,maxy-miny)
			//fmt.Println(maxx,minx,maxy,miny)
			//fmt.Println(m,c.MaxX())
			cm:=c.MaxX()
			if cm<m{
				return c
			}
			return Limiter{UnlimitedTranslated{Limiter{Translated{c,-(maxx+minx)/2,-(maxy+miny)/2},m/2},(maxx+minx)/2,(maxy+miny)/2},cm}
		}
	}
	return c
}



//// a composite of Translated's that are all 'one side' of the origin can be Limited to a smaller region. 
//func Centre (c Composite) Limited {
//	for i, p := range c {
//		if tp, is := p.(Translated); is {
//			m := tp.Limited.MaxX()
//			maxx,minx,maxy,miny:=tp.X+m,tp.X-m,tp.Y+m,tp.Y-m
//			//fmt.Println(m,maxx,minx,maxy,miny)
//			for _, p := range c[i+1:] {
//				if tp, is := p.(Translated); is {
//					m := tp.Limited.MaxX()
//					if m+tp.X>maxx {
//						maxx=m+tp.X
//					}else{
//						if tp.X-m<minx{
//							minx=tp.X-m
//						}
//					}
//					if m+tp.Y>maxy {
//						maxy=m+tp.Y
//					}else{
//						if tp.Y-m<miny{
//							miny=tp.Y-m
//						}
//					}
//				}
//			}
//			m=max2(maxx-minx,maxy-miny)
//			//fmt.Println(maxx,minx,maxy,miny)
//			//fmt.Println(m,c.MaxX())
//			cm:=c.MaxX()
//			if cm<m{
//				return c
//			}
//			return Limiter{UnlimitedTranslated{Limiter{UnlimitedTranslated{c,-(maxx+minx)/2,-(maxy+miny)/2},m/2},(maxx+minx)/2,(maxy+miny)/2},cm}
//		}
//	}
//	return c
//}

