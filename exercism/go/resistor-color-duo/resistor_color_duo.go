package resistorcolorduo

var colorValues = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Value returns the resistance value of a resistor with given colors.
func Value(colors []string) (val int) {
	for i := 0; i < len(colors) && i < 2; i++ {
		val *= 10
		val += colorValues[colors[i]]
	}
	return
}
