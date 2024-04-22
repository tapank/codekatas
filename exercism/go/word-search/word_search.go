package wordsearch

import (
	"errors"
	"fmt"
)

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	// create a rune matrix of the puzzle
	p := make([][]rune, len(puzzle))
	for i := range puzzle {
		p[i] = []rune(puzzle[i])
	}

	// create the answer map
	a := map[string][2][2]int{}

words:
	// check one word at a time
	for _, word := range words {
		if len(word) == 0 {
			return nil, errors.New("empty word in word list")
		}

		for r, row := range p {
			for c, ch := range row {
				if ch != rune(word[0]) {
					continue
				}
				start := [2]int{r, c}
				if end, ok := findword(p, word, start); ok {
					a[word] = [2][2]int{{c, r}, {end[1], end[0]}}
					continue words
				}
			}
		}
		// word not found
		return a, fmt.Errorf("word:'%s' not found", word)
	}

	// all words found
	return a, nil
}

// delta coordinates for traversing each direction
var bearing = [][2]int{
	{0, 1},   // E
	{0, -1},  // W
	{1, 0},   // S
	{-1, 0},  // N
	{-1, 1},  // NE
	{1, 1},   // SE
	{1, -1},  // SW
	{-1, -1}, // NW
}

// findword scans for a given word in all directions in the given matrix
// starting at given coordinates and returns the end coordinates and true
// if the word is found and false if not found
func findword(matrix [][]rune, word string, start [2]int) ([2]int, bool) {
tack:
	for _, d := range bearing {
		r, c := start[0], start[1]
		dr, dc := d[0], d[1]
		for _, ch := range word {
			if r >= 0 && r < len(matrix) && c >= 0 && c < len(matrix[r]) && matrix[r][c] == ch {
				r += dr
				c += dc
			} else {
				// no match, change bearing
				continue tack
			}
		}
		// word found
		return [2]int{r - dr, c - dc}, true
	}
	// word not found in any direction
	return [2]int{}, false
}
