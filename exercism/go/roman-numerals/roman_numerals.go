package romannumerals

import (
	"errors"
	"fmt"
)

var thousands = []string{"", "M", "MM", "MMM"}
var hundreds = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
var tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
var units = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("Invalid input: " + fmt.Sprint(input))
	}
	out := ""

	out += thousands[input/1000]
	input %= 1000
	out += hundreds[input/100]
	input %= 100
	out += tens[input/10]
	input %= 10
	out += units[input]
	return string(out), nil
}
