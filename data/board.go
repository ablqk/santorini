package data

//boardSize is the size of one side of the board, in squares.
const boardSize = 5

type Square struct {
	height int
}

type Board struct {
	Squares [][]Square
}

func makeBoard() Board {
	board := make([][]Square, boardSize)
	for i := range board {
		board[i] = make([]Square, boardSize)
		for j := range board[i] {
			board[i][j] = Square{}
		}
	}
	return Board{board}
}
