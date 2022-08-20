/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) <= 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] != val {
			left++
			continue
		}
		if nums[right] == val {
			right--
			continue
		}
		nums[left] = nums[right]
		nums[right] = val
		right--
		left++
	}
	if nums[left] == val {
		left--
	}
	return left + 1
}

// @lc code=end

