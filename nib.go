package patterns


// Nibs can create LimitedPatterns from lines and curves.
type Nib interface{
	Line(x,x,x,x) LimitedPattern
	QuadraticBezier(x,x,x,x,x,x) LimitedPattern
	CubicBezier(x,x,x,x,x,x,x,x) LimitedPattern
	Arc(x,x,x,x,float64,bool,bool,x,x) LimitedPattern
}

