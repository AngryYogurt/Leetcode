package main

import (
	"fmt"
	"github.com/AngryYogurt/leetcode/tools"
)

func main() {
	tools.AssertObjectEqual("6", multiply("2", "3"))
	tools.AssertObjectEqual("60000000000", multiply("20000000", "3000"))
	tools.AssertObjectEqual("19552538255754", multiply("84596879", "231126"))
}

/*
 * 执行用时 :128 ms, 在所有 Go 提交中击败了5.24%的用户
 * 内存消耗 :6.6 MB, 在所有 Go 提交中击败了5.93%的用户
 */
func multiply(num1 string, num2 string) string {
	result := "0"
	zeros := ""
	if num1 == "0" || num2 == "0"{
		return result
	}
	for i := len(num2) - 1; i >= 0; i-- {
		result = add(result, multiplyOne(num1, num2[i]) + zeros)
		zeros += "0"
	}
	return result
}

func add(num1, num2 string) string {
	result := ""
	idx1, idx2 := len(num1)-1, len(num2)-1
	n1, n2, r, tmp := 0, 0, 0, 0
	for {
		if idx1 < 0 && idx2 < 0 && tmp == 0 {
			break
		}
		if idx1 < 0 {
			n1 = 0
		} else {
			n1 = int(num1[idx1] - '0')
		}
		if idx2 < 0 {
			n2 = 0
		} else {
			n2 = int(num2[idx2] - '0')
		}
		r = n1 + n2 + tmp
		if r >= 10 {
			tmp = 1
			r = r - 10
		} else {
			tmp = 0
		}
		result = fmt.Sprintf("%d%s", r, result)
		idx1--
		idx2--
	}
	return result
}

func multiplyOne(num1 string, num2 byte) string {
	tmp := 0
	r := 0
	result := ""
	for i := len(num1) - 1; i >= 0; i-- {
		r = int(num1[i]-'0')*int(num2-'0') + tmp
		tmp = r / 10
		r = r % 10
		result = fmt.Sprintf("%d%s", r, result)
	}
	if tmp > 0 {
		result = fmt.Sprintf("%d%s", tmp, result)
	}
	return result
}
