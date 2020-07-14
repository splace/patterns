package pattern

import "fmt"
import "io"
import "strings"

// Path fmt.Stringer using one command per line.
func (p Path) String() string {
	b := new(strings.Builder)
	for _, s := range p {
		switch ts := s.(type) {
		case MoveTo:
			fmt.Fprintf(b, "M%v,%v\n", ts[0], ts[1])
		case MoveToRelative:
			fmt.Fprintf(b, "m%v,%v\n", ts[0], ts[1])
		case LineTo:
			fmt.Fprintf(b, "L%v,%v\n", ts[0], ts[1])
		case LineToRelative:
			fmt.Fprintf(b, "l%v,%v\n", ts[0], ts[1])
		case VerticalLineTo:
			fmt.Fprintf(b, "V%v\n", ts[0])
		case VerticalLineToRelative:
			fmt.Fprintf(b, "v%v\n", ts[0])
		case HorizontalLineTo:
			fmt.Fprintf(b, "H%v\n", ts[0])
		case HorizontalLineToRelative:
			fmt.Fprintf(b, "h%v\n", ts[0])
		case CloseRelative:
			fmt.Fprintf(b, "z\n")
		case Close:
			fmt.Fprintf(b, "Z\n")
		case QuadraticBezierTo:
			fmt.Fprintf(b, "Q%v,%v %v,%v\n", ts[0], ts[1], ts[2], ts[3])
		case SmoothQuadraticBezierTo:
			fmt.Fprintf(b, "T%v,%v\n", ts[0], ts[1])
		case QuadraticBezierToRelative:
			fmt.Fprintf(b, "q%v,%v %v,%v\n", ts[0], ts[1], ts[2], ts[3])
		case SmoothQuadraticBezierToRelative:
			fmt.Fprintf(b, "t%v,%v\n", ts[0], ts[1])
		case CubicBezierTo:
			fmt.Fprintf(b, "C%v,%v %v,%v %v,%v\n", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5])
		case SmoothCubicBezierTo:
			fmt.Fprintf(b, "S%v,%v %v,%v\n", ts[0], ts[1], ts[2], ts[3])
		case CubicBezierToRelative:
			fmt.Fprintf(b, "c%v,%v %v,%v %v,%v\n", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5])
		case SmoothCubicBezierToRelative:
			fmt.Fprintf(b, "s%v,%v %v,%v\n", ts[0], ts[1], ts[2], ts[3])
		case ArcTo:
			fmt.Fprintf(b, "A%v,%v %v %v %v %v,%v\n", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5], ts[6])
		case ArcToRelative:
			fmt.Fprintf(b, "a%v,%v %v %v %v %v,%v\n", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5], ts[6])
		case Path:
			fmt.Fprintf(b, "%v\n", MaxCompactStringer(ts))
		}
	}
	return b.String()[:b.Len()-1]
}

// Scan in a Path.
// accepts SVG path format options/compression.
func (p *Path) Scan(state fmt.ScanState, r rune) (err error) {
	var d []x // all Drawer's returned are slices into the same underlying x-typed array
	var lc, c rune
	for {
		state.SkipSpace()
		c, _, err = state.ReadRune()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		if c == ',' {
			state.SkipSpace()
			c, _, err = state.ReadRune()
			if err != nil {
				if err == io.EOF {
					return nil
				}
				return err
			}
		}

		for {
			switch c {
			case 'M':
				d = append(d, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, MoveTo(d[len(d)-2:]))
			case 'm':
				d = append(d, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, MoveToRelative(d[len(d)-2:]))
			case 'L':
				d = append(d, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, LineTo(d[len(d)-2:]))
			case 'l':
				d = append(d, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, LineToRelative(d[len(d)-2:]))
			case 'Z':
				*p = append(*p, Close{})
			case 'z':
				*p = append(*p, CloseRelative{})
			case 'H':
				d = append(d, 0)
				_, err = fmt.Fscan(state, &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, HorizontalLineTo(d[len(d)-1:]))
			case 'h':
				d = append(d, 0)
				_, err = fmt.Fscan(state, &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, HorizontalLineToRelative(d[len(d)-1:]))
			case 'V':
				d = append(d, 0)
				_, err = fmt.Fscan(state, &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, VerticalLineTo(d[len(d)-1:]))
			case 'v':
				d = append(d, 0)
				_, err = fmt.Fscan(state, &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, VerticalLineToRelative(d[len(d)-1:]))
			case 'Q': // quadratic Bézier curve
				d = append(d, 0, 0, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-4], &d[len(d)-3], &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, QuadraticBezierTo(d[len(d)-4:]))
			case 'q': // quadratic Bézier curve relative
				d = append(d, 0, 0, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-4], &d[len(d)-3], &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, QuadraticBezierToRelative(d[len(d)-4:]))
			case 'C': // cubic Bézier curve
				d = append(d, 0, 0, 0, 0, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-6], &d[len(d)-5], &d[len(d)-4], &d[len(d)-3], &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, CubicBezierTo(d[len(d)-6:]))
			case 'c': // cubic Bézier curve relative
				d = append(d, 0, 0, 0, 0, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-6], &d[len(d)-5], &d[len(d)-4], &d[len(d)-3], &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, CubicBezierToRelative(d[len(d)-6:]))
			case 'T': // smooth quadratic Bézier curveto
				d = append(d, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, SmoothQuadraticBezierTo(d[len(d)-2:]))
			case 't': // smooth quadratic Bézier curveto relative
				d = append(d, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, SmoothQuadraticBezierToRelative(d[len(d)-2:]))
			case 'S': // smooth cubic Bézier curve
				d = append(d, 0, 0, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-4], &d[len(d)-3], &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, SmoothCubicBezierTo(d[len(d)-4:]))
			case 's': // smooth cubic Bézier curve relative
				d = append(d, 0, 0, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-4], &d[len(d)-3], &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				*p = append(*p, SmoothCubicBezierToRelative(d[len(d)-4:]))
			case 'A': // elliptical Arc
				d = append(d, 0, 0, 0, 0, 0, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-7], &d[len(d)-6], &d[len(d)-5], &d[len(d)-4], &d[len(d)-3], &d[len(d)-2], &d[len(d)-1])
				if d[len(d)-3] != 0 && d[len(d)-3] != unitX || d[len(d)-4] != 0 && d[len(d)-4] != unitX {
					return fmt.Errorf("Arc flags not 0 or 1: %v %v", d[len(d)-3], d[len(d)-4])
				}
				if err != nil {
					return err
				}
				*p = append(*p, ArcTo(d[len(d)-7:]))
			case 'a': // elliptical Arc relative
				d = append(d, 0, 0, 0, 0, 0, 0, 0)
				_, err = fmt.Fscan(state, &d[len(d)-7], &d[len(d)-6], &d[len(d)-5], &d[len(d)-4], &d[len(d)-3], &d[len(d)-2], &d[len(d)-1])
				if err != nil {
					return err
				}
				if d[len(d)-3] != 0 && d[len(d)-3] != unitX || d[len(d)-4] != 0 && d[len(d)-4] != unitX {
					return fmt.Errorf("Arc flags not 0 or 1: %v %v", d[len(d)-3], d[len(d)-4])
				}
				*p = append(*p, ArcToRelative(d[len(d)-7:]))
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', '-', '+':
				state.UnreadRune()
				switch lc {
				case 'M':
					c = 'L'
				case 'm':
					c = 'l'
				default:
					// numeric parameter so use previous command
					c = lc
				}
				continue
			default:
				return fmt.Errorf("Unknown command:%v", string(c))
			}
			break
		}
		lc = c
	}
	return nil
}

// a derivative of Path with a fmt.Stringer producing skipped repeat command letters and one command type per line.
type CompactStringer Path

func (p CompactStringer) String() string {
	b := new(strings.Builder)
	var ls Drawer
	for _, s := range p {
		switch ts := s.(type) {
		case MoveTo:
			fmt.Fprintf(b, "\nM%v,%v", ts[0], ts[1])
		case MoveToRelative:
			fmt.Fprintf(b, "\nm%v,%v", ts[0], ts[1])
		case LineTo:
			switch ls.(type) {
			case LineTo, MoveTo:
				fmt.Fprintf(b, " %v,%v", ts[0], ts[1])
			default:
				fmt.Fprintf(b, "\nL%v,%v", ts[0], ts[1])
			}
		case LineToRelative:
			switch ls.(type) {
			case LineToRelative, MoveToRelative:
				fmt.Fprintf(b, " %v,%v", ts[0], ts[1])
			default:
				fmt.Fprintf(b, "\nl%v,%v", ts[0], ts[1])
			}
		case VerticalLineTo:
			fmt.Fprintf(b, "\nV%v", ts[0])
		case VerticalLineToRelative:
			fmt.Fprintf(b, "\nv%v", ts[0])
		case HorizontalLineTo:
			fmt.Fprintf(b, "\nH%v", ts[0])
		case HorizontalLineToRelative:
			fmt.Fprintf(b, "\nh%v", ts[0])
		case CloseRelative:
			fmt.Fprintf(b, "\nz")
		case Close:
			fmt.Fprintf(b, "\nZ")
		case QuadraticBezierTo:
			switch ls.(type) {
			case QuadraticBezierTo:
				fmt.Fprintf(b, " %v,%v %v,%v", ts[0], ts[1], ts[2], ts[3])
			default:
				fmt.Fprintf(b, "\nQ%v,%v %v,%v", ts[0], ts[1], ts[2], ts[3])
			}
		case SmoothQuadraticBezierTo:
			switch ls.(type) {
			case SmoothQuadraticBezierTo:
				fmt.Fprintf(b, " %v,%v", ts[0], ts[1])
			default:
				fmt.Fprintf(b, "\nT%v,%v", ts[0], ts[1])
			}
		case QuadraticBezierToRelative:
			switch ls.(type) {
			case QuadraticBezierToRelative:
				fmt.Fprintf(b, " %v,%v %v,%v", ts[0], ts[1], ts[2], ts[3])
			default:
				fmt.Fprintf(b, "\nq%v,%v %v,%v", ts[0], ts[1], ts[2], ts[3])
			}
		case SmoothQuadraticBezierToRelative:
			switch ls.(type) {
			case SmoothQuadraticBezierToRelative:
				fmt.Fprintf(b, " %v,%v", ts[0], ts[1])
			default:
				fmt.Fprintf(b, "\nt%v,%v", ts[0], ts[1])
			}
		case CubicBezierTo:
			switch ls.(type) {
			case CubicBezierTo:
				fmt.Fprintf(b, " %v,%v %v,%v %v,%v", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5])
			default:
				fmt.Fprintf(b, "\nC%v,%v %v,%v %v,%v", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5])
			}
		case SmoothCubicBezierTo:
			switch ls.(type) {
			case CubicBezierTo:
				fmt.Fprintf(b, " %v,%v %v,%v", ts[0], ts[1], ts[2], ts[3])
			default:
				fmt.Fprintf(b, "\nS%v,%v %v,%v", ts[0], ts[1], ts[2], ts[3])
			}
		case CubicBezierToRelative:
			switch ls.(type) {
			case CubicBezierToRelative:
				fmt.Fprintf(b, " %v,%v %v,%v %v,%v", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5])
			default:
				fmt.Fprintf(b, "\nc%v,%v %v,%v %v,%v", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5])
			}
		case SmoothCubicBezierToRelative:
			switch ls.(type) {
			case SmoothCubicBezierToRelative:
				fmt.Fprintf(b, " %v,%v %v,%v", ts[0], ts[1], ts[2], ts[3])
			default:
				fmt.Fprintf(b, "\ns%v,%v %v,%v", ts[0], ts[1], ts[2], ts[3])
			}
		case ArcTo:
			switch ls.(type) {
			case ArcTo:
				fmt.Fprintf(b, " %v,%v %v %v %v %v,%v", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5], ts[6])
			default:
				fmt.Fprintf(b, "\nA%v,%v %v %v %v %v,%v", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5], ts[6])
			}
		case ArcToRelative:
			switch ls.(type) {
			case ArcToRelative:
				fmt.Fprintf(b, " %v,%v %v %v %v %v,%v", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5], ts[6])
			default:
				fmt.Fprintf(b, "\na%v,%v %v %v %v %v,%v", ts[0], ts[1], ts[2], ts[3], ts[4], ts[5], ts[6])
			}
		case Path:
			fmt.Fprintf(b, "\n%v", CompactStringer(ts))
		}
		ls = s
	}
	return b.String()[1:]
}

// a derivative of Path with a fmt.Stringer producing SVG path compact form.
// skips repeated command letters/leading zeros/spaces/commas (if number starts with point or neg. sign)
type MaxCompactStringer Path

type addPrefixCommaWhenNecessery x

func (cx addPrefixCommaWhenNecessery) String() (s string) {
	s = fmt.Sprint(x(cx))
	switch s[0] {
	case '0':
		if len(s) > 1 {
			return "," + s[1:]
		}
		return "," + s
	case '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return "," + s
	}
	return
}

type noLeadingZerox x

func (cx noLeadingZerox) String() (s string) {
	s = fmt.Sprint(x(cx))
	if len(s) > 1 && s[0] == '0' {
		return s[1:]
	}
	return
}

func (p MaxCompactStringer) String() string {
	b := new(strings.Builder)
	var ls Drawer
	for _, s := range p {
		switch ts := s.(type) {
		case MoveTo:
			fmt.Fprintf(b, "M%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
		case MoveToRelative:
			fmt.Fprintf(b, "m%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
		case LineTo:
			switch ls.(type) {
			case LineTo, MoveTo:
				fmt.Fprintf(b, "%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
			default:
				fmt.Fprintf(b, "L%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
			}
		case LineToRelative:
			switch ls.(type) {
			case LineToRelative, MoveToRelative:
				fmt.Fprintf(b, "%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
			default:
				fmt.Fprintf(b, "l%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
			}
		case VerticalLineTo:
			fmt.Fprintf(b, "V%v", noLeadingZerox(ts[0]))
		case VerticalLineToRelative:
			fmt.Fprintf(b, "v%v", noLeadingZerox(ts[0]))
		case HorizontalLineTo:
			fmt.Fprintf(b, "H%v", noLeadingZerox(ts[0]))
		case HorizontalLineToRelative:
			fmt.Fprintf(b, "h%v", noLeadingZerox(ts[0]))
		case CloseRelative:
			fmt.Fprintf(b, "z")
		case Close:
			fmt.Fprintf(b, "Z")
		case QuadraticBezierTo:
			switch ls.(type) {
			case QuadraticBezierTo:
				fmt.Fprintf(b, "%v%v%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]))
			default:
				fmt.Fprintf(b, "Q%v%v%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]))
			}
		case SmoothQuadraticBezierTo:
			switch ls.(type) {
			case SmoothQuadraticBezierTo:
				fmt.Fprintf(b, "%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
			default:
				fmt.Fprintf(b, "T%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
			}
		case QuadraticBezierToRelative:
			switch ls.(type) {
			case QuadraticBezierToRelative:
				fmt.Fprintf(b, "%v%v%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]))
			default:
				fmt.Fprintf(b, "q%v%v%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]))
			}
		case SmoothQuadraticBezierToRelative:
			switch ls.(type) {
			case SmoothQuadraticBezierToRelative:
				fmt.Fprintf(b, "%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
			default:
				fmt.Fprintf(b, "t%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]))
			}
		case CubicBezierTo:
			switch ls.(type) {
			case CubicBezierTo:
				fmt.Fprintf(b, "%v%v%v%v%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]), addPrefixCommaWhenNecessery(ts[4]), addPrefixCommaWhenNecessery(ts[5]))
			default:
				fmt.Fprintf(b, "C%v%v%v%v%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]), addPrefixCommaWhenNecessery(ts[4]), addPrefixCommaWhenNecessery(ts[5]))
			}
		case SmoothCubicBezierTo:
			switch ls.(type) {
			case CubicBezierTo:
				fmt.Fprintf(b, "%v%v%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]))
			default:
				fmt.Fprintf(b, "S%v%v%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]))
			}
		case CubicBezierToRelative:
			switch ls.(type) {
			case CubicBezierToRelative:
				fmt.Fprintf(b, "%v%v%v%v%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]), addPrefixCommaWhenNecessery(ts[4]), addPrefixCommaWhenNecessery(ts[5]))
			default:
				fmt.Fprintf(b, "c%v%v%v%v%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]), addPrefixCommaWhenNecessery(ts[4]), addPrefixCommaWhenNecessery(ts[5]))
			}
		case SmoothCubicBezierToRelative:
			switch ls.(type) {
			case SmoothCubicBezierToRelative:
				fmt.Fprintf(b, "%v%v%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]))
			default:
				fmt.Fprintf(b, "s%v%v%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]))
			}
		case ArcTo:
			switch ls.(type) {
			case ArcTo:
				fmt.Fprintf(b, "%v%v%v%v%v%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]), addPrefixCommaWhenNecessery(ts[4]), addPrefixCommaWhenNecessery(ts[5]), addPrefixCommaWhenNecessery(ts[6]))
			default:
				fmt.Fprintf(b, "A%v%v%v%v%v%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]), addPrefixCommaWhenNecessery(ts[4]), addPrefixCommaWhenNecessery(ts[5]), addPrefixCommaWhenNecessery(ts[6]))
			}
		case ArcToRelative:
			switch ls.(type) {
			case ArcToRelative:
				fmt.Fprintf(b, "%v%v%v%v%v%v%v", addPrefixCommaWhenNecessery(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]), addPrefixCommaWhenNecessery(ts[4]), addPrefixCommaWhenNecessery(ts[5]), addPrefixCommaWhenNecessery(ts[6]))
			default:
				fmt.Fprintf(b, "a%v%v%v%v%v%v%v", noLeadingZerox(ts[0]), addPrefixCommaWhenNecessery(ts[1]), addPrefixCommaWhenNecessery(ts[2]), addPrefixCommaWhenNecessery(ts[3]), addPrefixCommaWhenNecessery(ts[4]), addPrefixCommaWhenNecessery(ts[5]), addPrefixCommaWhenNecessery(ts[6]))
			}
		}
		ls = s
	}
	return b.String()
}
