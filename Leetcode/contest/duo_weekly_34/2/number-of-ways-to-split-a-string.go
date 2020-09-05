package main

import "fmt"

func main() {
	fmt.Println(numWays("10101"))
}

func numWays(s string) int {
	if len(s) < 3 {
		return 0
	}
	tot := 0
	for _, c := range s {
		if c == '1' {
			tot++
		}
	}
	if tot%3 != 0 {
		return 0
	}
	res := 0
	if tot == 0 {
		for i := 1; i <= len(s)-2; i++ {
			res += i
			if res > 1000000007 {
				res -= 1000000007
			}
		}
		return res
	}
	st, en := tot/3, tot/3
	i, j := 0, len(s)-1
	lf, rt := 0, 0
	lok, rok := false, false
	for {
		if s[i] == '1' {
			st--
			if st == 0 {
				lf = i
				lok = true
			}
		}
		if s[j] == '1' {
			en--
			if en == 0 {
				rt = j
				rok = true
			}
		}
		i++
		j--
		if lok && rok {
			break
		}
	}

	lfZ, rtZ := 0, 0
	for i := lf + 1; s[i] == '0'; i++ {
		lfZ++
	}
	for i := rt - 1; s[i] == '0'; i-- {
		rtZ++
	}
	for i := 0; i < rtZ+1; i++ {
		res += lfZ + 1
		if res > 1000000007 {
			res -= 1000000007
		}
	}
	return res
}
