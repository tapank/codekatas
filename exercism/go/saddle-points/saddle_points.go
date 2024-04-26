package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix and Pair types.
type Matrix [][]int
type Pair [2]int

func New(s string) (*Matrix, error) {
	var m Matrix = [][]int{}
	if len(strings.Trim(s, " \n")) == 0 {
		return &m, nil
	}

	for r, row := range strings.Split(s, "\n") {
		m = append(m, []int{})
		for _, v := range strings.Split(row, " ") {
			if n, err := strconv.Atoi(v); err != nil {
				return nil, errors.New("cannot parse matrix")
			} else {
				m[r] = append(m[r], n)
			}
		}
	}
	return &m, nil
}

func (m *Matrix) Saddle() []Pair {
	pairs := []Pair{}
	for r, row := range *m {
		// find highest point in row
		maxInRow := 0
		for _, n := range row {
			if n > maxInRow {
				maxInRow = n
			}
		}

		// check if the highest in row is the lowest point in column
	row:
		for c, n := range row {
			if n == maxInRow {
				minInCol := maxInRow
				for i := 0; i < len(*m); i++ {
					if (*m)[i][c] < minInCol {
						continue row
					}
				}
				pairs = append(pairs, [2]int{r + 1, c + 1})
			}
		}
	}
	return pairs
}
