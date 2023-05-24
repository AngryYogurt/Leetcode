package main

import (
	"github.com/AngryYogurt/ProgrammingProblem/tools"
)

func main() {
	tools.AssertObjectEqual("cadb", removeDuplicateLetters("cbaddabaa"))
	tools.AssertObjectEqual("acdb", removeDuplicateLetters("cbacdcbc"))
	tools.AssertObjectEqual("abc", removeDuplicateLetters("bcabc"))
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.1 MB, 在所有 Go 提交中击败了81.48%的用户
 */
func removeDuplicateLetters2(s string) string {
	res := make([]byte, 0)
	used := make([]bool, 26)
	m := make(map[byte]int)
	for i := len(s) - 1; i >= 0; i-- {
		if _, ok := m[s[i]]; !ok {
			m[s[i]] = i
		}
	}
	for i := 0; i < len(s); i++ {
		if used[int(s[i]-'a')] {
			continue
		}
		j := len(res) - 1
		for ; j >= 0 && res[j] > s[i] && m[res[j]] > i; j-- {
			used[int(res[j]-'a')] = false
		}
		res = append(res[:j+1], s[i])
		used[int(s[i]-'a')] = true
	}
	return string(res)
}

func removeDuplicateLetters(s string) string {
	last := [26]int{}
	for idx, v := range s {
		last[v-'a'] = idx
	}
	stack := make([]byte, len(s))
	inStack := [26]bool{}
	idx := -1
	for i := range s {
		v := s[i]
		if inStack[v-'a'] {
			continue
		}
		for idx >= 0 {
			if stack[idx] > v && last[stack[idx]-'a'] > i {
				inStack[int(stack[idx]-'a')] = false
				idx--
			} else {
				break
			}
		}
		idx++
		stack[idx] = v
		inStack[v-'a'] = true
	}
	return string(stack[:idx+1])
}
