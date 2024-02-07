package armstrong

import (
	"math"
)

func IsNumber(n int) bool {
	if n == 0 {
		return true
	}

	if n < 0 {
		return false
	}

	// find the number of digits
	input, digits := n, 0
	for input > 0 {
		input /= 10
		digits++
	}

	sum := 0
	for input = n; input > 0; input /= 10 {
		sum += int(math.Pow(float64(input%10), float64(digits)))
	}
	return sum == n
}
