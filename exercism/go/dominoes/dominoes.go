package dominoes

type Domino [2]int

func Chainit(chain []Domino, pieces map[int]Domino) ([]Domino, bool) {
	match := chain[len(chain)-1][1]
	for key, domino := range pieces {
		// create a new copy of the domino
		d := Domino{domino[0], domino[1]}
		if d[1] == match {
			d[0], d[1] = d[1], d[0]
		}

		if d[0] == match {
			// create a copy of chain and append the match
			nextchain := []Domino{}
			copy(chain, nextchain)
			nextchain = append(chain, d)

			// create a copy of pieces and remove the match
			nextpieces := make(map[int]Domino)
			for k, v := range pieces {
				if k != key {
					nextpieces[k] = v
				}
			}
			if len(nextpieces) > 0 {
				if c, ok := Chainit(nextchain, nextpieces); ok {
					return c, true
				}
			} else if nextchain[0][0] == d[1] {
				return nextchain, true
			}
		}
	}
	return chain, false
}

func MakeChain(input []Domino) ([]Domino, bool) {
	if len(input) == 0 {
		return input, true
	}
	start, match := input[0][0], input[0][1]
	if len(input) == 1 {
		return input, start == match
	}

	// put remaining dominos in a map for further processing
	pieces := make(map[int]Domino, len(input)-1)
	for i, v := range input[1:] {
		pieces[i] = v
	}
	return Chainit([]Domino{input[0]}, pieces)
}
