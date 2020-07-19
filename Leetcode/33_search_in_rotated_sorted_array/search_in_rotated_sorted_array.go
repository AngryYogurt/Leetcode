package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(1, search([]int{1, 3}, 3))
	tools.AssertObjectEqual(-1, search([]int{3, 5, 1}, 2))
	tools.AssertObjectEqual(-1, search([]int{1, 3}, 0))
	tools.AssertObjectEqual(4, search([]int{4, 5, 6, 7, 0, 1, 2, 3}, 0))
	tools.AssertObjectEqual(5, search([]int{4, 5, 6, 7, 8, 0, 1, 2, 3}, 0))
}
/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.6 MB, 在所有 Go 提交中击败了71.43%的用户
 */
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	rotIdx := findRotateIndex(nums)
	left, right, length := rotIdx, rotIdx-1, len(nums)
	oldL, oldR := convertToOldIndex(left, rotIdx, length), convertToOldIndex(right, rotIdx, length)
	for ; oldL <= oldR; oldL, oldR = convertToOldIndex(left, rotIdx, length), convertToOldIndex(right, rotIdx, length) {
		mid := convertToNewIndex(oldL+(oldR-oldL)/2, rotIdx, length)
		if nums[mid] == target {
			return mid
		}
		if oldL == oldR {
			return -1
		}
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return -1
}

func convertToOldIndex(newIdx, rotIdx, length int) int {
	if newIdx < rotIdx {
		return newIdx + length - rotIdx
	}
	return newIdx - rotIdx
}

func convertToNewIndex(old, rotIdx, length int) int {
	if old >= (length-rotIdx) {
		return old - (length-rotIdx)
	}
	return old + rotIdx
}

func findRotateIndex(nums []int) int {
	start, end := 0, len(nums)-1
	if nums[end] > nums[start] {
		return 0
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		} else if nums[mid] < nums[start] {
			end = mid
		} else if nums[mid+1] > nums[end] {
			start = mid + 1
		}
	}
	return 0
}
