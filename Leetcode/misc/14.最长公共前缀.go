/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	idx := 0
	for idx < len(strs[0]) {
		ch := strs[0][idx]
		for _, str := range strs {
			if len(str) <= idx || str[idx] != ch {
				return str[:idx]
			}
		}
		idx++
	}
	return strs[0]
}

// @lc code=end

