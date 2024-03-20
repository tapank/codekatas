package encode

import (
	"strconv"
	"strings"
)

// RunLengthEncode encodes a string using run-length encoding.
func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	current, count, output := rune(input[0]), 0, ""
	for _, r := range input {
		if r == current {
			count++
		} else {
			if count > 1 {
				output += strconv.Itoa(count)
			}
			output += string(current)
			current, count = r, 1
		}
	}
	if count > 1 {
		output += strconv.Itoa(count)
	}
	output += string(current)
	return output
}

// RunLengthDecode decodes a run-length encoded string.
func RunLengthDecode(input string) string {
	count, output := 0, ""
	for _, r := range input {
		if r >= '0' && r <= '9' {
			count = count*10 + int(r-'0')
		} else {
			if count == 0 {
				count = 1
			}
			output += strings.Repeat(string(r), count)
			count = 0
		}
	}
	return output
}
