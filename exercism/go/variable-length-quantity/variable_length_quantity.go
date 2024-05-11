package variablelengthquantity

import "errors"

func EncodeVarint(input []uint32) (en []byte) {
	for _, v := range input {
		first := true
		ent := []byte{}
		for n := v; n != 0; {
			val := byte(n & 0b01111111)
			n >>= 7
			if first {
				first = false
			} else {
				val |= 0b10000000
			}
			ent = append([]byte{val}, ent...)
		}
		if len(ent) == 0 {
			ent = append(ent, 0)
		}
		en = append(en, ent...)
	}
	return en
}

func DecodeVarint(input []byte) (de []uint32, err error) {
	var current uint32
	var done bool
	for _, b := range input {
		current <<= 7
		current |= uint32(b & 0b01111111)
		if b&0b10000000 == 0 {
			de = append(de, current)
			current, done = 0, true
		}
	}
	if !done {
		err = errors.New("bad input")
	}
	return
}
