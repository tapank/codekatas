package bottlesong

import "fmt"

type NumWords struct {
	Camel  string
	Lower  string
	Bottle string
}

var numbers = map[int]NumWords{
	0:  {"No", "no", "bottles"},
	1:  {"One", "one", "bottle"},
	2:  {"Two", "two", "bottles"},
	3:  {"Three", "three", "bottles"},
	4:  {"Four", "four", "bottles"},
	5:  {"Five", "five", "bottles"},
	6:  {"Six", "six", "bottles"},
	7:  {"Seven", "seven", "bottles"},
	8:  {"Eight", "eight", "bottles"},
	9:  {"Nine", "nine", "bottles"},
	10: {"Ten", "ten", "bottles"},
}

func Recite(start, count int) []string {
	if start > 10 || start < 1 || count < 1 || count > start {
		return []string{}
	}

	song := []string{}
	l1_2fmt := "%s green %s hanging on the wall,"
	l3fmt := "And if one green bottle should accidentally fall,"
	l4fmt := "There'll be %s green %s hanging on the wall."
	for n := 0; n < count; n++ {
		line1_2 := fmt.Sprintf(l1_2fmt, numbers[start-n].Camel, numbers[start-n].Bottle)
		l4 := fmt.Sprintf(l4fmt, numbers[start-n-1].Lower, numbers[start-n-1].Bottle)

		song = append(song, line1_2, line1_2)
		song = append(song, l3fmt)
		song = append(song, l4)

		if n < count-1 {
			song = append(song, "")
		}
	}
	return song
}
