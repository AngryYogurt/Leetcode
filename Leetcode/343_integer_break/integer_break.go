package main

import "math"

func main() {

}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :1.9 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func integerBreak(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	} else if n == 3 {
		return 2
	}
	p := n / 3
	result := int64(1)
	if n%3 == 1 {
		p--
		result = 4
	} else if n%3 == 2 {
		result = 2
	}

	return int(result * int64(math.Pow(float64(3), float64(p))))
}
