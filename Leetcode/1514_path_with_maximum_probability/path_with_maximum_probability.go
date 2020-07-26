package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main(){
	tools.AssertObjectEqual(0.25, maxProbability(3, [][]int{{0,1}, {1,2},{0,2}}, []float64{0.5,0.5,0.2}, 0, 2))
	tools.AssertObjectEqual(0.3, maxProbability(3, [][]int{{0,1}, {1,2},{0,2}}, []float64{0.5,0.5,0.3}, 0, 2))
	tools.AssertObjectEqual(0, maxProbability(3, [][]int{{0,1}}, []float64{0.5}, 0, 2))
}

/*
 * 执行用时：228 ms, 在所有 Go 提交中击败了31.67%的用户
 * 内存消耗：9.9 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	edge := make(map[int]map[int]float64)
	for i, e := range edges{
		if _, ok := edge[e[0]];!ok{
			edge[e[0]] = make(map[int]float64, 0)
		}
		if _, ok := edge[e[1]];!ok{
			edge[e[1]] = make(map[int]float64, 0)
		}
		edge[e[0]][e[1]] = succProb[i]
		edge[e[1]][e[0]] = succProb[i]
	}
	q := make([]int, 0)
	prob := make(map[int]float64)
	q = append(q, start)
	prob[start] = 1
	i := 0
	for i < len(q){
		curr := q[i]
		for k, v:= range edge[curr]{
			if p := prob[curr] * v; p > prob[k]{
				prob[k] = p
				if k != end{
					q = append(q, k)
				}
			}
		}
		i++
	}
	return prob[end]
}