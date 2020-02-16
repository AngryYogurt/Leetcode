package main

import (
	"github.com/AngryYogurt/leetcode/tools"
	"sort"
)

func main() {
	root1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	}
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}
	tools.AssertObjectEqual(true, twoSumBSTs(root1, root2, 5))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
	t1, t2 := iterateTree(root1), iterateTree(root2)
	if root1 == nil {
		root1 = &TreeNode{
			Val: 0,
		}
	}
	if root2 == nil {
		root2 = &TreeNode{
			Val: 0,
		}
	}
	sort.Slice(t1, func(i, j int) bool { return t1[i] < t1[j] })
	sort.Slice(t2, func(i, j int) bool { return t2[i] < t2[j] })
	idx1, idx2 := 0, len(t2)-1
	for idx1 < len(t1) && idx2 >= 0 {
		if t1[idx1]+t2[idx2] == target {
			return true
		} else if t1[idx1]+t2[idx2] < target {
			idx1++
		} else if t1[idx1]+t2[idx2] > target {
			idx2--
		}
	}
	return false
}

func iterateTree(root *TreeNode) []int {
	result := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	idx := 0
	for idx < len(queue) {
		result = append(result, queue[idx].Val)
		if c := queue[idx].Left; c != nil {
			queue = append(queue, c)
		}
		if c := queue[idx].Right; c != nil {
			queue = append(queue, c)
		}
		idx++
	}
	return result
}
