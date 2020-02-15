package main

import (
	"fmt"
)

func main() {
	fmt.Println(permute([]int{1}))
	fmt.Println(permute([]int{1, 2, 3, 4, 5}))
	fmt.Println(permute([]int{}))
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.8 MB, 在所有 Go 提交中击败了46.13%的用户
 */
type Result struct {
	nums   []int
	result [][]int
}

func permute(nums []int) [][]int {
	result := Result{
		nums:   nums,
		result: make([][]int, 0),
	}

	gen(nums, []int{}, &result)
	return result.result
}

func gen(cands, curr []int, result *Result) {
	if len(cands) <= 0 {
		return
	} else if len(cands) == 1 {
		result.result = append(result.result, append(curr, cands[0]))
		return
	}
	for idx, _ := range cands {
		gen(append(append(make([]int, 0), cands[:idx]...), cands[idx+1:]...), append(curr, cands[idx]), result)
	}
}
