package main

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 0 || numRows == 1 {
		return s
	}

	n := len(s)
	delta := 2*numRows - 2
	var b bytes.Buffer

	for i := 0; i < numRows; i++ {

		for j := 0; j+i < n; j += delta {
			b.WriteByte(s[i+j])

			if i != 0 && i != numRows-1 && (j+delta-i < n) {
				b.WriteByte(s[j+delta-i])
			}
		}
	}

	return b.String()
}
