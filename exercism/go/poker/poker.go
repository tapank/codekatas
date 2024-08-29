package poker

import (
	"fmt"
	"maps"
	"slices"
	"sort"
)

type Suite rune

func (s Suite) String() string {
	return string(s)
}

const (
	HEART   Suite = Suite(0x2661)
	DIAMOND Suite = Suite(0x2662)
	SPADE   Suite = Suite(0x2664)
	CLUB    Suite = Suite(0x2667)
)

type Card struct {
	Rank  int
	Suite Suite
}

type Hand struct {
	Cards [5]Card
}

func New(s string) (Hand, error) {
	hand := Hand{Cards: [5]Card{}}
	var cardIndex int
	var expectingSpace, done bool
	for _, r := range s {
		// more than 5 cards
		if done {
			return hand, fmt.Errorf("too many cards: '%s'", s)
		}

		// space is expected between cards
		if expectingSpace {
			if r == ' ' {
				expectingSpace = false
				continue
			}
			return hand, fmt.Errorf("card must be separated by space: '%s'", s)
		}

		// parse card one rune at a time
		if rank := getRank(r); rank != -1 {
			if hand.Cards[cardIndex].Rank == 1 && rank == 0 {
				hand.Cards[cardIndex].Rank = 10
			} else if hand.Cards[cardIndex].Rank == 0 {
				hand.Cards[cardIndex].Rank = rank
			} else {
				return hand, fmt.Errorf("bad rank in hand'%s'", s)
			}
		} else {
			switch suite := Suite(r); suite {
			case SPADE, HEART, DIAMOND, CLUB:
				hand.Cards[cardIndex].Suite = suite
			default:
				return hand, fmt.Errorf("bad suite:'%s', in hand'%s'", string(suite), s)
			}
			if rank = hand.Cards[cardIndex].Rank; rank < 2 || rank > 14 {
				return hand, fmt.Errorf("bad rank:'%v', in hand'%s'", rank, s)
			}

			cardIndex++
			expectingSpace = true
			if cardIndex == 5 {
				done = true
			}
		}
	}

	// not enough cards
	if !done {
		return hand, fmt.Errorf("wrong number of cards: '%s'", s)
	}

	// sort cards by rank
	sort.Slice(hand.Cards[:], func(i, j int) bool {
		return hand.Cards[i].Rank > hand.Cards[j].Rank
	})
	return hand, nil
}

func getRank(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	}
	switch r {
	case 'J', 'j':
		return 11
	case 'Q', 'q':
		return 12
	case 'K', 'k':
		return 13
	case 'A', 'a':
		return 14
	}
	return -1
}

func BestHand(handStrings []string) ([]string, error) {
	keys := func(hands map[string]Hand) ([]string, error) {
		keys := []string{}
		for k := range maps.Keys(hands) {
			keys = append(keys, k)
		}
		return keys, nil
	}
	hands := map[string]Hand{}
	for _, hs := range handStrings {
		if hand, err := New(hs); err != nil {
			return nil, err
		} else {
			hands[hs] = hand
		}
	}
	if len(hands) == 1 {
		return keys(hands)
	}
	if h := StraightFlush(hands); len(h) > 0 {
		return keys(h)
	}
	if h := FourOfAKind(hands); len(h) > 0 {
		return keys(h)
	}
	if h := FullHouse(hands); len(h) > 0 {
		return keys(h)
	}
	if h := Flush(hands); len(h) > 0 {
		return keys(h)
	}
	if h := Straight(hands); len(h) > 0 {
		return keys(h)
	}
	if h := ThreeOfAKind(hands); len(h) > 0 {
		return keys(h)
	}
	if h := TwoPair(hands); len(h) > 0 {
		return keys(h)
	}
	if h := OnePair(hands); len(h) > 0 {
		return keys(h)
	}
	if h := HighCard(hands); len(h) > 0 {
		return keys(h)
	}
	return []string{}, nil
}

func HighCard(hands map[string]Hand) map[string]Hand {
	op := map[string]Hand{}
	maxRanks := []int{}
outer:
	for hs, hand := range hands {
		kRanks := []int{}
		for _, card := range hand.Cards {
			kRanks = append(kRanks, card.Rank)
		}
		slices.Sort(kRanks)
		slices.Reverse(kRanks)
		switch slices.Compare(maxRanks, kRanks) {
		case -1:
			maxRanks = kRanks
			op = map[string]Hand{}
		case 1:
			continue outer
		}
		op[hs] = hand
	}
	return op
}

func OnePair(hands map[string]Hand) map[string]Hand {
	op := map[string]Hand{}
	var maxPairRank int
	maxKRanks := []int{}
outer:
	for hs, hand := range hands {
		rmap := map[int]int{}
		for _, card := range hand.Cards {
			rmap[card.Rank]++
		}

		var pairRank int
		kRanks := []int{}
		for rank, count := range rmap {
			if count == 2 {
				if pairRank == 0 {
					pairRank = rank
				} else {
					continue outer
				}
			} else {
				kRanks = append(kRanks, rank)
			}
		}
		if pairRank == 0 || len(kRanks) != 3 || pairRank < maxPairRank {
			continue outer
		}
		slices.Sort(kRanks)
		slices.Reverse(kRanks)
		if pairRank > maxPairRank {
			maxPairRank = pairRank
			maxKRanks = kRanks
			op = map[string]Hand{}
		}
		switch slices.Compare(maxKRanks, kRanks) {
		case -1:
			maxPairRank = pairRank
			maxKRanks = kRanks
			op = map[string]Hand{}
		case 1:
			continue outer
		}
		op[hs] = hand
	}
	return op
}

func TwoPair(hands map[string]Hand) map[string]Hand {
	tp := map[string]Hand{}
	var maxPair1Rank, maxPair2Rank, maxKickerRank int
outer:
	for hs, hand := range hands {
		rmap := map[int]int{}
		for _, card := range hand.Cards {
			rmap[card.Rank]++
		}
		var pair1Rank, pair2Rank, kickerRank int
		for rank, count := range rmap {
			if count == 2 {
				if rank > pair1Rank {
					pair1Rank, pair2Rank = rank, pair1Rank
				} else {
					pair2Rank = rank
				}
			} else {
				if kickerRank == 0 {
					kickerRank = rank
				} else {
					continue outer
				}
			}
		}
		if pair1Rank == 0 || pair2Rank == 0 || kickerRank == 0 || pair1Rank < maxPair1Rank {
			continue outer
		}
		if pair1Rank > maxPair1Rank {
			maxPair1Rank = pair1Rank
			maxPair2Rank = pair2Rank
			maxKickerRank = kickerRank
			tp = map[string]Hand{}
		}
		if pair2Rank < maxPair2Rank {
			continue
		}
		if pair2Rank > maxPair2Rank {
			maxPair1Rank = pair1Rank
			maxPair2Rank = pair2Rank
			maxKickerRank = kickerRank
			tp = map[string]Hand{}
		}
		if kickerRank < maxKickerRank {
			continue
		}
		if kickerRank > maxKickerRank {
			maxPair1Rank = pair1Rank
			maxPair2Rank = pair2Rank
			maxKickerRank = kickerRank
			tp = map[string]Hand{}
		}
		tp[hs] = hand
	}
	return tp
}

func ThreeOfAKind(hands map[string]Hand) map[string]Hand {
	toak := map[string]Hand{}
	var maxTripletRank, maxKicker1Rank, maxKicker2Rank int
outer:
	for hs, hand := range hands {
		rmap := map[int]int{}
		for _, card := range hand.Cards {
			rmap[card.Rank]++
		}
		var tripletRank, kicker1Rank, kicker2Rank int
		for rank, count := range rmap {
			if count == 3 {
				tripletRank = rank
			} else {
				if rank > kicker1Rank {
					kicker1Rank = rank
				} else {
					kicker2Rank = rank
				}
			}
		}
		if tripletRank == 0 || kicker1Rank == 0 || kicker2Rank == 0 || tripletRank < maxTripletRank {
			continue outer
		}
		if tripletRank > maxTripletRank {
			maxTripletRank = tripletRank
			maxKicker1Rank = kicker1Rank
			maxKicker2Rank = kicker2Rank
			toak = map[string]Hand{}
		}
		if kicker1Rank < maxKicker1Rank {
			continue
		}
		if kicker1Rank > maxKicker1Rank {
			maxTripletRank = tripletRank
			maxKicker1Rank = kicker1Rank
			maxKicker2Rank = kicker2Rank
			toak = map[string]Hand{}
		}
		if kicker2Rank < maxKicker2Rank {
			continue
		}
		if kicker2Rank > maxKicker2Rank {
			maxTripletRank = tripletRank
			maxKicker1Rank = kicker1Rank
			maxKicker2Rank = kicker2Rank
			toak = map[string]Hand{}
		}
		toak[hs] = hand
	}
	return toak
}

func Straight(hands map[string]Hand) map[string]Hand {
	var straight map[string]Hand
	var maxStraight *Hand
outer:
	for hs, hand := range hands {
		r := hand.Cards[0].Rank + 1
		for i, card := range hand.Cards {
			if r-1 != card.Rank {
				if i != 1 && r != 14 {
					continue outer
				}
			}
			r = card.Rank
		}

		maxRank := hand.Cards[0].Rank
		if hand.Cards[0].Rank == 14 && hand.Cards[1].Rank == 5 {
			maxRank = 5
		}
		if maxStraight == nil || maxRank > maxStraight.Cards[0].Rank {
			straight = map[string]Hand{}
			maxStraight = &hand
		} else if maxRank < maxStraight.Cards[0].Rank {
			continue
		}
		straight[hs] = hand
	}
	return straight
}

func Flush(hands map[string]Hand) map[string]Hand {
	flush := map[string]Hand{}
	var maxFlush *Hand
outer:
	for hs, hand := range hands {
		smap := map[Suite]int{}
		for _, card := range hand.Cards {
			smap[card.Suite]++
		}
		if len(smap) != 1 {
			continue
		}
		if maxFlush == nil {
			maxFlush = &hand
			flush[hs] = hand
			continue
		}

		for i := 0; i < 5; i++ {
			if hand.Cards[i].Rank < maxFlush.Cards[i].Rank {
				continue outer
			} else if hand.Cards[i].Rank > maxFlush.Cards[i].Rank {
				maxFlush = &hand
				flush = map[string]Hand{}
				break
			}
		}
		flush[hs] = hand
	}
	return flush
}

func FullHouse(hands map[string]Hand) map[string]Hand {
	fh := map[string]Hand{}
	var maxTripletRank, maxPairRank int
outer:
	for hs, hand := range hands {
		rmap := map[int]int{}
		for _, card := range hand.Cards {
			rmap[card.Rank]++
		}
		var tripletRank, pairRank int
		for rank, count := range rmap {
			if count == 3 {
				tripletRank = rank
			} else if count == 2 {
				pairRank = rank
			} else {
				continue outer
			}
		}
		if tripletRank == 0 || pairRank == 0 || tripletRank < maxTripletRank {
			continue outer
		}
		if tripletRank > maxTripletRank {
			maxTripletRank = tripletRank
			maxPairRank = pairRank
			fh = map[string]Hand{}
		}
		if pairRank < maxPairRank {
			continue outer
		}
		if pairRank > maxPairRank {
			maxPairRank = pairRank
			fh = map[string]Hand{}
		}
		fh[hs] = hand
	}
	return fh
}

func FourOfAKind(hands map[string]Hand) map[string]Hand {
	foak := map[string]Hand{}
	var maxSequenceRank, maxKickerRank int
outer:
	for hs, hand := range hands {
		rmap := map[int]int{}
		for _, card := range hand.Cards {
			rmap[card.Rank]++
		}
		var sequenceRank, kickerRank int
		for rank, count := range rmap {
			if count == 4 {
				sequenceRank = rank
			} else if count == 1 {
				kickerRank = rank
			} else {
				continue outer
			}
		}
		if sequenceRank == 0 || kickerRank == 0 || sequenceRank < maxSequenceRank {
			continue outer
		}
		if sequenceRank > maxSequenceRank {
			maxSequenceRank = sequenceRank
			maxKickerRank = kickerRank
			foak = map[string]Hand{}
		}
		if kickerRank < maxKickerRank {
			continue
		}
		if kickerRank > maxKickerRank {
			maxKickerRank = kickerRank
			foak = map[string]Hand{}
		}
		foak[hs] = hand
	}
	return foak
}

func StraightFlush(hands map[string]Hand) map[string]Hand {
	straights := map[string]Hand{}
	maxRank := 0
outer:
	for hstring, hand := range hands {
		s, r := hand.Cards[0].Suite, hand.Cards[0].Rank+1
		for _, card := range hand.Cards {
			if s != card.Suite || r-1 != card.Rank {
				continue outer
			}
			r = card.Rank
		}
		if hand.Cards[0].Rank > maxRank {
			straights = map[string]Hand{}
			maxRank = hand.Cards[0].Rank
		}
		if hand.Cards[0].Rank == maxRank {
			straights[hstring] = hand
		}
	}
	return straights
}
