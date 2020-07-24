package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(7, minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

/*
 * 执行用时：8 ms, 在所有 Go 提交中击败了89.86%的用户
 * 内存消耗：3.9 MB, 在所有 Go 提交中击败了83.33%的用户
 */
func minPathSum(grid [][]int) int {
	if len(grid) <= 0 {
		return 0
	}
	tmp := make([]int, len(grid[0]))
	tmp[0] = grid[0][0]
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				tmp[j] = tmp[j-1] + grid[i][j]
			} else if j == 0 {
				tmp[j] = tmp[j] + grid[i][j]
			} else {
				tmp[j] = min(tmp[j-1], tmp[j]) + grid[i][j]
			}
		}
	}
	return tmp[len(tmp)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
