package palindrome

import "errors"

// Define Product type here.
type Product struct {
	Val            int
	Factorizations [][2]int
}

func Products(fmin, fmax int) (Product, Product, error) {
	var pmin, pmax *Product
	var err error
	for i := fmin; i <= fmax; i++ {
		for j := i; j <= fmax; j++ {
			if product := i * j; Palindrome(product) {
				if pmin == nil {
					pmin, pmax = &Product{product, [][2]int{{i, j}}}, &Product{product, [][2]int{{i, j}}}
				} else if product <= pmin.Val {
					if product == pmin.Val {
						pmin.Factorizations = append(pmin.Factorizations, [2]int{i, j})
					} else {
						pmin = &Product{product, [][2]int{{i, j}}}
					}
				} else if product >= pmax.Val {
					if product == pmax.Val {
						pmax.Factorizations = append(pmax.Factorizations, [2]int{i, j})
					} else {
						pmax = &Product{product, [][2]int{{i, j}}}
					}
				}
			}
		}
	}
	if pmin == nil {
		pmin, pmax = &Product{}, &Product{}
		if fmin > fmax {
			err = errors.New("fmin > fmax")
		} else {
			err = errors.New("no palindromes")
		}
	}
	return *pmin, *pmax, err
}

func Palindrome(n int) bool {
	digits := []int{}
	for ; n != 0; n /= 10 {
		digits = append(digits, n%10)
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
