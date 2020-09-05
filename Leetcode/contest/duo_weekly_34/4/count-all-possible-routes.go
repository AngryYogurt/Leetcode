package main

import "fmt"

func main() {
	fmt.Println(countRoutes([]int{1, 2, 3}, 0, 2, 40))
	fmt.Println(countRoutes([]int{4, 3, 1}, 1, 0, 6))
}

func countRoutes(locations []int, start int, finish int, fuel int) int {
	st = make([][]int, len(locations))
	for i, _ := range st {
		st[i] = make([]int, fuel+1)
	}
	st[start][fuel] = 1
	dp(locations, start, fuel)
	total := 0
	for i := 0; i < fuel+1; i++ {
		total += st[finish][i]
		if total > 1000000007 {
			total -= 1000000007
		}
	}
	return total
}

var st [][]int

func dp(l []int, curr int, fuel int) {
	for f := fuel; f >= 0; f-- {
		for i := 0; i < len(st); i++ {
			for j := 0; j < len(st); j++ {
				if j == i {
					continue
				}
				fue := l[i] - l[j]
				if fue < 0 {
					fue = -fue
				}
				if fue+f > fuel {
					continue
				}
				st[i][f] += st[j][f+fue]
				if st[i][f] > 1000000007 {
					st[i][f] -= 1000000007
				}
			}
		}
	}
}
