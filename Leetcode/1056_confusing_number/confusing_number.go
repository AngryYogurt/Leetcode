package main

import "github.com/AngryYogurt/ProgrammingProblem/tools"

func main() {
	tools.AssertObjectEqual(false, confusingNumber(11))
}

func confusingNumber(n int) bool {
	new := 0
	for n > 0 {
		m := n % 10
		new = new*10 + m
		switch m {
		case 2, 3, 4, 5, 7:
			return false
		}
		n /= 10
	}
	return new != n
}
