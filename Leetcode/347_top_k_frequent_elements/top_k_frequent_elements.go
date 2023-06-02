package main

func main() {

}

func topKFrequent(nums []int, k int) []int {
	cntStat := make(map[int]int)
	for _, v := range nums {
		cntStat[v]++
	}
	stats := make([][2]int, 0)
	for val, cnt := range cntStat {
		stats = append(stats, [2]int{val, cnt})
	}
	qsort(stats, 0, len(stats)-1)
	result := make([]int, 0)
	for i := 0; i < k; i++ {
		result = append(result, stats[i][0])
	}
	return result
}

func qsort(src [][2]int, s, e int) {
	l, r := s, e
	if l > r {
		return
	}
	for l < r {
		for src[l][1] >= src[s][1] {
			l++
		}
		for src[r][1] <= src[s][1] {
			r--
		}
		if l < r {
			src[l], src[r] = src[r], src[l]
			l++
			r--
		}
	}
	qsort(src, s, r)
	qsort(src, l, e)
}
