package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(permuteUnique([]int{1, 1, 2}))
	//fmt.Println(permuteUnique([]int{1, 1}))
	fmt.Println(permuteUnique([]int{1, 2, 2, 4, 3, 4, 5}))
	fmt.Println(permuteUnique([]int{}))
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :4.4 MB, 在所有 Go 提交中击败了30.23%的用户
 */
type Result struct {
	result [][]int
}

func permuteUnique(nums []int) [][]int {
	result := Result{
		result: make([][]int, 0),
	}
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })
	gen(nums, []int{}, &result)
	return result.result
}

func gen(cands, curr []int, result *Result) {
	if len(cands) <= 0 {
		result.result = append(result.result, curr)
		return
	}
	for idx, _ := range cands {
		if idx > 0 && cands[idx] == cands[idx-1] {
			continue
		}
		gen(append(append(make([]int, 0), cands[:idx]...), cands[idx+1:]...), append(append(make([]int, 0), curr...), cands[idx]), result)
	}
}
