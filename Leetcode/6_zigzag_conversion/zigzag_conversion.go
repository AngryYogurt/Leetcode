package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual("AB", convert("AB",1))
	tools.AssertObjectEqual("LDREOEIIECIHNTSG", convert("LEETCODEISHIRING", 4))
	tools.AssertObjectEqual("LCIRETOESIIGEDHN", convert("LEETCODEISHIRING",3))
}
/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :3.9 MB, 在所有 Go 提交中击败了99.73%的用户
 */
func convert(s string, numRows int) string {
	if len(s) <= numRows || numRows == 1{
		return s
	}
	result := make([]byte, len(s))
	idx := 0
	step := 2 * (numRows - 1)
	for i := 0; i < numRows; i++ {
		for j := 0; ; j += step {
			if j+i >= len(s) {
				break
			}
			result[idx] = s[j+i]
			idx++
			if i != 0 && i != numRows-1 && j+step-i < len(s) {
				result[idx] = s[j+step-i]
				idx++
			}
		}
	}
	return string(result)
}
