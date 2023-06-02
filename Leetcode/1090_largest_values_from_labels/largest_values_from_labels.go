package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	//tools.AssertObjectEqual(12, largestValsFromLabels([]int{5, 4, 3, 2, 1}, []int{1, 3, 3, 3, 2}, 3, 2))
	//tools.AssertObjectEqual(6, largestValsFromLabels([]int{2, 6, 1, 2, 6}, []int{2, 2, 2, 1, 1}, 1, 1))
	tools.AssertObjectEqual(20, largestValsFromLabels([]int{0, 6, 7, 5, 7}, []int{2, 0, 2, 0, 2}, 3, 4))
}

func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	qsort(values, labels, 0, len(values)-1)
	cntMap := make(map[int][]int)
	for idx := range values {
		v := values[idx]
		if _, ok := cntMap[v]; !ok {
			cntMap[v] = make([]int, 0)
		}
		cntMap[v] = append(cntMap[v], labels[idx])
	}
	labelSet := make(map[int]int)
	sum := 0
	for idx := range values {
		if idx > 0 && values[idx] == values[idx-1] {
			continue
		}
		v := values[idx]
		for _, label := range cntMap[v] {
			if labelSet[label] < useLimit {
				labelSet[label]++
				sum += v
				numWanted--
				if numWanted <= 0 {
					break
				}
			}
		}
		if numWanted <= 0 {
			break
		}
	}
	return sum
}

func qsort(values, labels []int, start, end int) {
	if start >= end {
		return
	}
	l, r := start, end
	key := values[l]
	for l <= r {
		for values[l] > key {
			l++
		}
		for values[r] < key {
			r--
		}
		if l <= r {
			values[l], values[r] = values[r], values[l]
			labels[l], labels[r] = labels[r], labels[l]
			l++
			r--
		}
	}
	//values[start], values[l] = values[l], values[start]
	//labels[start], labels[l] = labels[l], labels[start]
	qsort(values, labels, start, r)
	qsort(values, labels, l, end)
}
