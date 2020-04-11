package main

func main(){
	
}


type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}
/*
 * 执行用时 :24 ms, 在所有 Go 提交中击败了82.59%的用户
 * 内存消耗 :6.8 MB, 在所有 Go 提交中击败了33.01%的用户
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	queue := make([]*TreeNode,0, 1000)
	if p.Val > q.Val{
		p, q = q, p
	}
	idx := 0
	queue = append(queue, root)
	for ;idx < len(queue);idx++ {
		if queue[idx] == nil{
			continue
		}
		if queue[idx].Val >= p.Val && queue[idx].Val <= q.Val{
			return queue[idx]
		}
		queue = append(queue, []*TreeNode{queue[idx].Left, queue[idx].Right}...)
	}
	return nil
}