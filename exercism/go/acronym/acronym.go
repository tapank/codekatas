// Package acronym provides functions to manipulate phrases and such
package acronym

import "strings"

// Abbreviate abbreviates a given phrase
func Abbreviate(s string) string {
	abb := []rune{}

	atstart := true
	for _, r := range s {
		if r == ' ' || r == '-' {
			atstart = true
		} else if atstart && (r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			abb = append(abb, r)
			atstart = false
		}
	}
	return strings.ToUpper(string(abb))
}
