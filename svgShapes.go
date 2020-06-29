package patterns

func ChamferedRectangle(hw, hh, r x) Path {
	owr, ohr := hw-r, hh-r
	return Path{
		MoveTo{owr, hh},
		LineTo{hw, ohr},
		VerticalLineTo{-ohr},
		LineTo{owr, -hh},
		HorizontalLineTo{-owr},
		LineTo{-hw, -ohr},
		VerticalLineTo{ohr},
		LineTo{-owr, hh},
		Close{},
	}
}

func SmoothCorneredTrapezoid(hw, dww, hh, r x) Path {
	ow, oh := hw-r, hh-r
	return Path{
		MoveTo{ow + dww, hh},
		QuadraticBezierTo{hw + dww, hh, hw + dww, oh},
		LineTo{hw - dww, -oh},
		QuadraticBezierTo{hw - dww, -hh, ow - dww, -hh},
		HorizontalLineTo{-ow + dww},
		QuadraticBezierTo{-hw + dww, -hh, -hw + dww, -oh},
		LineTo{-hw - dww, oh},
		QuadraticBezierTo{-hw - dww, hh, -ow - dww, hh},
		Close{},
	}
}

func RoundedRectangle(hw, hh, r x) Path {
	return ArcCorneredTrapezoid(hw, 0, hh, r, r, 0, 0, false, false)
}

func ScallopedCorneredRectangle(hw, hh, r x) Path {
	return ArcCorneredTrapezoid(hw, 0, hh, r, r, 0, 0, false, true)
}

func BallCorneredRectangle(hw, hh, r x) Path {
	return ArcCorneredTrapezoid(hw, 0, hh, r, r, 0, 0, true, false)
}

func ArcCorneredTrapezoid(hw, dww, hh, rx, ry, dr, a x, large, sweep bool) Path {
	ow, oh := hw-rx, hh-ry
	rx += dr * unitX
	ry += dr * unitX
	var lx, sx x
	if large {
		lx = unitX
	}
	if sweep {
		sx = unitX
	}
	return Path{
		MoveTo{ow + dww, hh},
		ArcTo{rx, ry, a, lx, sx, hw + dww, oh},
		LineTo{hw - dww, -oh},
		ArcTo{rx, ry, -a, lx, sx, ow - dww, -hh},
		HorizontalLineTo{-ow + dww},
		ArcTo{rx, ry, a, lx, sx, -hw + dww, -oh},
		LineTo{-hw - dww, oh},
		ArcTo{rx, ry, -a, lx, sx, -ow - dww, hh},
		Close{},
	}
}
