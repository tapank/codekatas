package summultiples

func SumMultiples(limit int, divisors ...int) int {
	multiples := map[int]bool{}
	sum := 0
	for _, d := range divisors {
		if d < 1 {
			continue
		}
		for n := 1; n*d < limit; n++ {
			factor := n * d
			if !multiples[factor] {
				multiples[factor] = true
				sum += factor
			}
		}
	}
	return sum
}
