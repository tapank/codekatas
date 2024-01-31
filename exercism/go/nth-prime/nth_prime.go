package prime

import (
	"errors"
)

// Nth returns the nth prime number. An error is returned if the nth prime
// number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("invalid input")
	}
	primes := make([]int, 0, n)
	primes = append(primes, 2)
	for i := 3; len(primes) < n; i += 2 {
		prime := true
		for _, p := range primes {
			if i%p == 0 {
				prime = false
				break
			}
		}
		if prime {
			primes = append(primes, i)
		}
	}
	return primes[len(primes)-1], nil
}
