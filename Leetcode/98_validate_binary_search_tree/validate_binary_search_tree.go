package main

import "github.com/AngryYogurt/leetcode/tools"

func main() {
	tools.AssertObjectEqual(false, isValidBST(&TreeNode{
		Val: 0,
		Left: nil,
		Right: nil,
	}))
	tools.AssertObjectEqual(false, isValidBST(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
 * 执行用时 :8 ms, 在所有 Go 提交中击败了79.20%的用户
 * 内存消耗 :5.5 MB, 在所有 Go 提交中击败了80.00%的用户
 */

var pre = -int(^uint(0)>>1)-1

func isValidBST(root *TreeNode) bool {
	pre = -int(^uint(0)>>1)-1
	return Inorder(root)
}

func Inorder(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if !Inorder(root.Left) {
		return false
	}
	if pre >= root.Val {
		return false
	} else {
		pre = root.Val
	}
	if !Inorder(root.Right) {
		return false
	}
	return true
}
