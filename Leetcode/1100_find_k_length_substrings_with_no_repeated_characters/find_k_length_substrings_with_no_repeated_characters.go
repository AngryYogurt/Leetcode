package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(6, numKLenSubstrNoRepeats("havefunonleetcode", 5))
}
/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.2 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func numKLenSubstrNoRepeats(S string, K int) int {
	if len(S) < K || K > 26 {
		return 0
	}
	res := 0
	appeared := [26]int{}
	ok := true
	for i := 0; i < len(S)-K+1; i++ {
		if i == 0 {
			first := true
			for j := 0; j < K; j++ {
				if appeared[S[j]-'a'] > 0 {
					first = false
				}
				appeared[S[j]-'a']++
			}
			if first {
				res++
			}
			continue
		}
		appeared[S[i-1]-'a']--
		appeared[S[i+K-1]-'a']++
		ok = true
		for _, v := range appeared {
			if v > 1 {
				ok = false
				break
			}
		}
		if ok {
			res++
		}
	}
	return res
}
