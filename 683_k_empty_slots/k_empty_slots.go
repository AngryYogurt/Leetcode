package main

import (
	"github.com/AngryYogurt/leetcode/tools"
	"math"
)

func main() {
	tools.AssertObjectEqual(2, kEmptySlots([]int{1, 3, 2}, 1))
	tools.AssertObjectEqual(-1, kEmptySlots([]int{1, 2, 3}, 1))
	tools.AssertObjectEqual(8, kEmptySlots([]int{6, 5, 8, 9, 7, 1, 10, 2, 3, 4}, 2))
	tools.AssertObjectEqual(6, kEmptySlots([]int{3, 9, 2, 8, 1, 6, 10, 5, 4, 7}, 1))
}

/*
 * 执行用时 :172 ms, 在所有 Go 提交中击败了50.00%的用户
 * 内存消耗 :7 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func kEmptySlots(bulbs []int, K int) int {
	ot := make([]int, len(bulbs)+1)
	for idx, v := range bulbs {
		ot[v] = idx + 1
	}
	start, end := 1, K+2
	minRes := len(bulbs) + 1
	fail := false
	for end < len(ot) {
		idx := start + 1
		fail = false
		for ; idx <= end-1; idx++ {
			if ot[idx] < ot[start] || ot[idx] < ot[end] {
				start, end = idx, idx+K+1
				fail = true
				break
			}
		}
		if !fail {
			minRes = int(math.Min(float64(minRes), math.Max(float64(ot[start]), float64(ot[end]))))
			start, end = end, end+K+1
		}
	}
	if minRes == len(bulbs)+1 {
		return -1
	}
	return minRes
}
