package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual([]int{}, countSubTrees(7, [][]int{
		{0, 1},
		{0, 2},
		{1, 4},
		{1, 5},
		{2, 3},
		{2, 6}}, "abaedcd"))
}
/*
 * 执行用时：1132 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗：84.4 MB, 在所有 Go 提交中击败了100.00%的用户
 */
type N struct {
	L []*N
	V int
}

var re []int

func countSubTrees(n int, edges [][]int, labels string) []int {
	if n == 1 {
		return []int{1}
	}
	m := make(map[int]*N)
	for _, e := range edges {
		if _, ok := m[e[0]]; !ok {
			m[e[0]] = &N{
				L: make([]*N, 0),
				V: e[0],
			}
		}
		if _, ok := m[e[1]]; !ok {
			m[e[1]] = &N{
				L: make([]*N, 0),
				V: e[1],
			}
		}
		m[e[0]].L = append(m[e[0]].L, m[e[1]])
		m[e[1]].L = append(m[e[1]].L, m[e[0]])
	}

	re = make([]int, n)
	dfs(m[0], labels, nil)
	return re
}

func dfs(root *N, labels string, last *N) map[int]int {
	res := make(map[int]int)
	res[int(labels[root.V]-'a')] = 1
	if len(root.L) <= 1 && root.V > 0 {
		re[root.V] = 1
		return res
	}
	for _, n := range root.L {
		if n.V == root.V || (last != nil && last.V == n.V) {
			continue
		}
		m := dfs(n, labels, root)
		for k, v := range m {
			res[k] += v
		}
	}
	re[root.V] = res[int(labels[root.V]-'a')]
	return res
}
