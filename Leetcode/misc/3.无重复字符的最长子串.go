/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}
	start := 0
	set := make([]int, 2048)
	for i := range set {
		set[i] = -1
	}
	res := 1
	for idx, v := range []byte(s) {
		if last := set[v-'a']; last >= start {
			if now := idx - start; now > res {
				res = now
			}
			start = last + 1
		}
		set[v-'a'] = idx
	}
	if now := len(s) - start; now > res {
		res = now
	}
	return res
}

// @lc code=end

