package stateoftictactoe

import (
	"errors"
	"strings"
)

type State string

const (
	Win     State = "win"
	Ongoing State = "ongoing"
	Draw    State = "draw"
	X       rune  = 'X' // always starts
	O       rune  = 'O'
)

func StateOfTicTacToe(board []string) (state State, err error) {
	var boardfull bool
	if boardfull, err = validate(board); err != nil {
		return
	}

	switch winX, winO := won(X, board), won(O, board); {
	case winX && winO:
		err = errors.New("both cannot win")
	case winX || winO:
		state = Win
	case boardfull:
		state = Draw
	default:
		state = Ongoing
	}
	return
}

func won(r rune, board []string) bool {
	pattern := strings.Repeat(string(r), 3)

	// check rows
	for _, row := range board {
		if row == pattern {
			return true
		}
	}

	// check columns
	for i := 0; i < 3; i++ {
		col := string(board[0][i]) + string(board[1][i]) + string(board[2][i])
		if col == pattern {
			return true
		}
	}

	// check diagonals
	if cross := string(board[0][0]) + string(board[1][1]) + string(board[2][2]); cross == pattern {
		return true
	}
	if cross := string(board[0][2]) + string(board[1][1]) + string(board[2][0]); cross == pattern {
		return true
	}

	// nothing matched
	return false
}

func validate(board []string) (boardfull bool, err error) {
	// validate dimensions
	if len(board) != 3 || len(board[0]) != 3 || len(board[1]) != 3 || len(board[2]) != 3 {
		err = errors.New("bad board")
		return
	}

	// validate values
	var countX, countO, countSpace int
	for _, row := range board {
		for _, r := range row {
			switch r {
			case X:
				countX++
			case O:
				countO++
			case ' ':
				countSpace++
			default:
				err = errors.New("invalid cell value found")
				return
			}
		}
	}

	// validate consistency of moves
	if !(countX == countO+1 || countX == countO) {
		err = errors.New("inconsistent number of moves")
		return
	}

	// is the board full?
	if countSpace == 0 {
		boardfull = true
	}
	return
}
