package main

import "fmt"

func main() {
	fmt.Println(minInsertions("(((()(()((())))(((()())))()())))(((()(()()((()()))"))
}

func minInsertions(s string) int {
	l := 0
	r := 0
	res := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if r%2 == 0 {
				res += r / 2
			} else {
				if l > 0 {
					res++
					l--
				} else {
					res += (r+1)/2 + 1
				}
			}
			l++
			r = 0
		} else {
			r++
			if l > 0 && r >= 2 {
				r -= 2
				l--
			}
		}
	}
	if l > 0 {
		res += l*2 - r
	} else if r > 0 {
		if r%2 == 0 {
			res += r / 2
		} else {
			res += (r+1)/2 + 1
		}
	}
	return res
}
