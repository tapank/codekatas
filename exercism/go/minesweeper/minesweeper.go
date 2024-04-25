package minesweeper

// Annotate returns an annotated board
func Annotate(board []string) []string {
	m := [][]rune{} // rune matrix
	for _, row := range board {
		m = append(m, []rune(row))
	}

	for r, row := range m {
		for c, char := range row {
			if char != '*' {
				if n := count(m, [2]int{r, c}); n > 0 {
					m[r][c] = rune('0' + n)
				}
			}
		}
	}

	a := []string{} // answer
	for _, row := range m {
		a = append(a, string(row))
	}
	return a
}

var deltas = [8][2]int{
	{-1, 0},  // N
	{-1, 1},  // NE
	{0, 1},   // S
	{1, 1},   // SE
	{1, 0},   // S
	{1, -1},  // SW
	{0, -1},  // W
	{-1, -1}, // NW
}

func count(m [][]rune, pos [2]int) (n int) {
	for _, delta := range deltas {
		r, c := pos[0]+delta[0], pos[1]+delta[1]
		if r >= 0 && r < len(m) && c >= 0 && c < len(m[r]) {
			if m[r][c] == '*' {
				n++
			}
		}
	}
	return
}
