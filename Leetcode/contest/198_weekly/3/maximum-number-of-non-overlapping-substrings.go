package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual([]string{""}, maxNumOfSubstrings("ababa" ))
	tools.AssertObjectEqual([]string{""}, maxNumOfSubstrings("abbaccd"))
	tools.AssertObjectEqual([]string{""}, maxNumOfSubstrings("abab"))
}
//TODO 5466
func maxNumOfSubstrings(s string) []string {
	if len(s) <= 0 {
		return []string{}
	}
	if len(s) == 1 {
		return []string{s}
	}

	tmp := make([][]int, 26)
	for i, _ := range tmp {
		tmp[i] = make([]int, 2)
		tmp[i][0] = -1
		tmp[i][1] = -1
	}
	for i := 0; i < len(s); i++ {
		v := s[i]
		if tmp[v-'a'][0] == -1 {
			tmp[v-'a'][0] = i
		}
	}
	for i := len(s) - 1; i >= 1; i-- {
		v := s[i]
		if tmp[v-'a'][1] == -1 {
			tmp[v-'a'][1] = i
		}
	}
	tmp2 := make([][]int, 0)
	for _, v := range tmp {
		if v[0] == -1 {
			continue
		}
		tmp2 = append(tmp2, v)
	}
	tmp = tmp2
	// sort
	for i := 0; i < len(tmp); i++ {
		m := i
		for j := i + 1; j < len(tmp); j++ {
			if tmp[j][0] < tmp[m][0] {
				m = j
			}
		}
		tmp[i], tmp[m] = tmp[m], tmp[i]
	}
	del := make(map[int]bool)
	for i := 0; i < len(tmp); i++ {
		for j := i + 1; j < len(tmp); j++ {
			if tmp[j][0] < tmp[i][1] && tmp[j][1] > tmp[i][1]{
				tmp[i] = []int{tmp[i][0], tmp[j][1]}
				del[j] = true
			}
		}
	}
	tmp2 = make([][]int, 0)
	if len(del) > 0 {
		for i := 0; i < len(tmp); i++ {
			if del[i] {
				continue
			}
			tmp2 = append(tmp2, tmp[i])
		}
		tmp = tmp2
	}

	r, max := travel(tmp, 0, [][]int{}, 0)
	for c := 1; c < len(tmp); c++ {
		tmpR, tmpMax := travel(tmp, c, [][]int{}, 0)
		if len(tmpR) == len(r) && tmpMax < max {
			r, max = tmpR, tmpMax
		} else if len(tmpR) > len(r) {
			r, max = tmpR, tmpMax
		}
	}
	result := make([]string, len(r))
	for i, v := range r {
		result[i] = s[v[0] : v[1]+1]
	}
	return result
}

func travel(arr [][]int, curr int, currRes [][]int, currMax int) (res [][]int, maxL int) {
	res = append([][]int{}, currRes...)
	currA := arr[curr]
	if currA[1]-currA[0] > currMax {
		currMax = currA[1] - currA[0]+1
	}
	res = append(res, currA)
	for i := curr + 1; i < len(arr); i++ {
		if v := arr[i]; v[0] <= res[len(res)-1][1] {
			continue
		}
		tmpRes, tmpMax := travel(arr, i, res, currMax)
		if len(tmpRes) == len(res) && tmpMax < currMax {
			res, currMax = tmpRes, tmpMax
		} else if len(tmpRes) > len(res) {
			res, currMax = tmpRes, tmpMax
		}
	}
	return res, currMax
}
