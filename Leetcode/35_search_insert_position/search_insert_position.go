package main

func main() {
	searchInsert([]int{1, 3, 5, 6}, 5)
}

func searchInsert(nums []int, target int) int {
	if len(nums) <= 0 {
		return 0
	}
	return search(nums, target)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (right + left - 1) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid
		} else {
			left = mid
		}
	}
	return left
}
