package lsproduct

import "errors"

func LargestSeriesProduct(s string, span int) (int64, error) {
	if span < 1 || len(s) < span {
		return 0, errors.New("invalid input")
	}
	digits, err := parse(s)
	if err != nil {
		return 0, err
	}

	var max int64
	for index := 0; index <= len(digits)-span; index++ {
		var product int64 = 1
		for i := 0; i < span && (i+index) < len(digits); i++ {
			product *= digits[i+index]
		}
		if product > max {
			max = product
		}
	}
	return max, nil
}

func parse(s string) ([]int64, error) {
	digits := []int64{}
	for _, r := range s {
		if r >= '0' && r <= '9' {
			digits = append(digits, int64(r-'0'))
		} else {
			return nil, errors.New("invalid input")
		}
	}
	return digits, nil
}
