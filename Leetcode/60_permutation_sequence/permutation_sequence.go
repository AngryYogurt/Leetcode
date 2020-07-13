package main

import (
	"fmt"
	"github.com/AngryYogurt/ProgrammingProblem/tools"
	"math"
)

func main() {
	tools.AssertObjectEqual("312", getPermutation(3, 5))
	tools.AssertObjectEqual("231", getPermutation(3, 4))
	tools.AssertObjectEqual("365421", getPermutation(6, 360))
	tools.AssertObjectEqual("412356", getPermutation(6, 361))
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.1 MB, 在所有 Go 提交中击败了28.95%的用户
 */

func getPermutation(n int, k int) string {
	t := 1
	cands := make([]int, 0)
	for i := 1; i <= n; i++ {
		t *= i
		cands = append(cands, i)
	}

	result := ""
	for i := n; i > 0; i-- {
		scope := int(math.Ceil(float64(k) / (float64(t) / float64(i))))
		result = fmt.Sprintf("%s%d", result, cands[scope-1])
		k = k - (scope-1)*(t/i)
		t /= i
		if len(cands) > 1 {
			cands = append(cands[:scope-1], cands[scope:]...)
		} else {
			break
		}
	}
	return result
}
