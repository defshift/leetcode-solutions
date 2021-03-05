package integer_to_roman

import (
	"bytes"
)

const (
	I  = 1
	IV = 4
	V  = 5
	IX = 9
	X  = 10
	XL = 40
	L  = 50
	XC = 90
	C  = 100
	CD = 400
	D  = 500
	CM = 900
	M  = 1000
)

type pair struct {
	decimalVal int
	romanVal   string
}

var mapping = []pair{
	{I, "I"},
	{IV, "IV"},
	{V, "V"},
	{IX, "IX"},
	{X, "X"},
	{XL, "XL"},
	{L, "L"},
	{XC, "XC"},
	{C, "C"},
	{CD, "CD"},
	{D, "D"},
	{CM, "CM"},
	{M, "M"},
}

func intToRoman(num int) string {
	rem := num
	i := len(mapping)

	res := bytes.Buffer{}

	for i != 0 {
		div := mapping[i-1].decimalVal
		rs := rem / div
		rem = rem % div

		if rs > 0 {
			rem += div * (rs - 1)
		}

		if rs == 0 {
			i--
		} else {
			res.WriteString(mapping[i-1].romanVal)
		}
	}

	return res.String()
}
