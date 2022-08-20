/*
 * @lc app=leetcode.cn id=342 lang=golang
 *
 * [342] 4的幂
 */

// @lc code=start
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	idx := 0
	for {
		if n&1 == 1 {
			if n>>1 > 0 {
				return false
			}
			return idx%2 == 0
		}
		idx++
		n >>= 1
	}
}

// @lc code=end

