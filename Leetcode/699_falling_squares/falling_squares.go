package main

func main() {

}
// TODO
type N struct {
	start  int
	end    int
	height int
	next   *N
}

func fallingSquares(positions [][]int) []int {
	top, idx := 0, 0
	result := make([]int, len(positions))
	head := &N{
		start:  0,
		end:    1000,
		height: 0,
		next:   nil,
	}
	curr := head
	for i := range positions {

		idx++
		curr = head
	}
	return result
}
