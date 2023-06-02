package main

import (
	"sort"
	"strconv"
)

func main() {

}

func largestNumber(nums []int) string {
	result := ""
	sort.Slice(nums, func(i, j int) bool {
		a := strconv.Itoa(nums[i])
		b := strconv.Itoa(nums[j])
		return a+b > b+a
	})
	if nums[0] == 0 {
		return "0"
	}
	for idx := range nums {
		result += strconv.Itoa(nums[idx])
	}
	return result
}
