/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	mm, nn := m-1, n-1
	idx := m + n - 1
	for nn >= 0 {
		if mm >= 0 && nums1[mm] > nums2[nn] {
			nums1[idx] = nums1[mm]
			mm--
		} else {
			nums1[idx] = nums2[nn]
			nn--
		}
		idx--
	}
}

// @lc code=end

