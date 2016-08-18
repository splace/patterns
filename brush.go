package patterns

import (
	"math"
)

// brush holds state, style/position , for line based patterns
type Brush struct {
	Width    x
	In       y
	Relative bool
	x, y     x
	sx, sy   x
}

func (p Brush) Line(px1, py1, px2, py2 x) LimitedPattern {
	length := float32(math.Sqrt(float64(px2-px1)*float64(px2-px1) + float64(py2-py1)*float64(py2-py1)))
	
	return Translated{NewRotated(Reduced{Square{Filling{p.In}}, float32(unitX*2)/length,float32(unitX*2)/float32(p.Width) },math.Atan2(float64(px1-px2),float64(py2-py1)) ).(LimitedPattern),(px1+px2)/2, (py1+py2)/2}
}

func (p *Brush) LineTo(px, py x) LimitedPattern {
	if p.Relative {
		px += p.x
		py += p.y
	}
	s := p.Line(p.x, p.y, px, py)
	p.x, p.y = px, py
	return s
}

func (p *Brush) LineToVert(py x) LimitedPattern {
	if p.Relative {
		py += p.y
	}
	s := p.Line(p.x, p.y, p.x, py)
	p.y = py
	return s
}

func (p *Brush) LineToHor(px x) LimitedPattern {
	if p.Relative {
		px += p.x
	}
	s := p.Line(p.x, p.y, px, p.y)
	p.x = px
	return s
}

func (p *Brush) StartLine(px1, py1, px2, py2 x) LimitedPattern {
	p.MoveTo(px1,py1)
	return p.LineTo(px2, py2)
}

func (p *Brush) MoveTo(px, py x) {
	if p.Relative {
		px += p.x
		py += p.y
	}
	p.x, p.y, p.sx, p.sy = px, py, px, py
	return
}

func (p *Brush) CloseLine() LimitedPattern {
	// TODO add join to end
	s := p.Line(p.x, p.y, p.sx, p.sy)
	p.x, p.y = p.sx, p.sy
	return s
}

func (p *Brush) Polygon(coords [][2]x) Pattern {
	s := make([]Pattern, len(coords)) 
	s[0] = p.StartLine(coords[0][0], coords[0][1], coords[1][0], coords[1][1])
	for i := 2; i < len(coords); i++ {
		s[i-1] = p.LineTo(coords[i][0], coords[i][1])
	}
	s[len(coords)-1] = p.CloseLine()
	return NewComposite(s...)
}


