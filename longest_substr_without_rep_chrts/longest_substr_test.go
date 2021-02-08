package main

import "testing"

func assertStringResult(t *testing.T, s string, expLen int) {
	l := lengthOfLongestSubstring(s)
	if l != expLen {
		t.Errorf("String '%v' should have longest substring with length %d, actual %d", s, expLen, l)
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	assertStringResult(t, "bba", 2)
	assertStringResult(t, "abcb", 3)
	assertStringResult(t, "abcabcbb", 3)
	assertStringResult(t, "er er ewrwer34 34444 ", 6)
	assertStringResult(t, "dvdf", 3)
	assertStringResult(t, "abs", 3)
	assertStringResult(t, "abss", 3)
	assertStringResult(t, "", 0)
	assertStringResult(t, "abs absdf", 6)
	assertStringResult(t, "aaaaa", 1)
	assertStringResult(t, "aaaaaaaabca", 3)
	assertStringResult(t, "abcdawebom", 8)
}
