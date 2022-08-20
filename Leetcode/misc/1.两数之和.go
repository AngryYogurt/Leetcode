/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	set := make(map[int]int, len(nums)-1)
	for idx, v := range nums {
		if i, ok := set[target-v]; ok {
			return []int{idx, i}
		}
		set[v] = idx
	}
	return []int{0, 0}
}

// @lc code=end

