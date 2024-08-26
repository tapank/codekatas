package poker

import (
	"fmt"
)

type Suite rune

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

func (hand Hand) String() string {
	s := "[ "
	for _, card := range hand.Cards {
		s += fmt.Sprintf("%d%s ", card.Rank, string(card.Suite))
	}
	s += "]"
	return s
}

func New(s string) (Hand, error) {
	hand := Hand{[5]Card{}}
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

func BestHand(hands []string) ([]string, error) {
	fmt.Println("Given hands:", hands)
	for i, hand := range hands {
		if hs, err := New(hand); err != nil {
			return nil, err
		} else {
			fmt.Println(i, hs)
		}
	}
	return nil, nil
}
