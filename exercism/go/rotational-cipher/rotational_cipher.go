package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	in := []rune(plain)
	out := make([]rune, len(in))
	for i, r := range in {
		switch {
		case r >= 'A' && r <= 'Z':
			out[i] = 'A' + (rune(int(r)-'A'+shiftKey))%26
		case r >= 'a' && r <= 'z':
			out[i] = 'a' + (rune(int(r)-'a'+shiftKey))%26
		default:
			out[i] = r
		}
	}
	return string(out)
}
