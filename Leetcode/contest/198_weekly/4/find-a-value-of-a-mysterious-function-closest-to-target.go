package main

import "github.com/AngryYogurt/leetcode/tools"

func main() {
	tools.AssertObjectEqual(2, closestToTarget([]int{9, 12, 3, 7, 15}, 5))
	tools.AssertObjectEqual(999999, closestToTarget([]int{1000000, 1000000, 1000000}, 1))
	tools.AssertObjectEqual(0, closestToTarget([]int{1, 2, 4, 8, 16}, 0))
}

//TODO 5467
type node struct {
	Start, End   int
	V            int
	L, R, Parent *node
}

func closestToTarget(arr []int, target int) int {
	res := 0
	if len(arr) == 0 {
		return 0
	}
	r := build(arr, 0, len(arr)-1, nil)
	for l := 0; l < len(arr); l++ {
		for r := l; r < len(arr); r++ {

		}
	}
	return res
}

func build(arr []int, s, e int, p *node) *node {
	root := &node{
		Start:  s,
		End:    e,
		V:      0,
		Parent: p,
	}
	if s == e {
		root.V = arr[s]
		return root
	}
	root.L = build(arr, s, (e+s)/2, root)
	root.R = build(arr, (s+e)/2+1, e, root)
	root.V = root.L.V & root.R.V
	return root
}
