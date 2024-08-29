package knapsack

import (
	"sort"
)

type Item struct {
	Weight, Value int
}

// Knapsack takes in a maximum carrying capacity and a collection of items
// and returns the maximum value that can be carried by the knapsack
// given that the knapsack can only carry a maximum weight given by maximumWeight
func Knapsack(maximumWeight int, items []Item) (val int) {
	// sort items
	sort.Slice(items, func(i, j int) bool {
		if items[i].Value > 0 {
			iPerKg, jPerKg := float64(items[i].Value)/float64(items[i].Weight), float64(items[j].Value)/float64(items[j].Weight)
			if iPerKg > jPerKg {
				return true
			}
			if iPerKg == jPerKg {
				return items[i].Weight > items[j].Weight
			}
		}
		return false
	})
	var w int
	for _, it := range items {
		if it.Weight <= maximumWeight-w {
			val += it.Value
			w += it.Weight
		} else {
			return
		}
	}
	return
}
