package house

import "strings"

var first_line = []string{
	"",
	"This is the house that Jack built.",
	"This is the malt",
	"This is the rat",
	"This is the cat",
	"This is the dog",
	"This is the cow with the crumpled horn",
	"This is the maiden all forlorn",
	"This is the man all tattered and torn",
	"This is the priest all shaven and shorn",
	"This is the rooster that crowed in the morn",
	"This is the farmer sowing his corn",
	"This is the horse and the hound and the horn",
}

var other_lines = []string{
	"",
	"",
	"that lay in the house that Jack built.",
	"that ate the malt",
	"that killed the rat",
	"that worried the cat",
	"that tossed the dog",
	"that milked the cow with the crumpled horn",
	"that kissed the maiden all forlorn",
	"that married the man all tattered and torn",
	"that woke the priest all shaven and shorn",
	"that kept the rooster that crowed in the morn",
	"that belonged to the farmer sowing his corn",
}

func Verse(v int) string {
	if v < 1 || v > 12 {
		return ""
	}
	verse := first_line[v]
	for ; v > 1; v-- {
		verse += "\n"
		verse += other_lines[v]
	}
	return verse
}

func Song() string {
	song := []string{}
	for i := 1; i <= 12; i++ {
		song = append(song, Verse(i))
	}
	return strings.Join(song, "\n\n")
}
