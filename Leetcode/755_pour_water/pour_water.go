package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual([]int{4, 4, 4, 4, 3, 3, 3, 3, 3, 4, 3, 2, 1}, pourWater([]int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}, 10, 2))
	tools.AssertObjectEqual([]int{4, 4, 4}, pourWater([]int{3, 1, 3}, 5, 1))
	tools.AssertObjectEqual([]int{2, 3, 3, 4}, pourWater([]int{1, 2, 3, 4}, 2, 2))
	tools.AssertObjectEqual([]int{2, 2, 2, 3, 2, 2, 2}, pourWater([]int{2, 1, 1, 2, 1, 2, 2}, 4, 3))
}

/*
 * 执行用时 :0 ms, 在所有 Go 提交中击败了100.00%的用户
 * 内存消耗 :2.1 MB, 在所有 Go 提交中击败了100.00%的用户
 */
func pourWater(heights []int, V int, K int) []int {
	max, sum := heights[0], 0
	for _, v := range heights {
		if v > max {
			max = v
		}
		sum += v
	}
	if base := max * len(heights); V+sum >= base {
		V = V - (base - sum)
		extra := V / len(heights)
		V = V - extra*len(heights)
		for idx, _ := range heights {
			heights[idx] = max + extra
		}
	}
	for V > 0 {
		V--
		idx, target := K, K
		for idx > 0 {
			if heights[idx] > heights[idx-1] {
				target = idx - 1
			} else if heights[idx] < heights[idx-1] {
				break
			}
			idx--
		}
		if heights[target]+1 > heights[K] {
			target = K
		}
		if target == K {
			idx = K
			for idx < len(heights)-1 {
				if heights[idx] > heights[idx+1] {
					target = idx + 1
				} else if heights[idx] < heights[idx+1] {
					break
				}
				idx++
			}
		}
		if heights[target]+1 > heights[K] {
			target = K
		}
		heights[target]++
	}
	return heights
}
