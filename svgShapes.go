package patterns

func ChamferedBoxPath(hw,hh,r x) Path {
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

func RoundedBoxPath(hw,hh,r x) Path {
	owr,ohr:=hw-r,hh-r
	return Path{
		MoveTo{owr,hh},
		QuadraticBezierTo{hw,hh,hw,ohr},
		VerticalLineTo{-ohr},
		QuadraticBezierTo{hw,-hh,owr,-hh},
		HorizontalLineTo{-owr},
		QuadraticBezierTo{-hw,-hh,-hw,-ohr},
		VerticalLineTo{ohr},
		QuadraticBezierTo{-hw,hh,-owr,hh},
		Close{},
	}
}

