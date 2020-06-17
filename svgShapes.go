package patterns

func ChamferedRectangle(hw,hh,r x) Path {
	owr,ohr:=hw-r,hh-r
	return Path{
		MoveTo{owr,hh},
		LineTo{hw,ohr},
		VerticalLineTo{-ohr},
		LineTo{owr,-hh},
		HorizontalLineTo{-owr},
		LineTo{-hw,-ohr},
		VerticalLineTo{ohr},
		LineTo{-owr,hh},
		Close{},
	}
}

func RoundedRectangle(hw,hh,r x) Path {
	return ArcCorneredRectangle(hw,hh,r,r,0,0,false,false)
}

func ScallopedCorneredRectangle(hw,hh,r x) Path {
	return ArcCorneredRectangle(hw,hh,r,r,0,0,false,true)
}

func BallCorneredRectangle(hw,hh,r x) Path {
	return ArcCorneredRectangle(hw,hh,r,r,0,0,true,false)
}

func ArcCorneredRectangle(hw,hh,rx,ry,dr,a x, large,sweep bool) Path {
	ow,oh:=hw-rx,hh-ry
	rx+=dr*unitX
	ry+=dr*unitX
	var lx,sx x
	if large {lx=unitX}
	if sweep {sx=unitX}
	return Path{
		MoveTo{ow,hh},
		ArcTo{rx,ry,a,lx,sx,hw,oh},
		VerticalLineTo{-oh},
		ArcTo{rx,ry,a,lx,sx,ow,-hh},
		HorizontalLineTo{-ow},
		ArcTo{rx,ry,-a,lx,sx,-hw,-oh},
		VerticalLineTo{oh},
		ArcTo{rx,ry,-a,lx,sx,-ow,hh},
		Close{},
	}
}

