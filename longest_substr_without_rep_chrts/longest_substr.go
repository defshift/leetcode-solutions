package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	if len(s) == 1 {
		return 1
	}

	hm := make(map[byte]int)
	hm[s[0]] = 0

	max := 1
	currentMax := max
	lastOccurIndex := 0

	for i := 1; i < len(s); i++ {
		is := s[i]

		repIndex, ok := hm[is]

		if ok {
			if currentMax > max {
				max = currentMax
			}

			currentMax = len(s[repIndex+1 : i+1])

			for j := lastOccurIndex; j < repIndex; j++ {
				delete(hm, s[j])
			}

			lastOccurIndex = repIndex + 1
		} else {
			currentMax++
		}

		hm[is] = i
	}

	if currentMax > max {
		max = currentMax
	}

	return max
}
