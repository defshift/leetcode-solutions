package main

func twoSum(nums []int, target int) []int {
	for first := 0; first < len(nums)-1; first++ {
		for second := first + 1; second < len(nums); second++ {
			if res := nums[first] + nums[second]; res == target {
				return []int{first, second}
			}
		}
	}

	return nil
}
