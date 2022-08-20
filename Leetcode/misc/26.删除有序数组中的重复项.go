/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	lastVal, lastIdx := nums[0], 0
	for _, v := range nums {
		if v == lastVal {
			continue
		}
		lastVal = v
		lastIdx++
		nums[lastIdx] = v
	}
	return lastIdx + 1
}

// @lc code=end

