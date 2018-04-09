package data

//boardSize is the size of one side of the board, in squares.
const boardSize = 5

// Board represents a board.
type Board struct {
	Squares [][]int
}

func makeBoard() Board {
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
		for j := range board[i] {
			board[i][j] = 0
		}
	}
	return Board{board}
}
