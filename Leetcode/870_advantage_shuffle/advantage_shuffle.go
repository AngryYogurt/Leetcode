package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(advantageCount([]int{8, 2, 4, 4, 5, 6, 6, 0, 4, 7}, []int{0, 8, 7, 4, 4, 2, 8, 5, 2, 0}))
	fmt.Println(advantageCount([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(advantageCount([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
}
/*
 * 执行用时 :108 ms, 在所有 Go 提交中击败了66.67%的用户
 * 内存消耗 :6.7 MB, 在所有 Go 提交中击败了50.00%的用户
 */
func advantageCount(A []int, B []int) []int {
	result := make([]int, len(B))
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	idx := 0
	for i, v := range B {
		idx = findIdx(A, v)
		result[i] = A[idx]
		A = append(A[:idx], A[idx+1:]...)
	}
	return result
}

func findIdx(n []int, target int) int {
	l, r := 0, len(n)-1
	m := 0
	if n[r] <= target || n[l] > target {
		return 0
	}
	for l <= r {
		m = l + (r-l)>>1
		if n[m] < target {
			l = m + 1
		} else if n[m] > target {
			r = m - 1
		} else {
			break
		}
	}
	if n[m] <= target {
		for m < len(n) {
			if n[m] > target {
				break
			} else if m == len(n)-1 {
				m = 0
				break
			}
			m++
		}
	}
	return m
}
