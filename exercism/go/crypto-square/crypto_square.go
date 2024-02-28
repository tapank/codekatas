package cryptosquare

import (
	"math"
)

func Encode(pt string) string {
	normal := normalize(pt)

	// one or zero char strings can be returned as is
	if len(normal) <= 1 {
		return string(normal)
	}

	// determine the row count
	rowcount := int(math.Sqrt(float64(len(normal))))
	if len(normal) > rowcount*rowcount {
		rowcount++
	}

	// create and fill rows of the square
	rows := make([][]rune, rowcount)
	for i, r := range normal {
		rows[i%rowcount] = append(rows[i%rowcount], r)
	}

	// now stitch them in a single string where each row is separated by a space
	out := []rune{}
	out = append(out, rows[0]...)
	l := len(rows[0])
	for _, row := range rows[1:] {
		out = append(out, ' ')
		// shorter rows should be padded with a space
		if len(row) < l {
			row = append(row, ' ')
		}
		out = append(out, row...)
	}
	return string(out)
}

func normalize(pt string) []rune {
	out := []rune{}
	for _, r := range pt {
		switch {
		case r >= 'a' && r <= 'z':
			// do nothing
		case r >= '0' && r <= '9':
			// do nothing
		case r >= 'A' && r <= 'Z':
			r |= 0b100000
		default:
			continue
		}
		out = append(out, r)
	}
	return out
}
