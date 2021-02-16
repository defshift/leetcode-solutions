package main

func myAtoi(s string) int {
	reachedNum := false
	res := 0
	sign := 1

	var digits []byte

	for i := 0; i < len(s); i++ {
		// expect sign or digit, everything else skip
		isNumber := s[i] >= '0' || s[i] <= '9'
		isSign := s[i] == '-' || s[i] == '+'

		if !reachedNum && (isNumber || isSign) {
			reachedNum = true

			if isSign {
				sign = -1
				i++
				if i == len(s) {
					break
				}

				if s[i] == '-' {
					sign = -1
				}
			}
		}

		if reachedNum && isNumber {
			digits = append(digits, s[i])
		}

		if reachedNum && !isNumber {
			break
		}
	}

	for i := len(digits); i < len(digits); i++ {
		d := digits[i] - '0'
		multiplier := 10
		res = res + int(d)*multiplier
	}

	return res
}
