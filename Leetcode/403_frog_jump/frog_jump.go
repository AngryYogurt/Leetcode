package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(true, canCross([]int{0, 1, 3, 6, 10, 15, 16, 21}))
	tools.AssertObjectEqual(true, canCross([]int{0, 1, 3, 5, 6, 8, 12, 17}))
	tools.AssertObjectEqual(false, canCross([]int{0, 1, 2, 3, 4, 8, 9, 11}))
}
/*
 * 执行用时 :100 ms, 在所有 Go 提交中击败了16.67%的用户
 * 内存消耗 :7 MB, 在所有 Go 提交中击败了8.33%的用户
 */
func canCross(stones []int) bool {
	if len(stones) <= 0 {
		return true
	}
	m := make(map[int]map[int]bool)
	m[stones[0]] = map[int]bool{0: true, 1: true}
	max := 0
	for i := 1; i < len(stones); i++ {
		m[stones[i]] = make(map[int]bool)
		for j := i; j >= 0; j-- {
			if step := stones[i] - stones[j]; m[stones[j]][step] {
				m[stones[i]][step] = true
				m[stones[i]][step+1] = true
				m[stones[i]][step-1] = true
				if max < step{
					max = step
				}
			} else if step > i || step > max + 1{ // 剪枝
				break
			}
		}
	}
	if len(m[stones[len(stones)-1]]) > 0 {
		return true
	}
	return false
}
