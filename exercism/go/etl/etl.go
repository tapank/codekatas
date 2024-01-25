package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	out := map[string]int{}
	for k, v := range in {
		for _, letter := range v {
			out[strings.ToLower(letter)] = k
		}
	}
	return out
}
