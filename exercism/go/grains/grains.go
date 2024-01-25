package grains

import (
	"errors"
	"fmt"
)

func Square(number int) (grains uint64, err error) {
	if number < 1 || number > 64 {
		err = errors.New("Invalid square: " + fmt.Sprint(number))
		return
	}

	grains = 1
	for i := 2; i <= number; i++ {
		grains *= 2
	}
	return
}

func Total() (grains uint64) {
	for i := 1; i <= 64; i++ {
		g, _ := Square(i)
		grains += g
	}
	return
}
