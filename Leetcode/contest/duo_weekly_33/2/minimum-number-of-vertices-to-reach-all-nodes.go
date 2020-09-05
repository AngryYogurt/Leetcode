package main

func main() {
}
func findSmallestSetOfVertices(n int, edges [][]int) []int {
	res := make([]int, n)
	for i := 0; i < len(edges); i++ {
		res[edges[i][1]] = 1
	}
	rr := make([]int, 0)
	for i := 0; i < len(res); i++ {
		if res[i] == 0 {
			rr = append(rr, i)
		}
	}
	return rr
}
