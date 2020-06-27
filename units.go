package patterns


type filler interface {
	fill() y
}

type Filling y

func (f Filling) fill() y {
	return y(f)
}

// a Pattern with constant value
type Constant Filling

func (p Constant) at(px, py x) y {
	return y(p)
}

type Disc Filling

func (p Disc) at(px, py x) (v y) {
	if px>unitX || py>unitX{
		return 
	}
	x2:=int64(px)
	y2:=int64(py)
	if x2*x2+y2*y2 <= unitX*unitX {
		return y(p)
	}
	return
}

func (p Disc) MaxX() x {
	return unitX
}

type Square	Filling

func (p Square) at(px, py x) (v y) {
	if py < unitX && py >= -unitX && px >= -unitX && px < unitX {
		return y(p)
	}
	return
}

func (p Square) MaxX() x {
	return unitX
}


