package say

import "strings"

var words = map[int64]string{
	1:                 "one",
	2:                 "two",
	3:                 "three",
	4:                 "four",
	5:                 "five",
	6:                 "six",
	7:                 "seven",
	8:                 "eight",
	9:                 "nine",
	10:                "ten",
	11:                "eleven",
	12:                "twelve",
	13:                "thirteen",
	14:                "fourteen",
	15:                "fifteen",
	16:                "sixteen",
	17:                "seventeen",
	18:                "eighteen",
	19:                "nineteen",
	20:                "twenty",
	30:                "thirty",
	40:                "forty",
	50:                "fifty",
	60:                "sixty",
	70:                "seventy",
	80:                "eighty",
	90:                "ninety",
	100:               "hundred",
	1000:              "thousand",
	1_000_000:         "million",
	1_000_000_000:     "billion",
	1_000_000_000_000: "trillion",
}

func Say(n int64) (string, bool) {
	// validate
	if n < 0 || n > 999_999_999_999 {
		return "", false
	}

	if n == 0 {
		return "zero", true
	}

	parts := []string{}
	// billions
	if part := SayTillThousand(n % 1_000_000_000_000 / 1_000_000_000); part != "" {
		parts = append(parts, part+" billion")
	}

	// millions
	if part := SayTillThousand(n % 1_000_000_000 / 1_000_000); part != "" {
		parts = append(parts, part+" million")
	}

	// thousands
	if part := SayTillThousand(n % 1_000_000 / 1_000); part != "" {
		parts = append(parts, part+" thousand")
	}

	// remainder
	if part := SayTillThousand(n % 1000); part != "" {
		parts = append(parts, part)
	}

	return strings.Join(parts, " "), true
}

func SayTillThousand(n int64) string {
	if n <= 20 {
		return words[n]
	}

	parts := []string{}
	if hundreds := words[(n%1000)/100]; hundreds != "" {
		parts = append(parts, hundreds+" hundred")
	}

	if tens, units := words[(n%100)/10*10], words[n%10]; tens != "" && units != "" {
		parts = append(parts, tens+"-"+units)
	} else if w := tens + units; w != "" {
		parts = append(parts, w)
	}

	return strings.Join(parts, " ")
}
