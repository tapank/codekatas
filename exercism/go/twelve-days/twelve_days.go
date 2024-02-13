package twelve

import "fmt"

var days = map[int]string{
	12: "twelfth",
	11: "eleventh",
	10: "tenth",
	9:  "ninth",
	8:  "eighth",
	7:  "seventh",
	6:  "sixth",
	5:  "fifth",
	4:  "fourth",
	3:  "third",
	2:  "second",
	1:  "first",
}

var what = map[int]string{
	12: " twelve Drummers Drumming,",
	11: " eleven Pipers Piping,",
	10: " ten Lords-a-Leaping,",
	9:  " nine Ladies Dancing,",
	8:  " eight Maids-a-Milking,",
	7:  " seven Swans-a-Swimming,",
	6:  " six Geese-a-Laying,",
	5:  " five Gold Rings,",
	4:  " four Calling Birds,",
	3:  " three French Hens,",
	2:  " two Turtle Doves,",
	1:  " and a Partridge in a Pear Tree.",
}

func Verse(i int) string {
	if i < 1 || i > 12 {
		return ""
	}

	verse := fmt.Sprintf("On the %s day of Christmas my true love gave to me:", days[i])
	suffix := what[1]
	if i == 1 {
		suffix = " a Partridge in a Pear Tree."
	}
	for n := 2; n <= i; n++ {
		suffix = what[n] + suffix
	}
	return verse + suffix
}

func Song() string {
	song := ""
	for n := 1; n <= 12; n++ {
		song += Verse(n)
		if n < 12 {
			song += "\n"
		}
	}
	return song
}
