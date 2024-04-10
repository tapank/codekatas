package piglatin

import "strings"

// Sentence is a bad solution because it does not "feel" clean
// However, most production code I have seen is worse than this
// so I think we should be fine here. I am sure new tests can be
// added to prove this solution wrong or incomplete. At this point
// we only care that the tests pass.
func Sentence(s string) string {
	// split sentence into words
	words := strings.Split(s, " ")

	// no words found
	if len(words) == 0 {
		return s
	}

	// only one word found
	if len(words) == 1 {
		return word(words[0])
	}

	// more than one word
	sentence := ""
	for _, w := range words[:len(words)-1] {
		sentence += word(w) + " "
	}
	return sentence + word(words[len(words)-1])
}

func word(w string) string {
	switch {
	case len(w) == 0:
		return ""
	case isvowel(w[0]) || isvowely(w):
		return w + "ay"
	case w[0] == 'y':
		return w[1:] + "yay"
	case strings.Index(w, "qu") == 0:
		return w[2:] + w[0:2] + "ay"
	case strings.Index(w, "qu") == 1 && !isvowel(w[0]):
		return w[3:] + w[0:3] + "ay"
	}
	if ok, index := yconcluster(w); ok {
		if len(w) == 2 {
			return "y" + string(w[:1]) + "ay"
		}
		return w[index:] + w[0:index] + "ay"

	}
	if ok, index := consonantsound(w); ok {
		return w[index:] + w[0:index] + "ay"
	}
	return w
}

func yconcluster(s string) (bool, int) {
	if len(s) < 2 {
		return false, -1
	}
	for i, r := range s {
		if isvowel(byte(r)) {
			return false, -1
		}
		if r == 'y' && i > 0 {
			return true, i
		}
	}
	return false, -1
}

func consonantsound(s string) (bool, int) {
	for i, r := range s {
		if isvowel(byte(r)) {
			return true, i
		}
	}
	return false, -1
}

func isvowel(r byte) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

func isvowely(s string) bool {
	if len(s) > 1 {
		st := s[0:2]
		if st == "xr" || st == "yt" {
			return true
		}
	}
	return false
}
