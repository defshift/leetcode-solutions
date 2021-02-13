package main

func longestPFromCenter(s string, left, right int) int {
	l := left
	r := right

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l + 1
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start := 0
	end := 0
	for i := 0; i < len(s); i++ {
		l1 := longestPFromCenter(s, i, i)
		l2 := longestPFromCenter(s, i, i+1)

		l := l1
		if l2 > l1 {
			l = l2
		}

		if l > (end - start) {
			start = i - (l-1)/2
			end = i + l/2
		}
	}

	return s[start+1 : end]
}
