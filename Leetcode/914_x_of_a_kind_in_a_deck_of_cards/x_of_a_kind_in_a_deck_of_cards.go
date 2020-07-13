package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(true, hasGroupsSizeX([]int{1, 2, 3, 4, 4, 3, 2, 1}))
}

/*
 * 执行用时 :16 ms, 在所有 Go 提交中击败了93.43%的用户
 * 内存消耗 :6.3 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func hasGroupsSizeX(deck []int) bool {
	if len(deck) < 2 {
		return false
	}
	count := make([]int, 10000)
	for _, v := range deck {
		count[v]++
	}
	g := -1
	for _, v := range count {
		if v <= 0 {
			continue
		}
		if g == -1 {
			g = v
		} else {
			g = gcd(g, v)
		}
	}
	return g >= 2
}

func gcd(x, y int) int {
	for x != 0 {
		x, y = y%x, x
	}
	return y
}
