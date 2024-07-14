package main

import (
	"github.com/AngryYogurt/ProgrammingProblem/tools"
	"strings"
)

func main() {
	tools.AssertObjectEqual(true, isValid("{[]}"))
	tools.AssertObjectEqual(false, isValid("}"))
	tools.AssertObjectEqual(false, isValid("{[((("))
	tools.AssertObjectEqual(true, isValid(""))
	tools.AssertObjectEqual(false, isValid("{[}]"))
}

var mapping = map[byte]byte{
	'{': '}',
	'(': ')',
	'[': ']',
}

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for idx := range s {
		switch v := s[idx]; v {
		case '{', '(', '[':
			stack = append(stack, v)
		case '}', ')', ']':
			if len(stack) <= 0 || mapping[stack[len(stack)-1]] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

/*
 * 执行用时 :4 ms, 在所有 Go 提交中击败了9.64%的用户
 * 内存消耗 :7 MB, 在所有 Go 提交中击败了5.50%的用户
 */
func isValid3(s string) bool {
	l := len(s)
	for l > 0 {
		s = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(s, "()", ""), "{}", ""), "[]", "")
		if len(s) == l {
			return false
		} else if len(s) < l {
			l = len(s)
		}
	}
	return true
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2 MB, 在所有 Go 提交中击败了56.94%的用户
 */
func isValid2(s string) bool {
	tail := -1
	stack := make([]byte, len(s))
	for _, c := range s {
		switch c {
		case '(':
			tail++
			stack[tail] = '('
		case '[':
			tail++
			stack[tail] = '['
		case '{':
			tail++
			stack[tail] = '{'
		case ')':
			if tail < 0 || stack[tail] != '(' {
				return false
			}
			stack[tail] = 0
			tail--
		case ']':
			if tail < 0 || stack[tail] != '[' {
				return false
			}
			stack[tail] = 0
			tail--
		case '}':
			if tail < 0 || stack[tail] != '{' {
				return false
			}
			stack[tail] = 0
			tail--
		}
	}
	return tail < 0
}
