package main

import "github.com/AngryYogurt/leetcode/tools"

func main() {
	tools.AssertObjectEqual("AB", gcdOfStrings("ABABABABAB", "ABAB"))
	tools.AssertObjectEqual("", gcdOfStrings("ABABABABAB", "CDD"))
	tools.AssertObjectEqual("AAA", gcdOfStrings("AAAAAAAAAAAA", "AAAAAAAAA"))
}
/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.5 MB, 在所有 Go 提交中击败了57.14%的用户
*/
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if b > a {
		a, b = b, a
	}
	tmp := a % b
	if tmp == 0{
		return b
	}
	for tmp > 0 {
		tmp = a % b
		a = b
		b = tmp
	}
	return a
}
