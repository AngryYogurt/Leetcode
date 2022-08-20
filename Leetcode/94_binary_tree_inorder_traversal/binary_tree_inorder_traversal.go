package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	return stackInorder(root)
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

// 普通遍历
func stackInorder(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	curr := root
	for curr != nil {
		if curr.Left != nil {
			stack = append(stack, curr)
			curr = curr.Left
			continue
		}
		res = append(res, curr.Val)
		if curr.Right != nil {
			curr = curr.Right
			continue
		}
		if len(stack) > 0 {
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			continue
		}
		break
	}
	return res
}
