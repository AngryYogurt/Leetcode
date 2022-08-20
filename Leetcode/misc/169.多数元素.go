/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	res := nums[0]
	ct := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == res {
			ct++
		} else {
			ct--
			if ct < 0 {
				res = nums[i]
				ct = 1
			}
		}
	}
	return res
}

// @lc code=end

