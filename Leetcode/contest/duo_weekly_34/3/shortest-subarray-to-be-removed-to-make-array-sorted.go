package main

import "fmt"

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{2, 7, 4, 6, 11, 17, 4, 1, 5, 13, 10, 12}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 3, 2, 4}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3}))
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
	fmt.Println(findLengthOfShortestSubarray([]int{5, 4, 3, 2, 1}))
	fmt.Println(findLengthOfShortestSubarray([]int{1}))
}

func findLengthOfShortestSubarray(arr []int) int {
	if len(arr) <= 1 {
		return 0
	}
	lLen, rLen := 1, 1
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			break
		}
		lLen++
	}

	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			break
		}
		rLen++
	}
	mLen := 0
	i, j := 0, len(arr)-1
	l, r := arr[i], arr[j]
	if l <= r {
		mLen = 2
		for i < j-1 {
			if ((arr[i+1]-arr[i] >= arr[j]-arr[j-1] && arr[j]-arr[j-1] >= 0) || arr[j]-arr[j-1] >= 0) && arr[j-1] >= arr[i] {
				j--
				mLen++
				continue
			} else if ((arr[i+1]-arr[i] < arr[j]-arr[j-1] && arr[i+1]-arr[i] >= 0) || arr[i+1]-arr[i] >= 0) && arr[i+1] <= arr[j] {
				i++
				mLen++
				continue
			} else {
				break
			}
		}
	}
	return len(arr) - max(lLen, max(rLen, mLen))
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
