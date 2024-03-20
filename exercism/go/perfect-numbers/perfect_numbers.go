package perfect

import "errors"

// Classification type.
type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("only positive numbers are allowed")

// Classify classifies a number as perfect, abundant, or deficient.
// Returns an error if the number is not positive.
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return ClassificationDeficient, ErrOnlyPositive
	}
	sum := aliquotSum(n)
	switch {
	case sum == n:
		return ClassificationPerfect, nil
	case sum < n:
		return ClassificationDeficient, nil
	default:
		return ClassificationAbundant, nil
	}
}

// aliquotSum returns the sum of the proper divisors of n.
func aliquotSum(n int64) int64 {
	sum := int64(0)
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
