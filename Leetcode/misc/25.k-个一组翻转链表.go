/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	first := head
	var lastT *ListNode
	var res *ListNode
	for {
		ptr := first
		for i := 0; i < k-1; i++ {
			if ptr.Next == nil {
				break
			}
			ptr = ptr.Next
		}
		nextFirst := ptr.Next
		ptr.Next = nil
		h, t := reverseFull(first)
		if res == nil {
			res = h
		}
		if lastT != nil {
			lastT.Next = h
		}
		lastT = t
		first = nextFirst
	}
	return res
}

func reverseFull(head *ListNode) (h, t *ListNode) {
	t = head
	var last *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = last
		last = head
		head = tmp
	}
	return last, t
}

// @lc code=end

