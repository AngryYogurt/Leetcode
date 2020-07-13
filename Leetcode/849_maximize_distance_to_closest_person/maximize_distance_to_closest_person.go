package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(1, maxDistToClosest([]int{1, 0, 1}))
	tools.AssertObjectEqual(1, maxDistToClosest([]int{0, 1}))
	tools.AssertObjectEqual(2, maxDistToClosest([]int{1, 0, 0, 0, 1, 0, 1}))
	tools.AssertObjectEqual(3, maxDistToClosest([]int{1, 0, 0, 0}))
}
/*
 * 执行用时 :12 ms, 在所有 Go 提交中击败了96.77%的用户
 * 内存消耗 :5.4 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func maxDistToClosest(seats []int) int {
	res := 0
	j := len(seats) - 1
	for ; j >= 0; j-- {
		if seats[j] != 0 {
			break
		} else {
			res++
		}
	}
	i := 0
	for ; i < len(seats); i++ {
		if seats[i] != 0 {
			break
		}
	}
	if i > res {
		res = i
	}
	length, tmp := 0, 0
	for ; i <= j; i++ {
		if seats[i] == 0 {
			tmp++
		} else {
			if tmp > length {
				length = tmp
			}
			tmp = 0
		}
	}
	if res < (length+1)/2 {
		return (length + 1) / 2
	}
	return res
}
