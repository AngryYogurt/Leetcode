package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(1, numDupDigitsAtMostN(20))
	tools.AssertObjectEqual(262, numDupDigitsAtMostN(1000))
}
/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func numDupDigitsAtMostN(N int) int {
	if N <= 10 {
		return 0
	}
	digit := make([]int, 0)
	n := N
	for n > 0 {
		digit = append(digit, n%10)
		n /= 10
	}
	result := 0
	for idx := len(digit) - 2; idx >= 0; idx-- {
		t1 := 9
		for i := idx - 1; i >= 0; i-- {
			t1 *= 9 - (idx - 1 - i)
		}
		result += t1
	}
	//
	used := make([]int, 10)
	t2 := 0
	for idx := len(digit) - 1; idx >= 0; idx-- {
		j := 0
		if idx == len(digit)-1 {
			j = 1
		}
		for ; j < digit[idx]; j++ {
			if used[j] > 0 {
				continue
			}
			t2 += A(10-(len(digit)-idx), idx)
		}
		if used[digit[idx]] > 0 {
			break
		}
		used[digit[idx]]++
		if idx == 0 {
			t2 += 1
		}
	}
	result += t2
	result = N - result
	return result
}

func A(a, b int) int {
	res := 1
	for i := b; i > 0; i-- {
		res *= a-(b-i)
	}
	return res
}
