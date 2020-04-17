package main

import "github.com/AngryYogurt/leetcode/tools"

func main() {
	tools.AssertObjectEqual("bab", longestPalindrome("babad"))
}

/*
 * 执行用时 :4 ms, 在所有 Go 提交中击败了92.72%的用户
 * 内存消耗 :2.2 MB, 在所有 Go 提交中击败了44.83%的用户
 */
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	maxLen, start, end := 0, 0, 0
	for i := 0; i <= len(s)*2-1; i++ {
		pre, post := i/2, i/2+1
		if i%2 == 0 {
			pre, post = i/2-1, i/2+1
		}
		for pre >= 0 && post < len(s) && s[pre] == s[post] {
			pre--
			post++
		}
		if currLen := post - pre - 1; currLen > maxLen {
			maxLen = currLen
			start = pre + 1
			end = post - 1
		}
	}
	return s[start : end+1]
}
