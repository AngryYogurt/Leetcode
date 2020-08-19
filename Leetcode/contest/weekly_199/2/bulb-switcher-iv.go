package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(3, minFlips("10111"))
	tools.AssertObjectEqual(3, minFlips("101"))
	tools.AssertObjectEqual(0, minFlips("00000"))
	tools.AssertObjectEqual(5, minFlips("001011101"))

}

func minFlips(target string) int {
	t := []int{0, 0}
	for i := 1; i < len(target); i++ {
		if target[i] != target[i-1] {
			t[int(target[i-1])-'0']++
		}
	}
	t[int(target[len(target)-1]-'0')]++
	need := true
	if target[len(target)-1] == '1' {
		need = false
	}
	if t[1] == 0 {
		return 0
	}

	res1 := t[1] * 2
	if !need {
		res1--
	}
	res0 := t[0] * 2+1
	if target[0] == 0 {
		res0--
	}
	if res0 < res1 {
		return res0
	}
	return res1
}
