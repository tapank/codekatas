package phonenumber

import (
	"errors"
	"fmt"
)

/*
Number cleans up the given input and returns a ten-digit NANP number consisting
of:
  - a three-digit area code in the format NXX, followed by
  - a seven-digit local number in the format NXXXXXX

where N is any digit from 2 through 9 and X is any digit from 0 through 9
*/
func Number(phoneNumber string) (string, error) {
	digits := []rune{}
	for _, r := range phoneNumber {
		if r >= '0' && r <= '9' {
			digits = append(digits, r)
		}
	}

	if len(digits) < 10 {
		return "", errors.New("Invalid number " + string(phoneNumber))
	}

	if digits[0] == '1' {
		digits = digits[1:]
	}
	if l := len(digits); l == 10 && digits[0] > '1' && digits[3] > '1' {
		return string(digits), nil
	}
	return "", errors.New("Invalid number " + string(phoneNumber))
}

// AreaCode returns the three digit area code
func AreaCode(phoneNumber string) (num string, err error) {
	if num, err = Number(phoneNumber); err != nil {
		return "", err
	}
	return num[:3], nil
}

// Format outputs the number in (NXX) NXX-XXXX
func Format(phoneNumber string) (num string, err error) {
	if num, err = Number(phoneNumber); err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}
