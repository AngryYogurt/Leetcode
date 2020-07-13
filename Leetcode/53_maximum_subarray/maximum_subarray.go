package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(3, maxSubArray([]int{1, 2, -1, -2, 2, 1, -2, 1}))
	tools.AssertObjectEqual(-1, maxSubArray([]int{-1}))
	tools.AssertObjectEqual(6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	res := Merge(nums)

	return res
}

/*
 * 执行用时 :8 ms, 在所有 Go 提交中击败了64.06%的用户
 * 内存消耗 :3.3 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func SlidingWindow(nums []int) int {
	sum := 0
	max := nums[0]
	for _, v := range nums {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

/*
 * 执行用时 :4 ms, 在所有 Go 提交中击败了96.84%的用户
 * 内存消耗 :3.3 MB, 在所有 Go 提交中击败了54.55%的用户
*/

type Ans struct {
	tSum int
	lSum int
	rSum int
	mSum int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func meg(l, r Ans) Ans {
	return Ans{
		tSum: l.tSum + r.tSum,
		lSum: max(l.tSum+r.lSum, l.lSum),
		rSum: max(l.rSum+r.tSum, r.rSum),
		mSum: max(max(l.mSum, r.mSum), l.rSum+r.lSum),
	}
}

func calc(nums []int, l, r int) Ans {
	if l == r {
		return Ans{
			tSum: nums[l],
			lSum: nums[l],
			rSum: nums[l],
			mSum: nums[l],
		}
	}
	m := (l + r) / 2
	lA := calc(nums, l, m)
	rA := calc(nums, m+1, r)
	return meg(lA, rA)
}

func Merge(nums []int) int {
	return calc(nums, 0, len(nums)-1).mSum
}
