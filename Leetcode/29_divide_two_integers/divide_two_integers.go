package main

import (
	"github.com/AngryYogurt/ProgrammingProblem/tools"
	"math"
)

func main() {
	tools.AssertObjectEqual(2147483647, divide(-2147483648, -1))
	tools.AssertObjectEqual(3, divide(10, 3))
	tools.AssertObjectEqual(-2, divide(7, -3))
	tools.AssertObjectEqual(1, divide(2, 2))
}

/*
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：2.5 MB, 在所有 Go 提交中击败了50.00%的用户
 */

func divide(dividend int, divisor int) int {
	if divisor == -1 && dividend == math.MinInt32 {
		return math.MaxInt32
	}
	if divisor == -1 {
		return -dividend
	}
	if divisor == 1 {
		return dividend
	}
	positive := true
	if dividend < 0 {
		positive = !positive
		dividend = -dividend
	}
	if divisor < 0 {
		positive = !positive
		divisor = -divisor
	}
	result := 0
	tmp := make([]int, 32)
	v := make([]int, 32)
	t, i, vv := divisor, 0, 1
	for t <= dividend {
		tmp[i] = t
		v[i] = vv
		t += t
		vv += vv
		i++
	}
	i--
	for i >= 0 {
		if dividend >= tmp[i] {
			dividend -= tmp[i]
			result += v[i]
		} else {
			i--
		}
	}
	if !positive {
		result = -result
	}
	return result
}
