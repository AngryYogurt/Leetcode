package main

import "fmt"

func main() {
	fmt.Println(minOperations([]int{4,2, 5}))
}
func minOperations(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := -1
	res := 0
	for i := 0; i < len(nums); i++ {
		t := nums[i]
		ct := 0
		for t > 0 {
			if t&1 == 1 {
				res++
			}
			t >>= 1
			ct++
		}
		if ct-1 > max {
			max = ct - 1
		}
	}
	return res + max
}
