package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual([]int{}, findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
/*
 * 执行用时：24 ms, 在所有 Go 提交中击败了98.64%的用户
 * 内存消耗：6.3 MB, 在所有 Go 提交中击败了90.48%的用户
 */
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findDiagonalOrder(matrix [][]int) []int {

	if len(matrix) <= 0 {
		return []int{}
	}
	mi, mj := len(matrix), len(matrix[0])
	res := make([]int, mi*mj)
	i, j, idx := 0, 0, 0
	dirUp := true
	currLen := 1
	currRoad := 1
	for {
		res[idx] = matrix[i][j]
		currLen--
		if currLen == 0 {
			if currRoad >= mi+mj-1 {
				break
			}
			dirUp = !dirUp
			currRoad++
			currLen = currRoad
			if currLen > min(mi, mj) {
				if currLen > max(mi, mj) {
					currLen = mi + mj - currRoad
				} else {
					currLen = min(mi, mj)
				}
			}
			if (i == 0 && j != mj-1) || i == mi-1 {
				j++
			} else if j == 0 || j == mj-1 {
				i++
			}
		} else {
			if dirUp {
				i--
				j++
			} else {
				i++
				j--
			}
		}
		idx++
	}
	return res
}
