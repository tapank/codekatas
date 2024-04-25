package pythagorean

import "slices"

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	var a, b, c int
	triplets := []Triplet{}
	for c = max; c > b; c-- {
		for b = c - 1; b >= a; b-- {
			for a = b; a >= min; a-- {
				if a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	slices.Reverse(triplets)
	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	var a, b, c int
	triplets := []Triplet{}
	for c = p - 2; c > b; c-- {
		for b = c - 1; b >= a; b-- {
			for a = b; a >= 1; a-- {
				if a+b+c == p && a*a+b*b == c*c {
					triplets = append(triplets, Triplet{a, b, c})
				}
			}
		}
	}
	return triplets
}
