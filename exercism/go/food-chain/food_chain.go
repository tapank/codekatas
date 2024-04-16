package foodchain

import (
	"fmt"
	"strings"
)

var items = map[int][]string{
	1: {"fly", ""},
	2: {"spider", "It wriggled and jiggled and tickled inside her.\n"},
	3: {"bird", "How absurd to swallow a bird!\n"},
	4: {"cat", "Imagine that, to swallow a cat!\n"},
	5: {"dog", "What a hog, to swallow a dog!\n"},
	6: {"goat", "Just opened her throat and swallowed a goat!\n"},
	7: {"cow", "I don't know how she swallowed a cow!\n"},
	8: {"horse", "She's dead, of course!"},
}

func Verse(v int) string {
	var animal, action string
	if item, ok := items[v]; ok {
		animal, action = item[0], item[1]
	} else {
		return ""
	}

	verse := fmt.Sprintf("I know an old lady who swallowed a %s.\n", animal)
	if v > 1 {
		verse += action
	}
	if v == 8 {
		return verse
	}

	for i := v; i > 1; i-- {
		if i == 3 {
			verse += fmt.Sprintf("She swallowed the %s to catch the %s that wriggled and jiggled and tickled inside her.\n", items[i][0], items[i-1][0])
		} else {
			verse += fmt.Sprintf("She swallowed the %s to catch the %s.\n", items[i][0], items[i-1][0])
		}
	}
	verse += "I don't know why she swallowed the fly. Perhaps she'll die."
	return verse
}

func Verses(start, end int) string {
	var verses []string
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
	}
	return strings.Join(verses, "\n\n")
}

func Song() string {
	return Verses(1, 8)
}
