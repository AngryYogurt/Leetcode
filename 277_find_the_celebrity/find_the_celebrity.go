package main

func main() {

}

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */

/*
 * 执行用时 :56 ms, 在所有 Go 提交中击败了92.00%的用户
 * 内存消耗 :6.3 MB, 在所有 Go 提交中击败了77.78%的用户
*/
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		cand := 0
		for i := 1; i < n; i++ {
			if knows(cand, i) {
				cand = i
			}
		}
		for i := 0; i < n; i++ {
			if i == cand {
				continue
			}
			if !knows(i, cand) {
				return -1
			}
			if i < cand && knows(cand, i){
				return -1
			}
		}
		return cand
	}
}
