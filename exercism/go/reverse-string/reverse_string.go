package reverse

func Reverse(input string) string {
	in := []rune(input)
	l := len(in)
	out := make([]rune, l)

	for i, r := range in {
		out[l-1-i] = r
	}
	return string(out)
}
