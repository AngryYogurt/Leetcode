package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(2, numSplits("aacaba"))
	tools.AssertObjectEqual(4, numSplits("aaaaa"))
	tools.AssertObjectEqual(1, numSplits("abcd"))
}

func numSplits(s string) int {
	if len(s) <= 0 {
		return 0
	}
	res := 0
	l, r := make([]int, len(s)), make([]int, len(s))
	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[int(s[i]-'a')] = 1
		l[i] = len(m)
		if len(m) == 26 {
			break
		}
	}
	rr := len(s) - 1
	m = make(map[int]int)
	for ; rr >= 0; rr-- {
		m[int(s[rr]-'a')] = 1
		r[rr] = len(m)
		if len(m) == 26 {
			break
		}
	}
	if rr > len(l) {
		res += rr - len(l)
	} else {
		for i := rr+2; i < len(l); i++ {
			if r[i] == l[i-1] {
				res++
			}
		}
	}
	return res
}
