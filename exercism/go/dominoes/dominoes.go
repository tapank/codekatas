package dominoes

type Domino [2]int

func Chainit(chain []Domino, pieces map[Wrapper]struct{}) ([]Domino, bool) {
	match := chain[len(chain)-1][1]
	for k := range pieces {
		piece := Domino{k.d[0], k.d[1]}
		if piece[1] == match {
			piece[0], piece[1] = piece[1], piece[0]
		}

		if piece[0] == match {
			// create a copy of chain and append the match
			nextchain := []Domino{}
			copy(chain, nextchain)
			nextchain = append(chain, piece)

			// create a copy of pieces and remove the match
			nextpieces := make(map[Wrapper]struct{})
			for key, value := range pieces {
				if key != k {
					nextpieces[key] = value
				}
			}
			if len(nextpieces) > 0 {
				if c, ok := Chainit(nextchain, nextpieces); ok {
					return c, true
				}
			} else if nextchain[0][0] == piece[1] {
				return nextchain, true
			}
		}
	}
	return chain, false
}

type Wrapper struct {
	n int
	d Domino
}

func MakeChain(input []Domino) ([]Domino, bool) {
	chain := []Domino{}

	// empty input is ok
	if len(input) == 0 {
		return chain, true
	}

	start, match := input[0][0], input[0][1]
	chain = append(chain, input[0])

	// handle single domino
	if len(input) == 1 {
		return chain, start == match
	}

	// put remaining dominos in a map for further processing
	pieces := make(map[Wrapper]struct{}, len(input)-1)
	for i, v := range input[1:] {
		pieces[Wrapper{i, v}] = struct{}{}
	}
	return Chainit(chain, pieces)
}
