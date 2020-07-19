package main

import "github.com/AngryYogurt/leetcode/tools"

func main() {
	tools.AssertObjectEqual(13, numWaterBottles(9, 3))
	tools.AssertObjectEqual(6, numWaterBottles(5, 5))
	tools.AssertObjectEqual(19, numWaterBottles(15, 4))
}
/*
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：1.9 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func numWaterBottles(numBottles int, numExchange int) int {
	result := numBottles
	for numBottles >= numExchange {
		result += numBottles / numExchange
		numBottles = numBottles % numExchange + numBottles / numExchange
	}
	return result
}
