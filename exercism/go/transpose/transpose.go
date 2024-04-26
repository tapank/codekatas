package transpose

import "strings"

func Transpose(input []string) []string {
	// pad rows to ensure no row is shorter than the next row
	for i, l := len(input)-1, 0; i >= 0; i-- {
		if diff := l - len(input[i]); diff > 0 {
			input[i] += strings.Repeat(" ", diff)
		}
		l = len(input[i])
	}

	// transpose
	m := [][]rune{}
	for _, s := range input {
		for i, ch := range []rune(s) {
			if len(m) <= i {
				m = append(m, []rune{ch})
			} else {
				m[i] = append(m[i], ch)
			}
		}
	}

	// convert to list of strings
	out := []string{}
	for _, row := range m {
		out = append(out, string(row))
	}
	return out
}
