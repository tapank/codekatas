package romannumerals

import (
	"errors"
	"fmt"
)

var M = []string{"", "M", "MM", "MMM"}
var C = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var X = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var I = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3999 {
		return "", errors.New("Invalid input: " + fmt.Sprint(n))
	}
	return M[n/1000] + C[(n%1000)/100] + X[(n%100)/10] + I[n%10], nil
}
