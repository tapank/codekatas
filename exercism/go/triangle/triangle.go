package triangle

type Kind int

const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// KindFromSides returns the triangle type
func KindFromSides(a, b, c float64) Kind {
	if a+b < c || b+c < a || a+c < b || a+b+c <= 0 {
		return NaT
	}
	if a == b && a == c {
		return Equ
	}
	if a == b || b == c || c == a {
		return Iso
	}
	return Sca
}
