package diamond

import (
	"errors"
	"strings"
)

func Gen(b byte) (string, error) {
	// validate input
	if b < 'A' || b > 'Z' {
		return string(b), errors.New("illegal argument")
	}

	// construct empty square
	side := (b-'A')*2 + 1
	margin := side / 2
	grid := make([][]byte, side)
	for i := range grid {
		grid[i] = make([]byte, side)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	// fill in the alphabets
	for i, ch := 0, byte('A'); i < int(side)/2+1; i, ch = i+1, ch+1 {
		grid[i][margin] = ch
		grid[int(side)-i-1][margin] = ch
		grid[i][side-margin-1] = ch
		grid[int(side)-i-1][side-margin-1] = ch
		margin--
	}

	// stringify the grid
	rows := []string{}
	for i := range grid {
		rows = append(rows, string(grid[i]))
	}
	return strings.Join(rows, "\n"), nil
}
