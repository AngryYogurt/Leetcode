package main

func main() {

}
/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.2 MB, 在所有 Go 提交中击败了25.00%的用户
 */
func grayCode(n int) []int {
	result := make([]int, 0, n*2)
	result = append(result, 0)
	head := 1
	for i := 0; i < n; i++ {
		for j := len(result) - 1; j >= 0; j-- {
			result = append(result, head+result[j])
		}
		head <<= 1
	}
	return result
}
