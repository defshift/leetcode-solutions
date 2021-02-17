package main

import "testing"

func assert(t *testing.T, act, exp int) {
	if act != exp {
		t.Errorf("Actual: %v Expected: %v ", act, exp)
	}
}

func Test_myAtoi(t *testing.T) {
	assert(t, myAtoi("-233"), -233)
	assert(t, myAtoi(" 123333"), 123333)
	assert(t, myAtoi("-32424534545454554"), -2147483648)
	assert(t, myAtoi("2rt-233"), 2)
	assert(t, myAtoi("2147483646"), 2147483646)
}
