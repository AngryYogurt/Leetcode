/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	return recursiveInorder(root)
}

// morris 遍历
func morrisInorder(root *TreeNode) []int {
	res := make([]int, 0)
	curr := root
	for curr != nil {
		// 左子树非空，先遍历
		if curr.Left != nil {
			tmp := curr.Left
			// 找到左子树中序遍历根节点的前驱（直接找左子树根节点的最右）
			for tmp.Right != nil && tmp.Right != curr {
				tmp = tmp.Right
			}
			// 判断是否已挂上辅助线
			// 未挂上辅助线，左子树未遍历
			if tmp.Right == nil {
				tmp.Right = curr // 挂上辅助线，左子树遍历完成后可以回到当前节点
				curr = curr.Left // 开始遍历左子树
				continue         // 开始遍历
			} else {
				tmp.Right = nil // 已挂上辅助线，说明左子树遍历完成，删除辅助线
			}
		}
		// 左子树遍历完成 或 左子树为空，结果加入当前节点，开始遍历右子树
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}

// 栈遍历
func stackInorder(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	curr := root
	isPop := false
	for curr != nil {
		if curr.Left != nil && !isPop {
			stack = append(stack, curr)
			curr = curr.Left
			isPop = false
			continue
		}
		res = append(res, curr.Val)
		if curr.Right != nil {
			curr = curr.Right
			isPop = false
			continue
		}
		if len(stack) > 0 {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			isPop = true
			continue
		}
		break
	}
	return res
}

// 栈遍历2
func stackInorder2(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}

// 递归遍历
func recursiveInorder(root *TreeNode) []int {
	res := make([]int, 0)
	var recur func(*TreeNode)
	recur = func(curr *TreeNode) {
		if curr == nil {
			return
		}
		recur(curr.Left)
		res = append(res, curr.Val)
		recur(curr.Right)
	}
	recur(root)
	return res
}

// @lc code=end

