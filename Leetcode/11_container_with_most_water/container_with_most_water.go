package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(49, maxArea([]int{1,8,6,2,5,4,8,3,7}))
}

/*
 * 执行用时：16 ms, 在所有 Go 提交中击败了79.49%的用户
 * 内存消耗：5.8 MB, 在所有 Go 提交中击败了16.67%的用户
 */
func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		res = max(res, min(height[l], height[r])*(r-l))
		if height[r] < height[l] {
			r--
		} else {
			l++
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
