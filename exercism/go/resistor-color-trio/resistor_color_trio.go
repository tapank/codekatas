package resistorcolortrio

import "strconv"

var code = map[string]int{
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

const (
	G = 1_000_000_000
	M = 1_000_000
	K = 1_000
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
// WARNING: bad input not handled!
func Label(colors []string) string {
	if len(colors) < 3 {
		return ""
	}
	var value int
	value += code[colors[0]] * 10
	value += code[colors[1]]
	for i, zeros := 0, code[colors[2]]; i < zeros; i++ {
		value *= 10
	}

	var unit string
	switch {
	case value >= G && value%G == 0:
		unit = " gigaohms"
		value /= G
	case value >= M && value%M == 0:
		unit = " megaohms"
		value /= M
	case value >= K && value%K == 0:
		unit = " kiloohms"
		value /= K
	default:
		unit = " ohms"
	}

	return strconv.Itoa(value) + unit
}
