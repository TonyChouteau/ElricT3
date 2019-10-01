package engine

/*
Enum color value
*/
const (
	NONE   int = 0
	CROSS  int = 1
	CIRCLE int = 2
)

/*
Matrix3x3 declaration
*/
type Matrix3x3 [3][3]int

/*
CopyBoard : deep copy of a board
*/
func CopyBoard(board Matrix3x3) Matrix3x3 {
	newBoard := Matrix3x3{}
	for i := range board {
		for j := range board[i] {
			newBoard[i][j] = board[i][j]
		}
	}
	return newBoard
}

/*
CreateM : create Board matrix 3x3
*/
func CreateM() Matrix3x3 {
	return Matrix3x3{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
}

/*
ListLegal : give the list of legal move
*/
func ListLegal(board Matrix3x3) []int {
	l := []int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == 0 {
				l = append(l, 3*i+j)
			}
		}
	}
	return l
}

/*
Contains : if the slice contains a element
*/
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

/*
HasWon : if someone won
*/
func HasWon(board Matrix3x3, color int) bool {
	if (board[0][0] == color && board[0][1] == color && board[0][2] == color) || (board[1][0] == color && board[1][1] == color && board[1][2] == color) || (board[2][0] == color && board[2][1] == color && board[2][2] == color) || (board[0][0] == color && board[1][0] == color && board[2][0] == color) || (board[0][1] == color && board[1][1] == color && board[2][1] == color) || (board[0][2] == color && board[1][2] == color && board[2][2] == color) || (board[0][0] == color && board[1][1] == color && board[2][2] == color) || (board[0][2] == color && board[1][1] == color && board[2][0] == color) {
		return true
	} else {
		return false
	}
}

/*
IsLegal : is the move legal
*/
func IsLegal(board Matrix3x3, m int) bool {
	l := ListLegal(board)
	for _, e := range l {
		if e == m {
			return true
		}
	}
	return false
}

/*
NextColor : change color
*/
func NextColor(color int) int {
	if color == CROSS {
		return CIRCLE
	} else if color == CIRCLE {
		return CROSS
	} else {
		return NONE
	}

}

/*
Play : Play a move
*/
func Play(board Matrix3x3, move, color int) (Matrix3x3, int) {
	if color == 0 || !Contains(ListLegal(board), move) {
		return board, 99 // Error
	} else {
		board[move/3][move%3] = color
		if HasWon(board, color) {
			return board, color // Someone wins
		} else if len(ListLegal(board)) == 0 {
			return board, 3 // Draw
		} else {
			return board, 0 // Continue playing
		}
	}
}
