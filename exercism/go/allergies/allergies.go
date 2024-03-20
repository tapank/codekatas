package allergies

var allergens = map[uint]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

var allergenMap = func() map[string]uint {
	m := make(map[string]uint)
	for k, v := range allergens {
		m[v] = k
	}
	return m
}()

func Allergies(allergies uint) []string {
	anames := []string{}
	for k, v := range allergens {
		if allergies&k != 0 {
			anames = append(anames, v)
		}
	}
	return anames
}

func AllergicTo(allergies uint, allergen string) bool {
	return allergies&allergenMap[allergen] != 0
}
