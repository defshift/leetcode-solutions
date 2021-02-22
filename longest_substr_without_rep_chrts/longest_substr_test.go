package main

import "testing"

func assert(t *testing.T, act, exp int) {
	if act != exp {
		t.Errorf("Expected: %v Actual: %v", exp, act)
	}
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert(t, lengthOfLongestSubstring("bba"), 2)
	assert(t, lengthOfLongestSubstring("abcb"), 3)
	assert(t, lengthOfLongestSubstring("abcabcbb"), 3)
	assert(t, lengthOfLongestSubstring("er er ewrwer34 34444 "), 6)
	assert(t, lengthOfLongestSubstring("dvdf"), 3)
	assert(t, lengthOfLongestSubstring("abs"), 3)
	assert(t, lengthOfLongestSubstring("abss"), 3)
	assert(t, lengthOfLongestSubstring(""), 0)
	assert(t, lengthOfLongestSubstring("abs absdf"), 6)
	assert(t, lengthOfLongestSubstring("aaaaa"), 1)
	assert(t, lengthOfLongestSubstring("aaaaaaaabca"), 3)
	assert(t, lengthOfLongestSubstring("abcdawebom"), 8)
}
