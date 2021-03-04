package reverse_integer

import (
	"math"
)

const threshold = math.MaxInt32 / 10

func reverse(x int) int {
	res := 0
	rem := x

	for rem != 0 {
		digit := rem % 10
		rem = rem / 10

		if (math.Abs(float64(res)) > threshold) || ((res == threshold) && math.Abs(float64(digit)) > 7) {
			return 0
		}

		res = res*10 + digit
	}

	return res
}
