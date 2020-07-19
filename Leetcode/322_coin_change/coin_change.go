package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(3, coinChange([]int{1, 2, 5}, 11))
	tools.AssertObjectEqual(-1, coinChange([]int{2}, 3))
}

/*
 * 执行用时 :68 ms, 在所有 Go 提交中击败了13.69%的用户
 * 内存消耗 :6.4 MB, 在所有 Go 提交中击败了24.74%的用户
 */
func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return 0
	}
	if len(coins) < 1 {
		return -1
	}
	return change(coins, amount, make([]int, amount+1))
}

func change(coins []int, target int, count []int) int {
	if target < 0{
		return -1
	}
	if target == 0 {
		return 0
	}
	if count[target] != 0{
		return count[target]
	}
	min := target + 1
	for _, c := range coins {
		if res := change(coins, target-c, count); res >= 0 && res < min{
			min = res + 1
		}
	}
	count[target] = min
	if min == target +1{
		count[target] = -1
	}
	return count[target]
}
