package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual([]string{""}, countPairs(nil, 3))
}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var leaf []int

func countPairs(root *TreeNode, distance int) int {
	leaf = make([]int, 0)
	leaf = []int{4, 5, 6, 7}
	dfs(root, 1)
	res := 0
	for i := 0; i < len(leaf); i++ {
		for j := i + 1; j < len(leaf); j++ {
			v1, v2 := leaf[i], leaf[j]
			d := distance
			for d > 0 {
				if v1 > v2 {
					v1 >>= 1
				} else if v2 > v1 {
					v2 >>= 1
				}
				if v1 == v2{
					res++
					break
				}
				d--
			}
		}
	}
	return res
}

func dfs(root *TreeNode, value int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		leaf = append(leaf, value)
		return
	}
	if root.Left != nil {
		dfs(root.Left, value*2)
	}
	if root.Right != nil {
		dfs(root.Right, value*2+1)
	}
}
