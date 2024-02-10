package sublist

func Sublist(l1, l2 []int) Relation {
	sub := func(a, b []int) bool {
		la, lb := len(a), len(b)
		if la == 0 {
			return true
		}
		if la > lb {
			return false
		}
	outer:
		for bIndex := 0; bIndex < lb-la+1; bIndex++ {
			if a[0] == b[bIndex] {
				for aIndex := 0; aIndex < la && bIndex+aIndex < lb; aIndex++ {
					if a[aIndex] != b[aIndex+bIndex] {
						continue outer
					}
				}
				return true
			}
		}
		return false
	}

	s1, s2 := sub(l1, l2), sub(l2, l1)
	switch {
	case s1 && s2:
		return RelationEqual
	case s1:
		return RelationSublist
	case s2:
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}
