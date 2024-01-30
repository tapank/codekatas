package wordcount

import "strings"

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	f := Frequency{}
	chars := []rune{}
	for _, r := range phrase {
		if r >= 'A' && r <= 'Z' {
			r |= 0b100000
		}
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == ' ' || r == '\'' {
			chars = append(chars, r)
		} else {
			chars = append(chars, ' ')
		}
	}
	for _, word := range strings.Split(string(chars), " ") {
		word = strings.Trim(word, "'")
		if len(word) > 0 {
			f[strings.ToLower(word)]++
		}
	}
	return f
}
