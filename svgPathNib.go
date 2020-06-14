package patterns

// a Nib that doesn't actuall produce patterns, but appends to a Path the svg commandds needed to Draw the pattern. 
type SimpleSvgPathNib Path

func (p *SimpleSvgPathNib) Line(x1, y1, x2, y2 x) LimitedPattern {
	*p=append(*p,MoveTo([]x{x1,y1}),LineTo([]x{x2,y2}))
	return nil
}

func (p *SimpleSvgPathNib)	QuadraticBezier(sx, sy, cx, cy, ex, ey x) LimitedPattern{
	*p=append(*p,MoveTo([]x{sx,sy}),QuadraticBezierTo([]x{cx,cy,ex,ey}))
	return nil
}

func (p *SimpleSvgPathNib)	CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey x) LimitedPattern{
	*p=append(*p,MoveTo([]x{sx,sy}),CubicBezierTo([]x{c1x,c1y,c2x,c2y,ex,ey}))
	return nil
}

func (p *SimpleSvgPathNib)	Arc(sx,sy,rx,ry x, a float64, large,sweep bool, ex,ey x) LimitedPattern{
	if large{
		if sweep{
			*p=append(*p,MoveTo([]x{sx,sy}),ArcTo([]x{rx,ry,x(a*unitX),1,1,ex,ey}))
		}else{
			*p=append(*p,MoveTo([]x{sx,sy}),ArcTo([]x{rx,ry,x(a*unitX),1,0,ex,ey}))
		}
	}else{
		if sweep{
			*p=append(*p,MoveTo([]x{sx,sy}),ArcTo([]x{rx,ry,x(a*unitX),0,1,ex,ey}))
		}else{
			*p=append(*p,MoveTo([]x{sx,sy}),ArcTo([]x{rx,ry,x(a*unitX),0,0,ex,ey}))
		}
	}
	return nil
}