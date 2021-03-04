package reverse_integer

import (
	"testing"
)

func assert(t *testing.T, act, exp int) {
	if act != exp {
		t.Errorf("Actual: %v Expected: %v ", act, exp)
	}
}

func TestReverse(t *testing.T) {
	assert(t, reverse(-123456), -654321)
	assert(t, reverse(-100000000009), 0)
}
