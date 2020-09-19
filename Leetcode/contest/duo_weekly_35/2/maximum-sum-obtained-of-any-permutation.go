package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println()
}

func maxSumRangeQuery(nums []int, requests [][]int) int {
	in, out := make([]int, len(nums)+5), make([]int, len(nums)+5)
	for i, _ := range requests {
		in[requests[i][0]]++
		out[requests[i][1]]++
	}
	ts := make([]int, len(nums))
	curr := 0
	for i := 0; i < len(nums); i++ {
		curr += in[i]
		ts[i] = curr
		curr -= out[i]
	}
	sort.Slice(ts, func(i, j int) bool {
		return ts[i] < ts[j]
	})
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i] * ts[i]
		if sum > 1000000007 {
			sum -= 1000000007
		}
	}
	return sum
}
