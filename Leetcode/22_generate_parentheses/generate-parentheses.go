package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(0))
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.7 MB, 在所有 Go 提交中击败了84.03%的用户
 */
type Result struct {
	result []string
}

func generateParenthesis(n int) []string {
	r := Result{
		result: make([]string, 0, 0),
	}

	gen(n, n, "", &r)
	return r.result
}

func gen(l, r int, curr string, res *Result) {
	if l == r && l == 0 {
		res.result = append(res.result, curr)
		return
	} else if l < r {
		if l > 0 {
			gen(l-1, r, curr+"(", res)
		}
		gen(l, r-1, curr+")", res)
	} else if l == r {
		gen(l-1, r, curr+"(", res)
	}
}
