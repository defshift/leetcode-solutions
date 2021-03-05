package integer_to_roman

import "testing"

func assert(t *testing.T, act, exp string) {
	if act != exp {
		t.Errorf("Actual: %v Expected: %v ", act, exp)
	}
}

func TestIntegerToRoman(t *testing.T) {
	assert(t, intToRoman(20), "XX")
	assert(t, intToRoman(1994), "MCMXCIV")
	assert(t, intToRoman(7), "VII")
	assert(t, intToRoman(1993), "MCMXCIII")
}
