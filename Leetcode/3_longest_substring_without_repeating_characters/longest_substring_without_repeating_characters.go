package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(1, lengthOfLongestSubstring("bbbbb"))
	tools.AssertObjectEqual(3, lengthOfLongestSubstring("pwwkew"))
	tools.AssertObjectEqual(3, lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	charPos := map[byte]int{}

	start, end := 0, 0
	max := 0
	for ; start < len(s); start++ {
		if start > 0 {
			charPos[s[start-1]] = 0
		}
		for end < len(s) && charPos[s[end]] <= 0 {
			if max < end-start+1 {
				max = end - start + 1
			}
			charPos[s[end]]++
			end++
		}
	}
	return max
}
