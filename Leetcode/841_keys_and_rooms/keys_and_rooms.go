package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(true, canVisitAllRooms([][]int{
		{1}, {2}, {3}, {},
	}))
	tools.AssertObjectEqual(false, canVisitAllRooms([][]int{
		{1, 3}, {3, 0, 1}, {2}, {0},
	}))
}
/*
 * 执行用时：8 ms, 在所有 Go 提交中击败了71.43%的用户
 * 内存消耗：4.4 MB, 在所有 Go 提交中击败了23.53%的用户
 */
func canVisitAllRooms(rooms [][]int) bool {
	q := make([]int, 0, len(rooms))
	q = append(q, rooms[0]...)
	m := make(map[int]bool)
	m[0] = true
	idx := 0
	for idx < len(q) {
		if !m[q[idx]] {
			q = append(q, rooms[q[idx]]...)
			m[q[idx]] = true
		}
		idx++
	}
	if len(m) < len(rooms) {
		return false
	}
	return true
}
