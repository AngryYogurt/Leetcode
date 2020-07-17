package main

func main() {

}
/*
 * 执行用时：4 ms, 在所有 Go 提交中击败了17.00%的用户
 * 内存消耗：2.7 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func runningSum(nums []int) []int {
	if len(nums) <= 0 {
		return []int{}
	}
	res := make([]int, len(nums))
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = nums[i] + res[i-1]
	}
	return res
}
