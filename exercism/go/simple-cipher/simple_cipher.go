package cipher

type shift int
type vigenere []int

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return shift(distance)
}

func (c shift) Encode(input string) string {
	out := []rune{}
	for _, r := range input {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			r |= 0b100000
			offset := (int(r-'a') + int(c) + 26) % 26
			out = append(out, 'a'+rune(offset))
		}
	}
	return string(out)
}

func (c shift) Decode(input string) string {
	out := []rune{}
	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			offset := (int(r-'a') - int(c) + 26) % 26
			out = append(out, 'a'+rune(offset))
		}
	}
	return string(out)
}

func NewVigenere(key string) Cipher {
	var v vigenere = []int{}
	var valid bool
	for _, r := range key {
		if r >= 'a' && r <= 'z' {
			v = append(v, int(r-'a'))
			if r != 'a' {
				valid = true
			}
		} else {
			valid = false
			break
		}
	}
	if !valid {
		return nil
	}
	return v
}

func (v vigenere) Encode(input string) string {
	out := []rune{}
	pos := 0
	for _, r := range input {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			r |= 0b100000
			offset := (int(r-'a') + v[pos%len(v)]) % 26
			pos++
			out = append(out, 'a'+rune(offset))
		}
	}
	return string(out)
}

func (v vigenere) Decode(input string) string {
	out := []rune{}
	pos := 0
	for _, r := range input {
		if r >= 'a' && r <= 'z' {
			offset := (int(r-'a') - v[pos%len(v)] + 26) % 26
			pos++
			out = append(out, 'a'+rune(offset))
		}
	}
	return string(out)
}
