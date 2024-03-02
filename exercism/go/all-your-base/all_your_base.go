package allyourbase

import (
	"errors"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	// convert input to decimal
	input := 0
	for i, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		e := len(inputDigits) - i - 1
		input += Value(d, e, inputBase)
	}

	// convert decimal to output
	outputDigits := []int{}
	for input > 0 {
		outputDigits = append([]int{input % outputBase}, outputDigits...)
		input /= outputBase
	}

	// stuff in a zero if array is empty
	if len(outputDigits) == 0 {
		outputDigits = append(outputDigits, 0)
	}
	return outputDigits, nil
}

// compute positional value in decimal of a digit given its base
func Value(n int, pos int, base int) int {
	if pos == 0 {
		return n
	}
	out := 1
	for i := 0; i < pos; i++ {
		out *= base
	}
	return out * n
}
