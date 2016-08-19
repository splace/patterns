package patterns

import "io"
import "strings"
import "errors"

//import "fmt"

const bufLen = 1


type errParse struct {
	error
	Source      io.Reader
}

func (e errParse) Parsing() io.Reader {
	return e.Source
}

var ErrPolyFormat = errors.New("polygon format, odd coordinate count.")

func Polygon(p Brush, cs string) (Pattern, error) {
	sReader := strings.NewReader(cs)
	fReader := float32ListReader{Reader: sReader, buf: make([]byte, bufLen)}
	coordsBuf := make([]float32, 10)
	c, err := fReader.ReadFloat32s(coordsBuf)
	if err != nil {
		return nil,errParse{err,fReader}
	}
	if c%2 != 0 {
		return nil,errParse{errors.New("Coordinate count."),fReader}
	}
	coords := make([][2]x, c/2)
	for n := 0; n < c; n += 2 {
		coords[n/2]=[2]x{X(coordsBuf[n]),X(coordsBuf[n+1])}
	}
	return p.Polygon(coords), nil
}

var ErrNumParse = errors.New("Can't parse as number.")
var ErrMissingItem = errors.New("Separator without value.")
var ErrUnknownChar = errors.New("Uninterpretable character found.")

type float32ListReader struct {
	io.Reader
	charFound             bool  // something found in current section
	inSeparator           bool  // currently parsing separator
	neg                   bool  // nagative number
	partial               uint  // whole number section, read so far
	wholeFound            bool  // whole number section complete
	partialFraction       uint  // fraction section read so far
	partialFractionDigits uint8 // count of fractional section digits, used to turn partialFraction in required real number by power of ten division
	fractionFound         bool  // fractional section complete
	partialExponent       uint  // exponent section so far read
	negExponent           bool
	buf                   []byte
	unBuf                 []byte // slice into buf, pointeing to the unconsumed bytes remaining after last read stopped
}

// reads text into a float array
func (this *float32ListReader) ReadFloat32s(fs []float32) (c int, err error) {
	var power10 func(uint) float32
	power10 = func(n uint) float32 {
		switch n {
		case 0:
			return 1
		case 1:
			return 1e1
		case 2:
			return 1e2
		case 3:
			return 1e3
		case 4:
			return 1e4
		case 5:
			return 1e5
		case 6:
			return 1e6
		case 7:
			return 1e7
		case 8:
			return 1e8
		case 9:
			return 1e9
		default:
			return 1e10 * power10(n-10)
		}
	}
	var val func() float32
	// assemble parts into complete value
	val = func() float32 {
		return (float32(this.partial) + float32(this.partialFraction)/power10(uint(this.partialFractionDigits))) * power10(this.partialExponent)
	}
	var value func() float32
	value = func() float32 {
		return (float32(this.partial) + float32(this.partialFraction)/power10(uint(this.partialFractionDigits))) / power10(this.partialExponent)
	}
	// put value into array
	var setVal func()
	setVal = func() {
		if this.neg {
			if this.negExponent {
				fs[c] = -value()
			} else {
				fs[c] = -val()
			}
		} else {
			if this.negExponent {
				fs[c] = value()
			} else {
				fs[c] = val()
			}
		}
		//fmt.Println(c,":",fs[c])
		c++
	}
	var n int
	var b []byte
	for err == nil {
		if this.unBuf == nil || len(this.unBuf) == 0 { // use any unread first
			n, err = this.Reader.Read(this.buf)
			b = this.buf
		} else {
			n = len(this.unBuf)
			b = this.unBuf
		}
		if n == 0 {
			break
		}
		for i := 0; i < n; i++ {
			switch b[i] {
			case '0':
				this.inSeparator = false
				if this.wholeFound {
					if this.fractionFound {
						if this.charFound {
							this.partialExponent *= 10
						}
					} else {
						this.partialFraction *= 10
						this.partialFractionDigits++
					}
				} else {
					if this.charFound {
						this.partial *= 10
					}
				}
				this.charFound = true
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				this.inSeparator = false
				if this.wholeFound {
					if this.fractionFound {
						this.partialExponent *= 10
						this.partialExponent += uint(b[i]) - 48
					} else {
						this.partialFraction *= 10
						this.partialFractionDigits++
						this.partialFraction += uint(b[i]) - 48
					}
				} else {
					this.partial *= 10
					this.partial += uint(b[i]) - 48
				}
				this.charFound = true
			case '.':
				this.inSeparator = false
				this.wholeFound = true
				this.charFound = false
			case 'e', 'E':
				if !this.charFound {
					return c, ErrNumParse
				} // cant have e wihout something before
				this.wholeFound = true
				this.fractionFound = true
				this.charFound = false
			case ',': // separators
				this.inSeparator = true
				if this.wholeFound || this.charFound {
					setVal()
					this.partial = 0
					this.partialFraction = 0
					this.partialFractionDigits = 0
					this.wholeFound = false
					this.partialExponent = 0
					this.fractionFound = false
					this.neg = false
					this.negExponent = false
					if c >= len(fs) {
						return c, err
					}
				} else {
					if !this.inSeparator {
						return c, ErrMissingItem
					}
				}
				this.charFound = false
			case ' ', '\n', '\r', '\t': // accept these as separators, but multiple occurances onyl count once. // two tabs only one separator in an svg path
				if !this.inSeparator {
					this.inSeparator = true
					if this.wholeFound || this.charFound {
						setVal()
						this.partial = 0
						this.partialFraction = 0
						this.partialFractionDigits = 0
						this.charFound = false
						this.wholeFound = false
						this.partialExponent = 0
						this.fractionFound = false
						this.neg = false
						this.negExponent = false
						if c >= len(fs) {
							return c, err
						}
					}
				}
			case '-':
				this.inSeparator = false
				if !this.charFound {
					if !this.wholeFound {
						this.neg = true
					} else {
						if this.fractionFound {
							this.negExponent = true
						}
					}
				} else {
					return c, ErrNumParse
				}
				if this.wholeFound && !this.fractionFound {
					return c, ErrNumParse
				}
			case '+':
				this.inSeparator = false
				if this.charFound || this.wholeFound && !this.fractionFound {
					return c, ErrNumParse
				}

			default:
				this.unBuf = b[i:n]
				if this.wholeFound || this.charFound {
					setVal()
					this.partial = 0
					this.partialFraction = 0
					this.partialFractionDigits = 0
					this.charFound = false
					this.wholeFound = false
					this.partialExponent = 0
					this.fractionFound = false
					this.inSeparator = false
					this.neg = false
					this.negExponent = false
					if c == len(fs) {
						return c, err
					}
				}
				return c, ErrUnknownChar
			}
		}
		this.unBuf = this.unBuf[0:0]
	}
	if this.wholeFound || this.charFound {
		setVal()
	}
	return c, err
}



