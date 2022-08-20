/*
 * @lc app=leetcode.cn id=1413 lang=golang
 *
 * [1413] 逐步求和得到正数的最小值
 */

// @lc code=start
func minStartValue(nums []int) int {
	if len(nums) <= 0 {
		return 1
	}
	min := nums[0]
	sum := 0
	for _, v := range nums {
		sum += v
		if sum < min {
			min = sum
		}
	}
	if min >= 1 {
		return 1
	}
	return 1 - min
}

// @lc code=end

