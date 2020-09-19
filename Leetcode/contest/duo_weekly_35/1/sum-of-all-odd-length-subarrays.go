package main

import "fmt"

func main() {
	fmt.Println(sumOddLengthSubarrays([]int{1, 4, 2, 5, 3}))
}
func sumOddLengthSubarrays(arr []int) int {
	if len(arr) <= 0 {
		return 0
	}
	tl := len(arr)
	sum := 0
	for i, v := range arr {
		l := i
		r := tl - i - 1
		ct := ((l/2+1)*(r/2+1)) + ((l+1)/2*((r+1)/2))
		sum += ct*v
	}
	return sum
}
