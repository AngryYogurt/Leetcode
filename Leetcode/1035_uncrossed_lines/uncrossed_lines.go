package main

func main() {

}
/*
 * 执行用时：4 ms, 在所有 Go 提交中击败了80.95%的用户
 * 内存消耗：6.4 MB, 在所有 Go 提交中击败了6.90%的用户
 */
func maxUncrossedLines(A []int, B []int) int {
	if len(A) == 0 || len(B) == 0 {
		return 0
	}
	dp := make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}
	max := 0
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > max {
					max = dp[i][j]
				}
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return max
}
