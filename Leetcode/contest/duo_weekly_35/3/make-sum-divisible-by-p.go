package main

import "fmt"

func main() {
	fmt.Println(minSubarray([]int{8, 32, 31, 18, 34, 20, 21, 13, 1, 27, 23, 22, 11, 15, 30, 4, 2}, 148)) //7
	fmt.Println(minSubarray([]int{6, 3, 5, 2}, 9))                                                       //2
	fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6))                                                       //1

}
func minSubarray(nums []int, p int) int {
	if len(nums) <= 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0]%p == 0 {
			return 0
		} else {
			return -1
		}
	}

	m := make(map[int][]int)

	sum := 0
	lsum := make([]int, len(nums))
	rsum := make([]int, len(nums))
	for i, v := range nums {
		lsum[i] = sum
		sum += v
	}

	if sum%p == 0 {
		return 0
	}
	rsum[len(nums)-1] = 0
	m[0] = []int{len(nums) - 1}
	for i := len(nums) - 2; i >= 0; i-- {
		rsum[i] = nums[i+1] + rsum[i+1]
		t := rsum[i] % p
		if _, ok := m[t]; !ok {
			m[t] = make([]int, 0)
		}
		m[t] = append(m[t], i)
	}

	res := len(nums)
	if len(m[0]) > 1 {
		res = m[0][1] + 1
	}
	for i := 0; i < len(nums); i++ {
		t := lsum[i] % p
		if t == 0 && i < len(nums)-1 && len(nums)-1-i < res {
			res = len(nums) - i
		}
		for _, j := range m[p-t] {
			if j >= i && j-i+1 < res {
				res = j - i + 1
			}
		}
	}
	if res < len(nums) {
		return res
	}
	return -1
}
