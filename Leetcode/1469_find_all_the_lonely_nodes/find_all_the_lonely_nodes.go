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

/*
 * 执行用时：8 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：6.1 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func getLonelyNodes(root *TreeNode) []int {
	result = make([]int, 0)
	dfs(root)
	return result
}

var result []int

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right != nil {
		result = append(result, root.Right.Val)
	} else if root.Left != nil && root.Right == nil {
		result = append(result, root.Left.Val)
	}
	dfs(root.Left)
	dfs(root.Right)
	return
}
