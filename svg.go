package patterns

import "io"
import "strings"
import "errors"

//import "fmt"

const bufLen = 100


type errParsingReader struct {
	error
	Source      io.Reader
}

func (e errParsingReader) Parsing() io.Reader {
	return e.Source
}


func Polygon(p Brush, cs string) (Pattern, error) {
	reader := strings.NewReader(cs)
	fReader := float32ListReader{Reader: reader, buf: make([]byte, bufLen)}
	coordsBuf := make([]float32, 10)
	c, err := fReader.ReadFloat32s(coordsBuf)
	if err != io.EOF {
		return nil,errParsingReader{err,fReader}
	}
	if c%2 != 0 {
		return nil,errParsingReader{errors.New("Coordinate count."),reader}
	}
	coords := make([][2]x, c/2)
	for n := 0; n < c; n += 2 {
		coords[n/2]=[2]x{X(coordsBuf[n]),X(coordsBuf[n+1])}
	}
	return p.Polygon(coords), nil
}

type StageCompleted uint8

const (
	None StageCompleted = iota
	wholeSign
	whole
	Fraction
	ExponentSign
	Exponent
)

type float32ListReader struct {
	io.Reader
	charFound             bool  // something has been found in current section
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
	unBuf                 []byte // slice into buf, pointing to the unconsumed bytes remaining after last read stopped
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
	var value func() float32
	value = func() float32 {
		return (float32(this.partial) + float32(this.partialFraction)/power10(uint(this.partialFractionDigits))) 
	}
	// put value into array
	var setVal func()
	setVal = func() {
		if this.neg {
			if this.negExponent {
				fs[c] = -value()/ power10(this.partialExponent)
			} else {
				fs[c] = -value()* power10(this.partialExponent)
			}
		} else {
			if this.negExponent {
				fs[c] = value()/ power10(this.partialExponent)
			} else {
				fs[c] = value()* power10(this.partialExponent)
			}
		}
		//fmt.Println(c,":",fs[c])
		c++
	}
	var n int
	var b []byte
	for err == nil {
		if len(this.unBuf) != 0 { // use any unread first
			n = len(this.unBuf)
			b = this.unBuf
		} else {
			n, err = this.Reader.Read(this.buf)
			b = this.buf
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
					return c, errParsingReader{errors.New("Can't parse as number."),this.Reader}
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
						return c, errParsingReader{errors.New("Separator without value."),this.Reader}
					}
				}
				this.charFound = false
			case ' ', '\n', '\r', '\t': // accept these as separators, but multiple occurances only count once. // two tabs only one separator in an svg path
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
					return c, errParsingReader{errors.New("Can't parse as numbers."),this.Reader}

				}
				if this.wholeFound && !this.fractionFound {
					return c, errParsingReader{errors.New("Can't parse as numbers."),this.Reader}

				}
			case '+':
				this.inSeparator = false
				if this.charFound || this.wholeFound && !this.fractionFound {
					return c, errParsingReader{errors.New("Can't parse as numbers."),this.Reader}

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
				return c, errParsingReader{errors.New("Uninterpretable character found."),this.Reader}
			}
		}
		this.unBuf = this.unBuf[0:0]
	}
	if this.wholeFound || this.charFound {
		setVal()
	}
	return c, err
}



