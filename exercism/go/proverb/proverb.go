package proverb

import "fmt"

// Proverb generates a rhyme based on the given list of words.
func Proverb(rhyme []string) []string {
	if len(rhyme) == 0 {
		return nil
	}
	verse := []string{}
	for i := 1; i < len(rhyme); i++ {
		verse = append(verse, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	verse = append(verse, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	return verse
}
