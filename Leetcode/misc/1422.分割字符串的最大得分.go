/*
 * @lc app=leetcode.cn id=1422 lang=golang
 *
 * [1422] 分割字符串的最大得分
 */

// @lc code=start
func maxScore(s string) int {

	zero, one := 0, 0
	max := -1
	if s[len(s)-1] == '1' {
		one++
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zero++
		} else {
			zero--
			one++
		}
		if zero > max {
			max = zero
		}
	}
	return max + one
}

// @lc code=end

