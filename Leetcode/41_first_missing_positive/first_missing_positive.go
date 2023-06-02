package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(3, firstMissingPositive([]int{1, 2, 0}))
	tools.AssertObjectEqual(2, firstMissingPositive([]int{3, 4, -1, 1}))
	tools.AssertObjectEqual(1, firstMissingPositive([]int{7, 8, 9, 10, 11, 12}))
}

func firstMissingPositive(nums []int) int {
	max := len(nums)
	for i := 0; i < max; i++ {
		if nums[i] <= 0 {
			nums[i] = max + 1
		}
	}

	for idx := 0; idx < max; idx++ {
		n := abs(nums[idx])
		if n > max {
			continue
		}
		nums[n-1] = -abs(nums[n-1])
	}

	for i := 0; i < max; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return max + 1
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
