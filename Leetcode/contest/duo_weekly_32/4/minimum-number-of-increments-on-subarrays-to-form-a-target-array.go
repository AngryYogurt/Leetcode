package main

import "fmt"

func main() {
	fmt.Println(longestAwesome("51224"))
}

func longestAwesome(s string) int {
	total := map[int]int{}
	for i := 0; i < len(s); i++ {
		total[int(s[i]-'0')]++
	}
	oddC := make([]bool, 10)
	to := 0
	for k, v := range total {
		if v%2 == 1 {
			oddC[k] = true
			to++
		} else {
			oddC[k] = false
		}
	}
	tmpOC := append([]bool{}, oddC...)
	max := 0
	tmpTo := to
	for i := 0; i < len(s); i++ {
		if len(s)-i <= max {
			return max
		}
		if i > 0 {
			oddC[int(s[i-1]-'0')] = !oddC[int(s[i-1]-'0')]
			if oddC[int(s[i-1]-'0')] {
				tmpTo++
			} else {
				tmpTo--
			}
		}
		to = tmpTo
		oddC = append([]bool{}, tmpOC...)
		for j := len(s) - 1; j >= 0; j-- {
			if j-i+1 <= max {
				break
			}
			if j < len(s)-1 {
				oddC[int(s[j+1]-'0')] = !oddC[int(s[j+1]-'0')]
				if oddC[int(s[j+1]-'0')] {
					to++
				} else {
					to--
				}
			}
			if to <= 1 {
				max = j - i + 1
				break
			}
		}
	}
	return max
}
