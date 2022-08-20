package main

func main() {

}

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minCameraCover(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 0
}

func dfs(root *TreeNode)