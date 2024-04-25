package queenattack

import (
	"errors"
)

func CanQueenAttack(w, b string) (bool, error) {
	// validate input
	if len(w) != 2 || len(b) != 2 {
		return false, errors.New("bad input")
	}
	if w == b {
		return false, errors.New("illegal positions: " + w + ":" + b)
	}

	// compute white column (wc), white row (wr), black column (bc), black row (br)
	wc, wr, bc, br := int(w[0]-'a'), int(w[1]-'1'), int(b[0]-'a'), int(b[1]-'1')

	// validate coordinates
	if wc < 0 || wr < 0 || wc > 7 || wr > 7 || bc < 0 || br < 0 || bc > 7 || br > 7 {
		return false, errors.New("illegal positions: " + w + ":" + b)
	}

	// check attack positions
	return wr == br || wc == bc || diff(wr, wc) == diff(br, bc), nil
}

// always returns a positive difference
func diff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
