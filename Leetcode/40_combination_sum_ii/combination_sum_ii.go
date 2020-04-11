package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{1, 2, 3, 6, 7}, 7))
	fmt.Println(combinationSum2([]int{7, 3, 2}, 18))
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.7 MB, 在所有 Go 提交中击败了51.18%的用户
 */
type Result struct {
	result     [][]int
	target     int
	candidates []int
}

func combinationSum2(candidates []int, target int) [][]int {
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

	gen(i+1, append(curr, result.candidates[i]), currSum+result.candidates[i], result)

	m := i
	for ; m < len(result.candidates); m++ {
		if result.candidates[m] != result.candidates[i] {
			break
		} else if m == len(result.candidates)-1 {
			return
		}
	}
	gen(m, curr, currSum, result)
}
