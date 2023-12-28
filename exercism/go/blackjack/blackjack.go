package blackjack

var vals = map[string]int{
	"ace":   11,
	"eight": 8,
	"two":   2,
	"nine":  9,
	"three": 3,
	"ten":   10,
	"four":  4,
	"jack":  10,
	"five":  5,
	"queen": 10,
	"six":   6,
	"king":  10,
	"seven": 7,
}

const (
	Stand = "S"
	Hit   = "H"
	Split = "P"
	Win   = "W"
)

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return vals[card]
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	c1, c2, dc := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	c1n2 := c1 + c2
	decision := Stand
	switch {
	case c1 == 11 && c2 == 11:
		decision = Split
	case c1n2 == 21:
		if dc < 10 {
			decision = Win
		} else {
			decision = Stand
		}
	case c1n2 >= 17 && c1n2 <= 20:
		decision = Stand
	case c1n2 >= 12 && c1n2 <= 16:
		if dc >= 7 {
			decision = Hit
		} else {
			decision = Stand
		}
	case c1n2 <= 11:
		decision = Hit
	}
	return decision
}
