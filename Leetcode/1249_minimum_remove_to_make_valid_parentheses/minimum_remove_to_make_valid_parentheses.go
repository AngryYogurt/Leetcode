package main

import (
	"fmt"
)

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}
/*
 * 执行用时：12 ms, 在所有 Go 提交中击败了60.34%的用户
 * 内存消耗：6.4 MB, 在所有 Go 提交中击败了35.29%的用户
 */
func minRemoveToMakeValid(s string) string {
	res := []byte(s)
	rep('(', ')', res)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	rep(')', '(', res)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	rr := make([]byte, 0)
	for _, v := range res {
		if v != '.' {
			rr = append(rr, v)
		}
	}
	return string(rr)
}

func rep(a1, a2 byte, res []byte) {
	ct := 0
	for i := 0; i < len(res); i++ {
		if res[i] == a1 {
			ct++
		} else if res[i] == a2 {
			if ct > 0 {
				ct--
			} else {
				res[i] = '.'
			}
		} else {
			continue
		}
	}
}
