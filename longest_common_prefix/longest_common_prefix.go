package main

import (
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	s0 := strs[0]

	for i := 0; i < len(s0); i++ {
		for j := 1; j < len(strs); j++ {
			if !strings.HasPrefix(strs[j], s0[0:i+1]) {
				return s0[0:i]
			}
		}
	}

	return s0
}
