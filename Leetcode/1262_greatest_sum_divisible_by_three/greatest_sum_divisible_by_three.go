package main

import (
	"github.com/AngryYogurt/ProgrammingProblem/tools"
	"math"
)

func main() {
	tools.AssertObjectEqual(18, maxSumDivThree([]int{3, 6, 5, 1, 8}))
}

func maxSumDivThree(nums []int) int {
	minOne, minOne2, minTwo, minTwo2 := math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64
	sum := 0
	for _, v := range nums {
		sum += v
		switch v % 3 {
		case 1:
			minOne, minOne2 = min2(minOne, minOne2, v)
		case 2:
			minTwo, minTwo2 = min2(minTwo, minTwo2, v)
		}
	}
	minOne, minOne2, minTwo, minTwo2 = fixNum(minOne), fixNum(minOne2), fixNum(minTwo), fixNum(minTwo2)
	if minOne2 > 0 && minTwo2 > 0 {
		if minOne > 0 && minTwo > 0 {
			return sum
		} else {
			sum -= minOne
			sum -= minTwo
		}
	} else {
		sum -= minTwo
		sum -= minTwo2
		sum -= minOne
		sum -= minOne2
	}
	return sum
}

func min2(v1, v2, v3 int) (int, int) {
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	if v2 > v3 {
		v2, v3 = v3, v2
	}
	if v1 > v2 {
		v1, v2 = v2, v1
	}
	return v1, v2
}

func fixNum(value int) int {
	if value == math.MaxInt64 {
		return 0
	}
	return value
}
