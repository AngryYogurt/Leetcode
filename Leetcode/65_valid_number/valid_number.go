package main

import (
	"github.com/AngryYogurt/leetcode/tools"
	"regexp"
	"strings"
)

func main() {
	tools.AssertObjectEqual(true, isNumber("-90e3"))
	tools.AssertObjectEqual(false, isNumber("1e0.1"))
	tools.AssertObjectEqual(false, isNumber("e3"))
	tools.AssertObjectEqual(false, isNumber("1e"))
	tools.AssertObjectEqual(true, isNumber("1 "))
	tools.AssertObjectEqual(false, isNumber("."))
	tools.AssertObjectEqual(true, isNumber("0"))
	tools.AssertObjectEqual(true, isNumber("0.2"))
	tools.AssertObjectEqual(false, isNumber("abc"))
	tools.AssertObjectEqual(false, isNumber("1 a"))
	tools.AssertObjectEqual(true, isNumber("2e10"))
	tools.AssertObjectEqual(true, isNumber("02"))
	tools.AssertObjectEqual(true, isNumber("2e02"))
	tools.AssertObjectEqual(true, isNumber(".1"))
	tools.AssertObjectEqual(false, isNumber("5e.1"))
	tools.AssertObjectEqual(false, isNumber(".e1"))
	tools.AssertObjectEqual(true, isNumber("5."))
	tools.AssertObjectEqual(true, isNumber("5.e1"))
	tools.AssertObjectEqual(true, isNumber("-1."))
	tools.AssertObjectEqual(false, isNumber(""))

}
/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.8 MB, 在所有 Go 提交中击败了26.47%的用户
 */
var reg = "^\\s*[+|-]?(([0-9]*\\.?[0-9]+)|([0-9]+\\.?[0-9]*)|[0-9]+)(e[+|-]?[0-9]+)?\\s*$"
var r, _ = regexp.Compile(reg)
func isNumber(s string) bool {
	return r.Match([]byte(s))
}

// 未完成的题解：
var num = map[byte]bool{
	'0': true,
	'1': true,
	'2': true,
	'3': true,
	'4': true,
	'5': true,
	'6': true,
	'7': true,
	'8': true,
	'9': true,
}

func isNumber2(s string) bool  {
	s = strings.Trim(s, " ")
	canE := true
	canPM := true
	canPoint := true
	haveNum := false

	pointCount := strings.Count(s, ".")
	eCount := strings.Count(s, "e")
	pmCount := strings.Count(s, "+") + strings.Count(s, "-")
	if pointCount > 1 || eCount > 1 || pmCount-eCount > 1 {
		return false
	}
	if len(s) <= 0 || (len(s) == 1 && !num[s[0]]) {
		return false
	}
	for i := len(s) - 1; i >= 1; i-- {
		c := s[i]
		if i == len(s)-1 && !num[c] && c != '.' {
			return false
		}
		switch c {
		case 'e':
			if !canE || !num[s[i-1]] {
				return false
			}
			canE = false
			canPM = false
		case '+', '-':
			if !canPM {
				return false
			}
			if s[i-1] != 'e' {
				return false
			}
		case '.':
			if !canPoint {
				return false
			}
			canE = false
			canPoint = false

		default:
			if !num[c] {
				return false
			} else {
				haveNum = true
			}
		}
	}
	canE = false
	if c := s[0]; !num[c] {
		if c != '.' && c != '+' && c != '-' {
			return false
		}
		if (c == '.' && !canPoint) || (!canPM && (c == '+' || c == '-')) {
			return false
		}
	} else {
		haveNum = true
	}
	return haveNum
}
