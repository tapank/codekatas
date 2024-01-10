package isogram

func IsIsogram(word string) bool {
	tracker := map[rune]bool{}
	for _, r := range word {
		if r >= 'A' && r <= 'Z' {
			r |= 0b100000
		}
		if r >= 'a' && r <= 'z' {
			if tracker[r] {
				return false
			}
			tracker[r] = true
		}
	}
	return true
}
