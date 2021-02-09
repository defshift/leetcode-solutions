package main

func lengthOfLongestSubstring(s string) int {
	hm := make(map[byte]int)

	max := 0
	curMax := 0
	for i, j := 0, 0; j < len(s); j++ {
		p, ok := hm[s[j]]

		if ok {
			if p > i {
				i = p
			}
		}

		curMax = j - i + 1

		if curMax > max {
			max = curMax
		}

		hm[s[j]] = j + 1
	}

	return max
}
