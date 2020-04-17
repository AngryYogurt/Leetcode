package main

func main() {

}
/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.5 MB, 在所有 Go 提交中击败了42.86%的用户
 */

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	rotIdx := findRotateIndex(nums)
	return nums[rotIdx]
}

func findRotateIndex(nums []int) int {
	start, end := 0, len(nums)-1
	if nums[end] > nums[start] {
		return 0
	}
	for start < end {
		mid := (start + end) / 2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		} else if nums[mid] < nums[start] {
			end = mid
		} else if nums[mid+1] > nums[end] {
			start = mid + 1
		}
	}
	return 0
}
