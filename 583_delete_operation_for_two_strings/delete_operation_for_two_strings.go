package main

import "github.com/AngryYogurt/leetcode/tools"

func main() {
	tools.AssertObjectEqual(2, minDistance("sea", "eat"))
	tools.AssertObjectEqual(0, minDistance("a", "a"))
}
/*
 * 执行用时 :8 ms, 在所有 Go 提交中击败了75.00%的用户
 * 内存消耗 :3.1 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func minDistance(word1 string, word2 string) int {
	lcs := lcs(word1, word2)
	return len(word1) - lcs + len(word2) - lcs
}

// 动态规划最长公共子序列
func lcs(word1 string, word2 string) int {
	if len(word1) < len(word2) {
		word1, word2 = word2, word1
	}
	a1 := make([]int, len(word1)+1)
	a2 := make([]int, len(word1)+1)
	for i := 1; i <= len(word2); i++ {
		a1, a2 = a2, a1
		for j := 1; j <= len(word1); j++ {
			if word1[j-1] == word2[i-1] {
				a2[j] = a1[j-1] + 1
			} else {
				a2[j] = max(a1[j], a2[j-1])
			}
		}
	}
	return a2[len(word1)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
