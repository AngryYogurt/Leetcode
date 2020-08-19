package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(3, minNumberOperations([]int{1, 2, 3, 2, 1}))
	tools.AssertObjectEqual(4, minNumberOperations([]int{3,1,1,2}))
	tools.AssertObjectEqual(7, minNumberOperations([]int{3,1,5,4,2}))
}
func minNumberOperations(target []int) int {
	res := target[0]
	for i := 1; i < len(target); i++ {
		if v := target[i] - target[i-1]; v > 0 {
			res += v
		}
	}
	return res
}
