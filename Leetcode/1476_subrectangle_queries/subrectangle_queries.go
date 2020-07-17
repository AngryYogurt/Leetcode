package main

func main() {

}
/*
 * 执行用时：56 ms, 在所有 Go 提交中击败了92.22%的用户
 * 内存消耗：7.2 MB, 在所有 Go 提交中击败了100.00%的用户
 */
type SubrectangleQueries struct {
	UpdateLog [][5]int
	Rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		UpdateLog: make([][5]int, 0),
		Rectangle: rectangle,
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	t := sort(row1, col1, row2, col2)
	this.UpdateLog = append(this.UpdateLog, [5]int{t[0], t[1], t[2], t[3], newValue})
}

func sort(row1 int, col1 int, row2 int, col2 int) [4]int {
	return [4]int{min(row1, row2), min(col1, col2), max(row1, row2), max(col1, col2)}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	for i := len(this.UpdateLog) - 1; i >= 0; i-- {
		if l := this.UpdateLog[i]; row >= l[0] && row <= l[2] && col >= l[1] && col <= l[3] {
			return l[4]
		}
	}
	return this.Rectangle[row][col]
}
