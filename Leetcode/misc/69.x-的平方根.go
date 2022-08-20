/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	left, right := 1, x
	for left < right-1 {
		mid := (left + right) >> 1
		tmp := mid * mid
		if tmp > x {
			right = mid
		} else if tmp < x {
			left = mid
		} else {
			return mid
		}
	}
	return left
}

// @lc code=end

