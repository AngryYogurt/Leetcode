package main

func main() {

}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2 MB, 在所有 Go 提交中击败了65.48%的用户
 */
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] ++
			return digits
		}
	}
	return append([]int{1}, digits...)
}
