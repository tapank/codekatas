package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	given := normalize(subject)
	anagrams := []string{}
	for _, word := range candidates {
		if normalize(word) == given && !strings.EqualFold(subject, word) {
			anagrams = append(anagrams, word)
		}
	}
	return anagrams
}

func normalize(word string) string {
	word = strings.ToLower(word)
	runes := []rune{}
	for _, r := range word {
		runes = append(runes, r)
	}
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
