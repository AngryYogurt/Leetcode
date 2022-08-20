/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	digits[len(digits)-1]++
	if digits[len(digits)-1] < 10 {
		return digits
	}
	digits[len(digits)-1] = 0
	carry := 1
	for i := len(digits) - 2; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			carry = 0
			break
		}
		digits[i] = 0
	}
	if carry > 0 {
		return append([]int{1}, digits...)
	}
	return digits
}

// @lc code=end

