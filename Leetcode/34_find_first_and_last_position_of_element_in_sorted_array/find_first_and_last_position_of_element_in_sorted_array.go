package main

import "fmt"

func main() {
	res := mergeKLists([]*ListNode{
		buildList([]int{1, 4, 5}),
		buildList([]int{1, 3, 4}),
		buildList([]int{2, 6}),
	})
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

func mergeKLists(lists []*ListNode) *ListNode {
	queue := buildPriorityQueue(lists)
	result := &ListNode{}
	head := result
	if len(queue) <= 0 {
		return nil
	}
	for {
		p := popQueue(&queue)
		if p == nil {
			return head.Next
		}
		result.Next = p
		result = result.Next
	}
}

func buildPriorityQueue(lists []*ListNode) []*ListNode {
	queue := make([]*ListNode, 0)
	for _, n := range lists {
		if n == nil {
			continue
		}
		queue = append(queue, n)
		idx := len(queue) - 1
		for idx > 0 && queue[idx].Val < queue[parent(idx)].Val {
			tmp := queue[idx]
			queue[idx] = queue[parent(idx)]
			queue[parent(idx)] = tmp
			idx = parent(idx)
		}
	}
	return queue
}

func parent(x int) int {
	return (x - 1) >> 1
}

func left(x int) int {
	return x*2 + 1
}

func right(x int) int {
	return x*2 + 2
}

func popQueue(q *[]*ListNode) *ListNode {
	if len(*q) <= 0 {
		return nil
	}
	res := (*q)[0]
	(*q)[0] = (*q)[0].Next
	if (*q)[0] == nil {
		(*q)[0] = (*q)[len((*q))-1]
		(*q) = (*q)[:len((*q))-1]
	}
	shiftDown(q)
	return res
}

func shiftDown(q *[]*ListNode) {
	root := 0
	for len(*q) > 0 && (*q)[root] != nil {
		if (left(root) < len((*q)) && (*q)[root].Val > (*q)[left(root)].Val) && (right(root) < len((*q)) && (*q)[root].Val > (*q)[right(root)].Val) {
			if (*q)[left(root)].Val > (*q)[right(root)].Val {
				tmp := (*q)[right(root)]
				(*q)[right(root)] = (*q)[root]
				(*q)[root] = tmp
				root = right(root)
			} else {
				tmp := (*q)[left(root)]
				(*q)[left(root)] = (*q)[root]
				(*q)[root] = tmp
				root = left(root)
			}
			continue
		}
		if left(root) < len((*q)) && (*q)[root].Val > (*q)[left(root)].Val {
			tmp := (*q)[left(root)]
			(*q)[left(root)] = (*q)[root]
			(*q)[root] = tmp
			root = left(root)
			continue
		}
		if right(root) < len((*q)) && (*q)[root].Val > (*q)[right(root)].Val {
			tmp := (*q)[right(root)]
			(*q)[right(root)] = (*q)[root]
			(*q)[root] = tmp
			root = right(root)
			continue
		}
		break
	}
}
