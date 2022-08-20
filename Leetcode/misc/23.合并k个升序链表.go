/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	queue := buildPriorityQueue(lists)
	result := &ListNode{}
	for {
		p := popQueue(queue)
		if p == nil {
			return result.Next
		}
		result.Next = p
	}
	return result.Next
}

func buildPriorityQueue(lists []*ListNode) []*ListNode {
	queue := make([]*ListNode, 0)
	for _, n := range lists {
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

func popQueue(queue []*ListNode) *ListNode {
	res := queue[0]
	queue[0] = queue[0].Next
	if queue[0] == nil {
		queue[0] = queue[len(queue)-1]
		queue = queue[:len(queue)-1]
	}
	shiftDown(queue)
	return res
}

func shiftDown(queue []*ListNode) {
	root := 0
	for queue[root] != nil {
		if (left(root) < len(queue) && queue[root].Val > queue[left(root)].Val) && (right(root) < len(queue) && queue[root].Val > queue[right(root)].Val) {
			if queue[left(root)].Val > queue[right(root)].Val {
				tmp := queue[right(root)]
				queue[right(root)] = queue[root]
				queue[root] = tmp
				root = right(root)
			} else {
				tmp := queue[left(root)]
				queue[left(root)] = queue[root]
				queue[root] = tmp
				root = left(root)
			}
			continue
		}
		if left(root) < len(queue) && queue[root].Val > queue[left(root)].Val {
			tmp := queue[right(root)]
			queue[right(root)] = queue[root]
			queue[root] = tmp
			root = right(root)
			continue
		}
		if right(root) < len(queue) && queue[root].Val > queue[right(root)].Val {
			tmp := queue[left(root)]
			queue[left(root)] = queue[root]
			queue[root] = tmp
			root = left(root)
			continue
		}
		break
	}
}

// @lc code=end
