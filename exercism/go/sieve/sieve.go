package sieve

func Sieve(limit int) []int {
	// sieve is a bool array where the index indicates an int and the value at
	// that index indicates if that number is composite (true for composite,
	// false for not composite)
	sieve := make([]bool, limit+1)

	// mark 0 and 1 as composites because they are not prime
	sieve[0], sieve[1] = true, true

	// start iterating from index 2
	for n := 2; n < len(sieve); n++ {
		if !sieve[n] {
			// if the current i is not a composite, then mark all its multiples
			// as composite
			for i, j := 2, n*2; j < len(sieve); i, j = i+1, n*i {
				if j%n == 0 {
					sieve[j] = true
				}
			}
		}
	}

	// now gather primes from the sieve
	primes := []int{}
	for i, b := range sieve {
		if !b {
			primes = append(primes, i)
		}
	}
	return primes
}
