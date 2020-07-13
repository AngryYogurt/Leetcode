package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

/*
 * 执行用时 :4 ms, 在所有 Go 提交中击败了90.07%的用户
 * 内存消耗 :2.8 MB, 在所有 Go 提交中击败了66.67%的用户
 */

// 两种方法只是循环逻辑稍有差异，2更简洁
func trap(height []int) int {
	res := 0
	if len(height) <= 2 {
		return 0
	}
	left, right := 0, len(height)-1
	maxL, maxR := 0, 0
	for left < right {
		for height[left] < height[right] {
			if height[left] <= maxL {
				res += maxL - height[left]
			} else {
				maxL = height[left]
			}
			left++
			if left >= right {
				return res
			}
		}
		for height[left] >= height[right] {
			if height[right] <= maxR {
				res += maxR - height[right]
			} else {
				maxR = height[right]
			}
			right--

			if left >= right {
				return res
			}
		}
	}
	return res
}

func trap2(height []int) int {
	res := 0
	if len(height) <= 2 {
		return 0
	}
	left, right := 0, len(height)-1
	maxL, maxR := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] <= maxL {
				res += maxL - height[left]
			} else {
				maxL = height[left]
			}
			left++
		} else {
			if height[right] <= maxR {
				res += maxR - height[right]
			} else {
				maxR = height[right]
			}
			right--
		}
	}
	return res
}
