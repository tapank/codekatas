package ocr

import "strings"

var dmap = map[string]string{
	" _ \n| |\n|_|\n   ": "0",
	"   \n  |\n  |\n   ": "1",
	" _ \n _|\n|_ \n   ": "2",
	" _ \n _|\n _|\n   ": "3",
	"   \n|_|\n  |\n   ": "4",
	" _ \n|_ \n _|\n   ": "5",
	" _ \n|_ \n|_|\n   ": "6",
	" _ \n  |\n  |\n   ": "7",
	" _ \n|_|\n|_|\n   ": "8",
	" _ \n|_|\n _|\n   ": "9",
}

func Recognize(s string) []string {
	if len(s) == 0 || s[0] != '\n' {
		return []string{"?"}
	}

	s = s[1:]
	parsed := splitInput(s)
	if parsed == nil {
		return []string{"?"}
	}

	digits := []string{}
	for _, row := range parsed {
		aNumber := ""
		for _, word := range row {
			aNumber += recognizeDigit(word)
		}
		digits = append(digits, aNumber)
	}
	return digits
}

func splitInput(s string) [][]string {
	rows := strings.Split(s, "\n")
	// number of rows should be a multiple of 4
	if len(rows) == 0 || len(rows)%4 != 0 {
		return nil
	}

	tokens := [][]string{}
	for i := 0; i < len(rows); i += 4 {
		l := len(rows[i])
		// sets of 4 rows should be equal length and a multiple of 3
		if l == 0 || l%3 != 0 || l != len(rows[i+1]) || l != len(rows[i+2]) || l != len(rows[i+3]) {
			return nil
		}
		line := []string{}
		// pluck 3 bytes from each row and join them with new lines
		for j := 0; j < l; j += 3 {
			s := rows[i+0][j:j+3] + "\n"
			s += rows[i+1][j:j+3] + "\n"
			s += rows[i+2][j:j+3] + "\n"
			s += rows[i+3][j : j+3]
			line = append(line, s)
		}
		tokens = append(tokens, line)
	}
	return tokens
}

func recognizeDigit(s string) string {
	if d, ok := dmap[s]; ok {
		return d
	}
	return "?"
}
