/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	res := 0
	for res < x {
		res = res*10 + x%10
		x /= 10
	}
	return res == x || res/10 == x
}

// @lc code=end

