package frac

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var Pow10 = [...]int64{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
	10000000000,
	100000000000,
	1000000000000,
	10000000000000,
	100000000000000,
	1000000000000000,
	10000000000000000,
	100000000000000000,
	1000000000000000000,
}

type MagicFloat struct {
	Num int64
	Den int
}

func FloatToMagic(num interface{}, prec int) (MagicFloat, error) {

	switch n := num.(type) {
	case float32:
		var out MagicFloat
		out.Den = prec
		out.Num = int64(n * float32(Pow10[prec]))
		return out, nil
	case float64:
		var out MagicFloat
		out.Den = prec
		out.Num = int64(n * float64(Pow10[prec]))
		return out, nil
	default:
		return MagicFloat{}, errors.New("Not A Float")
	}
}

func ToMagic(num interface{}) (MagicFloat, error) {
	switch n := num.(type) {
	case int:
		return MagicFloat{Num: int64(n), Den: 1}, nil
	case int16:
		return MagicFloat{Num: int64(n), Den: 1}, nil
	case int32:
		return MagicFloat{Num: int64(n), Den: 1}, nil
	case int64:
		return MagicFloat{Num: int64(n), Den: 1}, nil
	case uint:
		return MagicFloat{Num: int64(n), Den: 1}, nil
	case uint16:
		return MagicFloat{Num: int64(n), Den: 1}, nil
	case uint32:
		return MagicFloat{Num: int64(n), Den: 1}, nil
	case uint64:
		return MagicFloat{Num: int64(n), Den: 1}, nil

	case string:
		// This function uses the "." in the string to find the num and den
		if strings.ContainsRune(n, '.') {
			part1 := n[:strings.IndexByte(n, '.')]
			part2 := n[(strings.IndexByte(n, '.') + 1):]
			Den := len(part2)
			numstring := part1 + part2
			Num, err := strconv.ParseInt(numstring, 10, 64)

			out := MagicFloat{
				Num: Num,
				Den: int(Den),
			}
			return out, err
		} else {

			Num, err := strconv.ParseInt(n, 10, 64)
			out := MagicFloat{
				Num: Num,
				Den: 1,
			}
			return out, err
		}
	default:
		return MagicFloat{}, errors.New("Not An Int or String")
	}
}

func (m MagicFloat) Format(f fmt.State, c rune) {
	bottom := float64(Pow10[m.Den])
	top := m.Num

	out := float64(top) / bottom

	strout := strconv.FormatFloat(out, 'f', -1, 64)

	f.Write([]byte(strout))
}

func ToFloat(a MagicFloat) float64 {
	return float64(a.Num) / float64(Pow10[a.Den])

}

//Add will retain the higher precision number
func Add(a MagicFloat, b MagicFloat) MagicFloat {
	if b.Den > a.Den {
		precDiff := b.Den - a.Den
		a.Num = a.Num * Pow10[precDiff]
		return MagicFloat{
			Num: a.Num + b.Num,
			Den: b.Den,
		}
	} else {
		precDiff := a.Den - b.Den
		b.Num = b.Num * Pow10[precDiff]
		return MagicFloat{
			Num: a.Num + b.Num,
			Den: a.Den,
		}
	}
}
func Sub(a MagicFloat, b MagicFloat) MagicFloat {
	b.Neg()
	if b.Den > a.Den {
		precDiff := b.Den - a.Den
		a.Num = a.Num * Pow10[precDiff]
		return MagicFloat{
			Num: a.Num + b.Num,
			Den: b.Den,
		}
	} else {
		precDiff := a.Den - b.Den
		b.Num = b.Num * Pow10[precDiff]
		return MagicFloat{
			Num: a.Num + b.Num,
			Den: a.Den,
		}
	}
}

//Mult will retain the higher precision number
func Mult(a MagicFloat, b MagicFloat) MagicFloat {
	if b.Den > a.Den {
		precDiff := b.Den - a.Den
		a.Num = a.Num * Pow10[precDiff]
		return MagicFloat{
			Num: (a.Num * b.Num) / Pow10[b.Den],
			Den: b.Den,
		}
	} else {
		precDiff := a.Den - b.Den
		b.Num = b.Num * Pow10[precDiff]
		return MagicFloat{
			Num: (a.Num * b.Num) / Pow10[a.Den],
			Den: a.Den,
		}
	}
}

//Mult will retain the higher precision number
func Div(a MagicFloat, b MagicFloat) MagicFloat {
	if b.Den > a.Den {
		precDiff := b.Den - a.Den
		a.Num = a.Num * Pow10[precDiff]
		return MagicFloat{
			Num: ((a.Num * Pow10[b.Den]) / (b.Num)),
			Den: b.Den,
		}
	} else {
		precDiff := a.Den - b.Den
		b.Num = b.Num * Pow10[precDiff]
		return MagicFloat{
			Num: ((a.Num * Pow10[a.Den]) / (b.Num)),
			Den: a.Den,
		}
	}
}

func (m *MagicFloat) Neg() {
	m.Num = -m.Num
}

func (m *MagicFloat) Abs() {
	if m.Num < 0 {
		m.Num = -m.Num
	}
}

func (m *MagicFloat) SetPrec(n int) {
	if n > m.Den {
		m.Den = n
		m.Num = m.Num * Pow10[n-m.Den]
	} else {
		m.Den = n
		m.Num = m.Num / Pow10[m.Den-n]
	}
}
