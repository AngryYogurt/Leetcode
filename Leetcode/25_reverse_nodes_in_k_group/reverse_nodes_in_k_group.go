package main

import "fmt"

func main() {
	res := reverseKGroup(buildList([]int{1, 2, 3, 4, 5}), 2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func buildList(src []int) *ListNode {
	res := &ListNode{}
	head := res
	for _, v := range src {
		res.Next = &ListNode{
			Val: v,
		}
		res = res.Next
	}
	return head.Next
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	first := head
	var lastT *ListNode
	var res *ListNode
	breakLast := false
	for first != nil {
		ptr := first
		for i := 0; i < k-1; i++ {
			if ptr.Next == nil {
				breakLast = true
				break
			}
			ptr = ptr.Next
		}
		if breakLast {
			lastT.Next = first
			return res
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
