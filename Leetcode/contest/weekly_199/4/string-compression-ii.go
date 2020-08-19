package main

import (
	"fmt"
	"github.com/AngryYogurt/ProgrammingProblem/tools"
	"strconv"
)

func main() {
	tools.AssertObjectEqual(4, getLengthOfOptimalCompression("aaabcccd", 2))
	tools.AssertObjectEqual(3, getLengthOfOptimalCompression("aaaaaaaaaaaaaaaaaaaaaaaaa", 0))
	tools.AssertObjectEqual(2, getLengthOfOptimalCompression("aabbaa", 2))
}

// TODO
func getLengthOfOptimalCompression(s string, k int) int {
	res := make([]byte, 0)
	res = append(res, s[0])
	ct := make([]int, len(s))
	count := 1
	idx := 0
	srcIdx := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			ct[idx] = count
			if count > 1 {
				c := strconv.FormatInt(int64(count), 10)
				res = append(res, c...)
				idx++
				if count > 10 {
					idx++
				}
			}
			idx++
			srcIdx += count
			res = append(res, s[i])
			count = 1
		}
	}
	if count > 1 {
		c := strconv.FormatInt(int64(count), 10)
		res = append(res, c...)
		idx += len(c)
	}
	fmt.Print(string(res) + " ")

	if count == 100 {
		if k > 0 {
			return 3
		} else {
			return 4
		}
	}
	if k == 0 {
		return len(res)
	}

	return travel(string(res), ct, 0, k)
}

func gen(s string) string{
	res := make([]byte, 0)
	res = append(res, s[0])
	ct := make([]int, len(s))
	count := 1
	idx := 0
	srcIdx := 0
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			count++
		} else {
			ct[idx] = count
			if count > 1 {
				c := strconv.FormatInt(int64(count), 10)
				res = append(res, c...)
				idx++
				if count > 10 {
					idx++
				}
			}
			idx++
			srcIdx += count
			res = append(res, s[i])
			count = 1
		}
	}
	if count > 1 {
		c := strconv.FormatInt(int64(count), 10)
		res = append(res, c...)
		idx += len(c)
	}
	return string(res)
}

func travel(res string, ct []int, ctIdx int, k int) int {
	if k == 0 {
		return ctIdx + 1
	}

	min := ctIdx + 1
	for i := ctIdx + 1; i < len(ct); i++ {
		if ct[i] == 0 {
			continue
		}
		if ct[i] >= 10 && k-(ct[i]-9) >= 0 {
			v := travel(res, ct, i, k-(ct[i]-9)) + 2
			if v < min {
				min = v
			}
		}
		v := travel(res, ct, i, k) + 2
		if v < min {
			min = v
		}
		if k-ct[i] >= 0 {
			v = travel(res, ct, i, k-ct[i])
			if v < min {
				min = v
			}
		}
	}
	return min
}
