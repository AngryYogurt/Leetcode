/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	h1, h2 := list1, list2
	res := &ListNode{}
	head := res
	for h1 != nil || h2 != nil {
		if h1 == nil {
			head.Next = h2
			break
		}
		if h2 == nil {
			head.Next = h1
			break
		}
		if h1.Val > h2.Val {
			head.Next = h2
			h2 = h2.Next
		} else {
			head.Next = h1
			h1 = h1.Next
		}
		head = head.Next
	}
	return res.Next
}

// @lc code=end

