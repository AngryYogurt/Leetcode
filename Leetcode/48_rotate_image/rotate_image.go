package main

import "fmt"

func main() {
	src := [][]int{
		{1, 2, 3,},
		{4, 5, 6,},
		{7, 8, 9,},
	}
	rotate(src)
	for _, v := range src {
		fmt.Println(v)
	}

	src = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(src)
	for _, v := range src {
		fmt.Println(v)
	}
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.2 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func rotate(matrix [][]int) {
	l, le := len(matrix), len(matrix)
	tmp := 0
	round := 0
	step := 0
	for l > 1 {
		step = round
		for i := 0; i < l-1; i++ {
			tmp = matrix[round][step]
			matrix[round][step] = matrix[le-1-step][round]
			matrix[le-1-step][round] = matrix[le-1-round][le-1-step]
			matrix[le-1-round][le-1-step] = matrix[step][le-1-round]
			matrix[step][le-1-round] = tmp
			step++
		}
		round++
		l -= 2
	}
}
