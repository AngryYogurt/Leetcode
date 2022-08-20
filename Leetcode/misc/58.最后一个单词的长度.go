/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	res := 0
	cur := 0
	for _, v := range s {
		if v == ' ' {
			if cur > 0 {
				res = cur
			}
			cur = 0
		} else {
			cur++
		}
	}
	if cur > 0 {
		return cur
	}
	return res
}

// @lc code=end

