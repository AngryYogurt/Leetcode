package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(-1, minimumTotal([][]int{
		{-1},
		{3, 2},
		{-3, 1, -1},
	}))
	tools.AssertObjectEqual(11, minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}
/*
 * 执行用时：4 ms, 在所有 Go 提交中击败了95.08%的用户
 * 内存消耗：3.1 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	tmp := make([]int, len(triangle)+1)
	for _, v := range triangle {
		tmp[len(v)] = tmp[len(v)-1] + v[len(v)-1]
		for i := len(v) - 2; i >= 0; i-- {
			tmp[i+1] = v[i+1] + min(tmp[i+1], tmp[i])
		}
		tmp[0] += v[0]
	}
	res := tmp[0]
	for i := 1; i < len(tmp)-1; i++ {
		if tmp[i] < res {
			res = tmp[i]
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
