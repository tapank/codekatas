package isbn

func IsValidISBN(s string) bool {
	isbn, ok := extractISBN(s)
	if !ok {
		return false
	}

	sum := 0
	for i := 0; i < 10; i++ {
		sum += (10 - i) * isbn[i]
	}
	return sum%11 == 0
}

func extractISBN(s string) (isbn []int, ok bool) {
	for _, d := range s {
		if d >= '0' && d <= '9' {
			isbn = append(isbn, int(d-'0'))
		} else if len(isbn) == 9 && (d == 'x' || d == 'X') {
			isbn = append(isbn, 10)
		} else if d != '-' {
			// invalid!
			break
		}
		if len(isbn) > 10 {
			break
		}
	}
	return isbn, len(isbn) == 10
}
