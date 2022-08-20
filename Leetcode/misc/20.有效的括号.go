/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	if len(s)&1 > 0 {
		return false
	}
	pair := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]byte, len(s))
	idx := 0
	for _, v := range []byte(s) {
		if p := pair[v]; p > 0 {
			if idx > 0 && stack[idx-1] == p {
				idx--
			} else {
				return false
			}
		} else {
			stack[idx] = v
			idx++
		}
	}
	return idx == 0
}

// @lc code=end

