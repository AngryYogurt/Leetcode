package main

import "github.com/AngryYogurt/leetcode/tools"

func main() {
	tools.AssertObjectEqual(false, search([]int{1, 1}, 0))
}

/*
 * 执行用时 :4 ms, 在所有 Go 提交中击败了93.08%的用户
 * 内存消耗 :3.2 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return true
		}
		return false
	}
	rotIdx := findRotateIndex(nums)
	left, right, length := rotIdx, rotIdx-1, len(nums)
	oldL, oldR := convertToOldIndex(left, rotIdx, length), convertToOldIndex(right, rotIdx, length)
	for ; oldL <= oldR; oldL, oldR = convertToOldIndex(left, rotIdx, length), convertToOldIndex(right, rotIdx, length) {
		mid := convertToNewIndex(oldL+(oldR-oldL)/2, rotIdx, length)
		if nums[mid] == target {
			return true
		}
		if oldL == oldR {
			return false
		}
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid
		}
	}
	return false
}

func convertToOldIndex(newIdx, rotIdx, length int) int {
	if newIdx < rotIdx {
		return newIdx + length - rotIdx
	}
	return newIdx - rotIdx
}

func convertToNewIndex(old, rotIdx, length int) int {
	if old >= (length - rotIdx) {
		return old - (length - rotIdx)
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
		} else {
			start++
		}
	}
	return 0
}
