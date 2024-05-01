package rectangles

type Cell struct {
	r, c int
}

func Count(diagram []string) (count int) {
	if len(diagram) < 2 {
		return
	}

	// create rune matrix of the puzzle
	rm := make([][]rune, len(diagram))
	for i, row := range diagram {
		rm[i] = []rune(row)
	}

	// scan matrix
	for r, row := range rm {
		// scan row
		for c, ch := range row {
			// possible top left corner of rectangle(s)
			if ch == '+' {
			innterloop:
				// scan forward to the possible top right corner of rectangle(s)
				for i := c + 1; i < len(row); i++ {
					switch row[i] {
					case '-':
						// keep scanning
					case '+':
						// possible top right corner of rectangle(s)
						// now scan down
						count += scanDown(rm, Cell{r, c}, Cell{r, i})
					default:
						// gap in the top side of rectangle
						break innterloop
					}
				}
			}
		}
	}
	return
}

func scanDown(rm [][]rune, left, right Cell) (count int) {
	for i := left.r + 1; i < len(rm); i++ {
		switch rm[i][left.c] {
		case '|', '+':
		default:
			return
		}

		switch rm[i][right.c] {
		case '|', '+':
		default:
			return
		}

		if rm[i][left.c] == '+' && rm[i][right.c] == '+' {
			if areConnected(rm, Cell{i, left.c}, Cell{i, right.c}) {
				count++
			}
		}
	}
	return
}

func areConnected(rm [][]rune, left, right Cell) bool {
	for i := left.c + 1; i < right.c; i++ {
		switch rm[left.r][i] {
		case '+', '-':
			// keep going
		default:
			return false
		}
	}
	return true
}
