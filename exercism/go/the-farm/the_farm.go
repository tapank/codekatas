package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood calculates food per cow
func DivideFood(fc FodderCalculator, n int) (float64, error) {
	var total float64
	var err error
	if total, err = fc.FodderAmount(n); err != nil {
		return 0.0, err
	}
	if f, err := fc.FatteningFactor(); err != nil {
		return 0.0, err
	} else {
		total *= f
	}
	return total / float64(n), nil
}

// DivideFood validates the number of cows and calculates food per cow
func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
	if err := ValidateNumberOfCows(n); err != nil {
		return 0.0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, n)
}

// DivideFood validates the number of cows
func ValidateNumberOfCows(n int) error {
	if n < 0 {
		return fmt.Errorf("%d cows are invalid: there are no negative cows", n)
	}
	if n == 0 {
		return fmt.Errorf("%d cows are invalid: no cows don't need food", n)
	}
	return nil
}
