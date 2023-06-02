package quick_sort

func quickSort(nums []int) {
	_quickSort(nums, 0, len(nums)-1)
}

func _quickSort(nums []int, start, end int) {
	l, r := start, end
	if start > end {
		return
	}
	tmp := nums[start]
	for l <= r {
		for nums[l] < tmp {
			l++
		}
		for nums[r] > tmp {
			r--
		}
		if l <= r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	_quickSort(nums, l, end)
	_quickSort(nums, start, r)
}
