package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	n1.a += n2.a
	n1.b += n2.b
	return n1
}

func (n1 Number) Subtract(n2 Number) Number {
	n1.a -= n2.a
	n1.b -= n2.b
	return n1
}

func (n1 Number) Multiply(n2 Number) Number {
	n1.a, n1.b = n1.a*n2.a-n1.b*n2.b, n1.b*n2.a+n1.a*n2.b
	return n1
}

func (n Number) Times(factor float64) Number {
	n.a *= factor
	n.b *= factor
	return n
}

func (n1 Number) Divide(n2 Number) Number {
	div := n2.a*n2.a + n2.b*n2.b
	n1.a, n1.b = (n1.a*n2.a+n1.b*n2.b)/div, (n1.b*n2.a-n1.a*n2.b)/div
	return n1
}

func (n Number) Conjugate() Number {
	n.b *= -1
	return n
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.a*n.a + n.b*n.b)
}

func (n Number) Exp() Number {
	n1 := Number{math.Cos(n.b), math.Sin(n.b)}
	n1 = n1.Times(math.Pow(math.E, n.a))
	return n1
}
