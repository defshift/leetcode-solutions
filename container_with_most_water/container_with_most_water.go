package main

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1

	maxArea := 0
	for i != j {
		min := height[i]
		if height[j] < min {
			min = height[j]
		}

		area := (j - i) * min

		if area > maxArea {
			maxArea = area
		}

		if height[j] > height[i] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}
