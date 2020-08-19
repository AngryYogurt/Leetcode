package main

func countOdds(low int, high int) int {
	if high%2 == 1 || low%2 == 1 {
		return (high-low)/2 + 1
	}
	return (high - low) / 2
}
