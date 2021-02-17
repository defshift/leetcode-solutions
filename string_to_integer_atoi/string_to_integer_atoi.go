package main

import (
	"math"
)

const threshold = math.MaxInt32 / 10

func myAtoi(s string) int {
	reachedNum := false
	res := 0
	sign := 1

	for i := 0; i < len(s); i++ {
		isNumber := s[i] >= '0' && s[i] <= '9'
		isSign := s[i] == '-' || s[i] == '+'
		isSpace := s[i] == ' '

		if !reachedNum && !(isNumber || isSign || isSpace) {
			break
		}

		if !reachedNum && (isNumber || isSign) {
			reachedNum = true

			if isSign {
				if s[i] == '-' {
					sign = -1
				}
				continue
			}
		}

		if reachedNum && isNumber {
			if (res > threshold) || ((res == threshold) && (int(s[i]-'0') > 7)) {
				if sign == 1 {
					res = math.MaxInt32
				} else {
					res = math.MinInt32
				}
				return res
			}
			res = res*10 + int(s[i]-'0')
		}

		if reachedNum && !isNumber {
			break
		}
	}

	return res * sign
}
