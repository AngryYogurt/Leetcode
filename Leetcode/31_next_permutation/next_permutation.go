package main

func main() {
	nextPermutation([]int{1, 2, 3})
	nextPermutation([]int{3, 2, 1})
	nextPermutation([]int{2, 1, 3})
	nextPermutation([]int{1, 1, 3})
	nextPermutation([]int{1, 2, 3, 4, 5, 10, 9, 8, 7, 6, 4, 3})
}

/*
 * 执行用时 :4 ms, 在所有 Go 提交中击败了70.62%的用户
 * 内存消耗 :2.5 MB, 在所有 Go 提交中击败了100.00%的用户
 */

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			for j := len(nums) - 1; j > i; j-- {
				if nums[j] > nums[i] {
					swap(nums, i, j)
					reverse(nums, i+1, len(nums) - 1)
					return
				}
			}
		}
	}
	reverse(nums, 0, len(nums)-1)
}
func reverse(nums []int, s, e int) {
	for i := s; ; i++ {
		if i < e-i+s {
			swap(nums, i, e-i+s)
		} else {
			return
		}
	}
}

func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
	return
}
