package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
 * 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：2.3 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func splitBST(root *TreeNode, V int) []*TreeNode {
	res := []*TreeNode{nil, nil}
	if root == nil {
		return res
	}
	if root.Val <= V{
		res = splitBST(root.Right, V)
		root.Right = res[0]
		res[0] = root
		return res
	}else{
		res = splitBST(root.Left, V)
		root.Left = res[1]
		res[1] = root
		return res
	}
}
