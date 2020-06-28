package patterns

// Facetted is a Nib producing curves using a number of straight lines.
// curves are divided according to CurveDivision:  (power of 2 number of divisions.)
// default 0 - no division, all curves are a single straight line
// if a Nib is provided its Straight method is used to draw the straight lines.
type Facetted struct{
	Nib
	LineNib
	CurveDivision uint8
//	Lwidth x // last width to make tappered lines
}

func (f Facetted) Straight(x1, y1, x2, y2 x) LimitedPattern {
	if f.Nib!=nil{
		return f.Nib.Straight(x1, y1, x2, y2)
	}
	return f.LineNib.Straight(x1, y1, x2, y2)
}

func (f Facetted) Curved(sx, sy, c1x, c1y, c2x,c2y, ex, ey x) LimitedPattern {
	if c1x==c2x && c1y==c2y {
		return  f.QuadraticBezier(sx, sy, c1x, c1y, ex, ey)
	}
	return f.CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey)
}

func (f Facetted) QuadraticBezier(sx, sy, cx, cy, ex, ey x) LimitedPattern {
		return f.polygon(sx,sy,ex,ey,Divide(1<<f.CurveDivision).QuadraticBezier(sx, sy, cx, cy, ex, ey))
}

func (f Facetted) CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey x) LimitedPattern {
	return f.polygon(sx,sy,ex,ey,Divide(1<<f.CurveDivision).CubicBezier(sx, sy, c1x, c1y, c2x,c2y, ex, ey))
}

func (f Facetted) QuinticBezier(sx, sy, c1x, c1y, c2x,c2y, c3x,c3y, ex, ey x) LimitedPattern {
	return f.polygon(sx,sy,ex,ey,Divide(1<<f.CurveDivision).QuinticBezier(sx, sy, c1x, c1y, c2x,c2y, c3x, c3y,ex, ey))
}

func (f Facetted) Conic(sx,sy,rx,ry x, a float64, large,sweep bool, ex,ey x) LimitedPattern {
	return f.polygon(sx,sy,ex,ey,Divide(1<<f.CurveDivision).Arc(sx,sy,rx,ry, a, large,sweep, ex,ey))
}

func (f Facetted) polygon(sx,sy,ex,ey x, pts <- chan [2]x) LimitedPattern {
	var s []Pattern
	joiner:=Shrunk{Disc(Filling(f.In)),2*unitX/float32(f.Width) }
	l:=Limits{sx,sy,sx,sy}
	for p:=range(pts){
		s= append(s,f.Straight(sx,sy,p[0],p[1]))
		s= append(s,Translated{&joiner,p[0],p[1]})
		sx,sy=p[0],p[1]
		l.Include(p)
	}
	s= append(s,f.Straight(sx,sy,ex,ey))
	l.Include([2]x{ex,ey})
	return Translated{Limiter{UnlimitedTranslated{NewComposite(s...),(l.MaxX+l.MinX)>>1,(l.MaxY+l.MinY)>>1},max((l.MaxX-l.MinX)>>1,(l.MaxY-l.MinY)>>1)+f.Width},-((l.MaxX+l.MinX)>>1),-((l.MaxY+l.MinY)>>1) }
}

func (f Facetted) Box(x,y x) LimitedPattern {
	return Limiter{Composite{f.Straight(-x,y, x,y),f.Straight(x,y,x,-y),f.Straight(x,-y,-x,-y),f.Straight(-x,-y,-x,y)},max(x+f.Width,y+f.Width)}
}

func (f Facetted) Polygon(coords ...[2]x) LimitedPattern {
	s := make([]Pattern, len(coords)) 
	l:=Limits{coords[0][0], coords[0][1],coords[0][0], coords[0][1]}
	for i := 1; i < len(s); i++ {
		s[i-1] = f.Straight(coords[i-1][0], coords[i-1][1],coords[i][0], coords[i][1])
		l.Include([2]x{coords[i][0], coords[i][1]})
	}
	s[len(coords)-1] = f.Straight(coords[len(coords)-1][0], coords[len(coords)-1][1],coords[0][0], coords[0][1])
	return Translated{Limiter{UnlimitedTranslated{NewComposite(s...),(l.MaxX+l.MinX)>>1,(l.MaxY+l.MinY)>>1},max((l.MaxX-l.MinX)>>1,(l.MaxY-l.MinY)>>1)+f.Width},-((l.MaxX+l.MinX)>>1),-((l.MaxY+l.MinY)>>1) } 
}


// max and min points
type Limits struct{
	MinX,MinY,MaxX,MaxY x
}

func (d *Limits) Include(p [2]x) {
	if p[0]<d.MinX{
		d.MinX=p[0]
	}else{
		if p[0]>d.MaxX{d.MaxX=p[0]}
	}
	if p[1]<d.MinY{
		d.MinY=p[1]
	}else{
		if p[1]>d.MaxY{d.MaxY=p[1]}
	}
}

