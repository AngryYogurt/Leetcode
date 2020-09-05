package main

func main() {

}

func diagonalSum(mat [][]int) int {
	if len(mat) == 0{
		return 0
	}
	sum := 0
	x, y := 0, 0
	for i := 0; i < len(mat); i++ {
		sum += mat[x][y]
		x++
		y++
	}

	x, y = 0, len(mat)-1
	for i := 0; i < len(mat); i++ {
		sum += mat[x][y]
		x++
		y--
	}
	if len(mat)%2 == 1 {
		sum -= mat[len(mat)/2][len(mat)/2]
	}
	return sum
}