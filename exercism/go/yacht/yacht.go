package yacht

import (
	"sort"
)

func Score(dice []int, category string) int {
	sort.Ints(dice)
	switch category {
	case "yacht":
		if maxCount(dice) == 5 {
			return 50
		}
	case "ones":
		return count(dice, 1)
	case "twos":
		return count(dice, 2) * 2
	case "threes":
		return count(dice, 3) * 3
	case "fours":
		return count(dice, 4) * 4
	case "fives":
		return count(dice, 5) * 5
	case "sixes":
		return count(dice, 6) * 6
	case "full house":
		c1, c2 := count(dice, dice[0]), count(dice, dice[4])
		if (c1 == 3 && c2 == 2) || (c1 == 2 && c2 == 3) {
			return total(dice)
		}
	case "four of a kind":
		c1, c2 := count(dice, dice[0]), count(dice, dice[4])
		if c1 >= 4 {
			return dice[0] * 4
		}
		if c2 >= 4 {
			return dice[4] * 4
		}
	case "little straight":
		if isstraight(dice) && dice[0] == 1 {
			return 30
		}
	case "big straight":
		if isstraight(dice) && dice[0] == 2 {
			return 30
		}
	case "choice":
		return total(dice)
	}
	return 0
}

func maxCount(d []int) int {
	var max, count int
	prev := -1
	for _, n := range d {
		if n == prev {
			count++
		} else {
			prev = n
			count = 1
		}
		if count > max {
			max = count
		}
	}
	return max
}

func count(d []int, n int) int {
	count := 0
	for _, v := range d {
		if v == n {
			count++
		}
	}
	return count
}

func total(d []int) (total int) {
	for _, v := range d {
		total += v
	}
	return
}

func isstraight(d []int) bool {
	for i := 0; i < len(d)-1; i++ {
		if d[i+1]-d[i] != 1 {
			return false
		}
	}
	return true
}
