/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	res := sum
	carry := 0
	for l1 != nil && l2 != nil {
		sum.Next = &ListNode{}
		sum = sum.Next
		sum.Val = l1.Val + l2.Val + carry
		carry = 0
		if sum.Val > 9 {
			sum.Val -= 10
			carry = 1
		}
		l1, l2 = l1.Next, l2.Next
	}
	rest := l1
	if l1 == nil {
		rest = l2
	}
	for carry > 0 || rest != nil {
		sum.Next = &ListNode{}
		sum = sum.Next
		sum.Val = carry
		if rest != nil {
			sum.Val += rest.Val
			rest = rest.Next
		}
		carry = 0
		if sum.Val > 9 {
			sum.Val -= 10
			carry = 1
		}
	}
	return res.Next
}

// @lc code=end

