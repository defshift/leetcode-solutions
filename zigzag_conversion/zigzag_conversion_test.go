package main

import (
	"testing"
)

func Test_convert(t *testing.T) {
	val := "Artkenrkgne23423fdwfs"
	exp := "Aeg3fsrknkn243dftre2w"
	act := convert(val, 3)

	if act != exp {
		t.Errorf("Expected: %v Actual: %v", exp, act)
	}

	val = "ABC"
	exp = "ABC"
	act = convert(val, 3)

	if act != exp {
		t.Errorf("Expected: %v Actual: %v", exp, act)
	}
}
