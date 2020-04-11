package main

func main() {

}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	return recursive(head)
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.9 MB, 在所有 Go 提交中击败了7.11%的用户
 */
func recursive(head *ListNode) *ListNode {
	return rec(head, nil)
}

func rec(head, last *ListNode) *ListNode {
	if head == nil {
		return last
	}
	n := head.Next
	head.Next = last
	return rec(n, head)
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.9 MB, 在所有 Go 提交中击败了15.20%的用户
 */
func recursive2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := recursive2(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.6 MB, 在所有 Go 提交中击败了69.87%的用户
 */
func iterative(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var last, next *ListNode = nil, head
	for head != nil {
		next = head.Next
		head.Next = last
		last = head
		head = next
	}
	return last
}
