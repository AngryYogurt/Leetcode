package main

import (
	"github.com/AngryYogurt/ProgrammingProblem/tools"
)

func main() {
	o := Constructor(3)
	tools.AssertObjectEqual(0, o.Move(0,0,1))
	tools.AssertObjectEqual(0, o.Move(0,2,2))
	tools.AssertObjectEqual(0, o.Move(2,2,1))
	tools.AssertObjectEqual(0, o.Move(1,1,2))
	tools.AssertObjectEqual(0, o.Move(2,0,1))
	tools.AssertObjectEqual(0, o.Move(1,0,2))
	tools.AssertObjectEqual(1, o.Move(2,1,1))

	o = Constructor(3)
	tools.AssertObjectEqual(0, o.Move(1, 1, 2))
	tools.AssertObjectEqual(0, o.Move(0, 2, 2))
	tools.AssertObjectEqual(2, o.Move(2, 0, 2))
}

/*
 * O(1) 时间解
 */

/*
 * 执行用时 :36 ms, 在所有 Go 提交中击败了94.12%的用户
 * 内存消耗 :9 MB, 在所有 Go 提交中击败了100.00%的用户
 */
type TicTacToe struct {
	n      int
	states [][]int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	s := make([][]int, 2)
	s[0] = make([]int, 2*n+2)
	s[1] = make([]int, 2*n+2)
	return TicTacToe{
		n:      n,
		states: s,
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	this.states[player-1][row]++
	this.states[player-1][this.n+col]++
	if row == col {
		this.states[player-1][len(this.states[player-1])-2]++
	}
	if row+col == this.n-1 {
		this.states[player-1][len(this.states[player-1])-1]++
	}
	/*
	   switch this.n {
		case this.states[player-1][row], this.states[player-1][this.n+col], this.states[player-1][len(this.states[player-1])-2], this.states[player-1][len(this.states[player-1])-1]:
			return player
		}
	 */
	if this.states[player-1][row] == this.n ||
		this.states[player-1][this.n+col] == this.n ||
		this.states[player-1][len(this.states[player-1])-2] == this.n ||
		this.states[player-1][len(this.states[player-1])-1] == this.n {
		return player
	}
	return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
