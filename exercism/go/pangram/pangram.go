package pangram

func IsPangram(input string) bool {
	tracker := map[rune]bool{}
	for _, r := range input {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') {
			// convert to lower case
			tracker[r|0b100000] = true
		}
		// we return as soon as we have seen every alphabet
		if len(tracker) == 26 {
			return true
		}
	}
	return false
}
