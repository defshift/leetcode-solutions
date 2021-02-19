package main

func maxArea(height []int) int {
	maxArea := 0
	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			min := height[i]
			if height[j] < min {
				min = height[j]
			}

			area := min * (j - i)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}
