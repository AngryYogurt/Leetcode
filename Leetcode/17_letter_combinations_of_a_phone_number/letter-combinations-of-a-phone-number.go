package main

import (
	"fmt"
)

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("23992"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("11132421189"))

}


/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.1 MB, 在所有 Go 提交中击败了8.87%的用户
 */
var km = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	return gen(digits)
}

func gen(digits string) []string {
	t := make([]string, 0, 0)
	if len(digits) <= 0{
		return t
	}
	if len(digits) == 1 {
		return append(t, km[digits]...)
	}
	for _, v := range km[string(digits[len(digits)-1])] {
		r := gen(digits[0 : len(digits)-1])
		for _, item := range r {
			t = append(t, item+v)
		}
	}
	return t
}
