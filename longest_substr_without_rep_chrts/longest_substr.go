package main

func lengthOfLongestSubstring(s string) int {
	hm := [128]int{}

	max := 0
	curMax := 0
	for i, j := 0, 0; j < len(s); j++ {
		p := hm[s[j]]

		if p != 0 {
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
