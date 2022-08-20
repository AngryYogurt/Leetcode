package main

import (
	"fmt"
	"math/rand"
)

func main() {
	q := []int{3, 2, 1, 5, 6, 4}
	fmt.Println(findKthLargest(q, 2))
}

/*
 * 执行用时：4 ms, 在所有 Go 提交中击败了99.10%的用户
 * 内存消耗：3.4 MB, 在所有 Go 提交中击败了78.75%的用户
 */

func findKthLargest(nums []int, k int) int {

	idx := qSortP(nums, 0, len(nums)-1)
	l, r := 0, len(nums)-1
	for idx != len(nums)-k {
		if idx > len(nums)-k {
			r = idx - 1
		} else if idx < len(nums)-k {
			l = idx + 1
		}
		idx = qSortP(nums, l, r)
	}
	return nums[idx]
}

func qSortP(a []int, l, r int) int {
	rr := rand.Intn(r-l+1) + l
	a[rr], a[l] = a[l], a[rr]

	x := a[l]
	s := l
	for l < r {
		for l < r && a[r] > x {
			r--
		}
		for l < r && a[l] <= x {
			l++
		}
		if l < r {
			a[l], a[r] = a[r], a[l]
		}
	}
	a[s], a[l] = a[l], a[s]
	return l
}
