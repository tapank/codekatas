package luhn

func Valid(id string) bool {
	// extract digits and validate input
	digits, valid := extractDigits(id)
	if !valid {
		return false
	}

	// double every second from the end
	for i := len(digits) - 2; i >= 0; i -= 2 {
		if digits[i] *= 2; digits[i] > 9 {
			digits[i] -= 9
		}
	}

	// sum up the numbers
	sum := 0
	for _, n := range digits {
		sum += n
	}

	// we have a winner if sum is divisible by 10
	return sum%10 == 0
}

func extractDigits(id string) ([]int, bool) {
	digits := []int{}
	for _, r := range id {
		if r >= '0' && r <= '9' {
			// gather digits
			digits = append(digits, int(r-'0'))
		} else if r != ' ' {
			// non digits except spaces are invalid
			return nil, false
		}
	}

	// one or less digits are invalid
	if len(digits) < 2 {
		return nil, false
	}
	return digits, true
}
