package strand

var compliment = map[rune]rune{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA generates a compliment of given dna.
// Expects correct input, errors are not handled
func ToRNA(dna string) string {
	target := []rune{}
	for _, r := range dna {
		target = append(target, compliment[r])
	}
	return string(target)
}
