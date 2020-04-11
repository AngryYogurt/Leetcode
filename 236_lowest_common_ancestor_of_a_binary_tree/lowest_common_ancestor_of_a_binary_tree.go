package main

import "github.com/AngryYogurt/leetcode/tools"

func main() {
	params := []*TreeNode{}
	v := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	for idx, _ := range v {
		if v[idx] == -1 {
			params = append(params, nil)
		} else {
			params = append(params, &TreeNode{
				Val: v[idx],
			})
		}
	}
	for idx, _ := range params {
		if ni := idx*2 + 1; ni < len(params) {
			params[idx].Left = params[ni]
		}
		if ni := idx*2 + 2; ni < len(params) {
			params[idx].Right = params[ni]
		}
	}
	/*
	        3
	    5      1
	  6   2   0 8
	     7 4
	 */
	tools.AssertObjectEqual(params[1], lowestCommonAncestor(params[0], params[1], params[10]))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
/*
 * 执行用时 :12 ms, 在所有 Go 提交中击败了92.68%的用户
 * 内存消耗 :7.2 MB, 在所有 Go 提交中击败了25.31%的用户
 */
var ans1 *TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    ans1 = nil
    travel(root, p, q)
    return ans1
}
func travel(root, l, r *TreeNode) int {
    if ans1 != nil {
        return 0
    }
    curr := 0
    if root == nil {
        return 0
    }
    if root == l || root == r {
        curr = 1
    }
    var exL, exR int
    exL = travel(root.Left, l, r)
    exR = travel(root.Right, l, r)
    if exL+exR+curr >= 2 {
        ans1 = root
        return 1
    }
    return exL + exR + curr
}
