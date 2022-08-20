/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}
	return search(nums, target)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right + left - 1) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}

// @lc code=end

