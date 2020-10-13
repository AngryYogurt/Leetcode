package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a*"))
}
/*
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：2.4 MB, 在所有 Go 提交中击败了36.33%的用户
 */
func isMatch(s string, p string) bool {
	if s == p && s == "" {
		return true
	}
	if len(p) <= 0 || p[0] == '*' {
		return false
	}
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if i == 0 && p[j-1] != '*'{
				continue
			} else {
				if p[j-1] == '*' {
					dp[i][j] = dp[i][j] || dp[i][j-2]

					if i > 0 && match(s[i-1], p[j-2]) {
						dp[i][j] = dp[i-1][j] || dp[i][j]
					}
				} else if i > 0 && match(s[i-1], p[j-1]) {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

func match(a, b byte) bool {
	if b == '.' || a == b {
		return true
	}
	return false
}
