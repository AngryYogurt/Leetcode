package main

import "fmt"

func main() {

	fmt.Println()
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	res := make([][]int, len(rowSum))
	for i := 0; i < len(rowSum); i++ {
		res[i] = make([]int, len(colSum))
	}
	for i := 0; i < len(rowSum); i++ {
		for j := 0;j < len(colSum);j++{
			// 不会
		}
	}
	return res
}
