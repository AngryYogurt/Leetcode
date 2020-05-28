package main

func main() {
	buildTree([]int{1,2,3}, []int{3,2,1})
	buildTree([]int{3,9,20,15,7}, []int{9,3,15,20,7})
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
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :4 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0 {
		return nil
	}
	return f(preorder, inorder)
}

func f(preorder []int, inorder []int) *TreeNode {
	if len(preorder) <= 0{
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{
			Val:   preorder[0],
			Left:  nil,
			Right: nil,
		}
	}
	rootVal := preorder[0]
	rtIdx := 0
	for i, _ := range inorder {
		if inorder[i] == rootVal {
			rtIdx = i
			break
		}
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  f(preorder[1:rtIdx+1], inorder[:rtIdx]),
		Right: f(preorder[rtIdx+1:], inorder[rtIdx+1:]),
	}
}
