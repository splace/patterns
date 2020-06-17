package patterns

func ChamferedBox(hw,hh,r x) Path {
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

func RoundedBox(hw,hh,r x) Path {
	ow,oh:=hw-r,hh-r
	return Path{
		MoveTo{ow,hh},
		ArcTo{r,r,0,0,0,hw,oh},
		VerticalLineTo{-oh},
		ArcTo{r,r,0,0,0,ow,-hh},
//		HorizontalLineTo{-ow},
//		ArcTo{r,r,0,0,1,-hw,-oh},
//		VerticalLineTo{oh},
//		ArcTo{r,r,0,0,1,-ow,hh},
//		Close{},
	}
}

