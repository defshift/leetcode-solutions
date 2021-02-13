package main

import "testing"

func assertStringResult(t *testing.T, s, exp string) {
	a := longestPalindrome(s)
	if a != exp {
		t.Errorf("Expected: %v Actual: %v", exp, a)
	}
}

func Test_longestPalindrome(t *testing.T) {
	assertStringResult(t, "babdrrggbbbabbb", "bbbabbb")
	assertStringResult(t, "rvwe  uuIuu mmmmpmmm", "mmmpmmm")
}
