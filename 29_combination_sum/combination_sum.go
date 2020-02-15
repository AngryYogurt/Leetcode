package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{1, 2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{7, 3, 2}, 18))
}

/*
 * 执行用时 :4 ms, 在所有 Go 提交中击败了83.04%的用户
 * 内存消耗 :3.3 MB, 在所有 Go 提交中击败了45.54%的用户
 */

type Result struct {
	result     [][]int
	target     int
	candidates []int
}

func combinationSum(candidates []int, target int) [][]int {
	var result = Result{}
	result.target = target
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
	result.candidates = candidates
	gen(0, make([]int, 0), 0, &result)
	return result.result
}

func gen(i int, curr []int, currSum int, result *Result) {
	if currSum == result.target {
		result.result = append(result.result, append(make([]int, 0), curr...))
	}
	if i >= len(result.candidates) {
		return
	}
	if currSum+result.candidates[i] > result.target {
		return
	}
	gen(i, append(curr, result.candidates[i]), currSum+result.candidates[i], result)
	gen(i+1, curr, currSum, result)
}
