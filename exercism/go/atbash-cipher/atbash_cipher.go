package atbash

func Atbash(s string) string {
	// does two things after changing upper case to lower case
	// 1. shift alphabet and leave alone digits
	// 2. eat up everything else
	// return true for case 1 and false for case 2
	f := func(r rune) (rune, bool) {
		if r >= 'A' && r <= 'Z' {
			r |= 0b100000
		}
		switch {
		case r >= 'a' && r <= 'z':
			return rune('z' - (r - 'a')), true
		case r >= '0' && r <= '9':
			return r, true
		}
		return 0, false
	}

	// accumulator
	ct := []rune{}
	// counter for group size
	ctr, max := 0, 5
	for _, r := range s {
		if cr, ok := f(r); ok {
			// when group size reached, add space and reset counter
			if ctr == max {
				ct = append(ct, ' ')
				ctr = 0
			}
			ct = append(ct, cr)
			ctr++
		}
	}
	return string(ct)
}
