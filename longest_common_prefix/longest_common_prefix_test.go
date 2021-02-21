package main

import "testing"

func assert(t *testing.T, act, exp string) {
	if act != exp {
		t.Errorf("Actual: %v Expected: %v ", act, exp)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	assert(t, longestCommonPrefix([]string{"ab", "a"}), "a")
}
