package main

import "fmt"

func main() {
	fmt.Println(containsCycle([][]byte{
		{'b'},
		{'b'},
	}))
	fmt.Println(containsCycle([][]byte{
		{'a', 'b', 'b'},
		{'b', 'z', 'b'},
		{'b', 'b', 'a'},
	}))
	fmt.Println(containsCycle([][]byte{
		{'c', 'a', 'd'},
		{'a', 'a', 'a'},
		{'a', 'a', 'd'},
		{'a', 'c', 'd'},
		{'a', 'b', 'c'},
	}))
	fmt.Println(containsCycle([][]byte{
		{'a', 'a', 'a', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'b', 'b', 'a'},
		{'a', 'a', 'a', 'a'},
	}))
}

func containsCycle(grid [][]byte) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '!' {
				continue
			}
			conn := make([][]int, 0)
			if dfs(grid, i, j, grid[i][j], i, j, &conn, -1, -1, 0) {
				return true
			}
			for k := 0; k < len(conn); k++ {
				grid[conn[k][0]][conn[k][1]] = '!'
			}
		}
	}
	return false
}

func dfs(g [][]byte, i, j int, v byte, si, sj int, conn *[][]int, li, lj int, l int) bool {
	if i == si && j == sj && l > 3 {
		return true
	}
	*conn = append(*conn, []int{i, j})
	if g[i][j] == v+'z' {
		return true
	}else{
		g[i][j] = v+'z'
	}
	if ni, nj := i+1, j; i+1 < len(g) && ni != li && (g[ni][nj] == v || g[ni][nj] == v+'z') {
		if dfs(g, ni, nj, v, si, sj, conn, i, j, l+1) {
			return true
		}
	}
	if ni, nj := i, j+1; j+1 < len(g[0]) && nj != lj && (g[ni][nj] == v || g[ni][nj] == v+'z') {
		if dfs(g, ni, nj, v, si, sj, conn, i, j, l+1) {
			return true
		}
	}
	if ni, nj := i-1, j; ni >= 0 && ni != li && (g[ni][nj] == v || g[ni][nj] == v+'z') {
		if dfs(g, ni, nj, v, si, sj, conn, i, j, l+1) {
			return true
		}
	}
	if ni, nj := i, j-1; nj >= 0 && nj != lj && (g[ni][nj] == v || g[ni][nj] == v+'z') {
		if dfs(g, ni, nj, v, si, sj, conn, i, j, l+1) {
			return true
		}
	}
	return false
}
