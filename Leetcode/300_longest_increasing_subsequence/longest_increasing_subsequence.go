package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{2, 2}))
}
/*
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：2.3 MB, 在所有 Go 提交中击败了95.42%的用户
 */
func lengthOfLIS(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	tails := make([]int, 0, len(nums))
	tails = append(tails, nums[0])
	for i := 1; i < len(nums); i++ {
		l, r := 0, len(tails)
		if tails[len(tails)-1] < nums[i] {
			tails = append(tails, nums[i])
		} else {
			for l < r {
				if tails[l+(r-l)/2] < nums[i] {
					l = l + (r-l)/2 + 1
				} else {
					r = l + (r-l)/2
				}
			}
			tails[l] = nums[i]
		}
	}
	return len(tails)
}
