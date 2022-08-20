/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	k, r := x, 1
	if k < 0 {
		k = -k
		r = -1
	}
	res := 0
	for k > 0 {
		res *= 10
		res += k % 10
		k = k / 10
		if (r > 0 && res >= (1<<31)-1) || (r < 0 && res >= 1<<31) {
			return 0
		}
	}
	return res * r
}

// @lc code=end

