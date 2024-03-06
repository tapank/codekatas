package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix [][]int

// New creates a new Matrix from the given string.
func New(s string) (Matrix, error) {
	matrix := [][]int{}
	length := -1
	for _, row := range strings.Split(s, "\n") {
		var r []int
		for _, col := range strings.Fields(row) {
			n, err := strconv.Atoi(col)
			if err != nil {
				return matrix, err
			}
			r = append(r, n)
		}
		if length == -1 {
			length = len(r)
		} else if length != len(r) {
			return matrix, errors.New("uneven rows")
		}
		matrix = append(matrix, r)
	}
	return matrix, nil
}

// Cols returns the columns without affecting the matrix.
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for _, row := range m {
		for j, val := range row {
			cols[j] = append(cols[j], val)
		}
	}
	return cols
}

// Rows returns the rows without affecting the matrix.
func (m Matrix) Rows() [][]int {
	matrix := make([][]int, len(m))
	for i, r := range m {
		matrix[i] = make([]int, len(r))
		copy(matrix[i], r)
	}
	return matrix
}

// Set sets the value at the given row and column.
func (m Matrix) Set(row, col, val int) bool {
	if row < 0 || row >= len(m) {
		return false
	}
	if col < 0 || col >= len(m[0]) {
		return false
	}
	m[row][col] = val
	return true
}
