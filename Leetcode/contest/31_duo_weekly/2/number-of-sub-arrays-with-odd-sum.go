package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(45, numOfSubarrays([]int{97,10,8,74,51,1,14,84,2,63,33,29,28}))
	tools.AssertObjectEqual(16, numOfSubarrays([]int{1, 2, 3, 4, 5, 6, 7}))
	tools.AssertObjectEqual(4, numOfSubarrays([]int{1, 3, 5}))
	tools.AssertObjectEqual(0, numOfSubarrays([]int{2, 4, 6}))
}

var M = 1000000007

func numOfSubarrays(arr []int) int {
	oddIdx := make([]int, 0)
	last := -1
	for i, v := range arr {
		if v&1 == 1 {
			if i == 0 {
				oddIdx = append(oddIdx, 1)
			} else {
				oddIdx = append(oddIdx, i-last)
			}
			last = i
		}
	}
	oddIdx = append(oddIdx, len(arr) - last)
	res := 0
	sum := 0
	for i := len(oddIdx) - 1; i >= 1; i -= 2 {
		sum += oddIdx[i]
		res += sum * oddIdx[i-1]
		if res > M {
			res %= M
		}
	}
	sum = 0
	for i := len(oddIdx) - 2; i >= 1; i -= 2 {
		sum += oddIdx[i]
		res += sum * oddIdx[i-1]
		if res > M {
			res %= M
		}
	}
	return res
}
