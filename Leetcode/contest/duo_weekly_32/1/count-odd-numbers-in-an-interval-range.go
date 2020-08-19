package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 5))
}

func findKthPositive(arr []int, k int) int {
	idx := 0
	for i := 1; idx < len(arr); i++ {
		if arr[idx] == i {
			idx++
		} else {
			k--
			if k == 0 {
				return i
			}
		}
	}
	return arr[len(arr)-1] + k
}
