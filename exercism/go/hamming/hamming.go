package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("unequal lengths")
	}

	hd := 0
	for i := range a {
		if a[i] != b[i] {
			hd++
		}
	}
	return hd, nil
}
