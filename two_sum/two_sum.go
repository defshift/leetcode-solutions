package main

func twoSum2rounds(nums []int, target int) []int {
	valueMap := map[int]int{}

	for i, v := range nums {
		valueMap[v] = i
	}


	for i, v := range nums {
		complement := target - v
	
		if c, ok := valueMap[complement]; ok && i != c {
			return []int{i, c}
		}
	}

	return []int{0, 0} 
}

func twoSum1Round(nums []int, target int) []int {
	valueMap := map[int]int{}

	for i, v := range nums {
		if index, ok := valueMap[v]; ok {
			return []int{i, index}
		} else {
			valueMap[target - v] = i
		}
	}

	return []int{0, 0} 
}

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
