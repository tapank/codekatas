package strain

func Keep[T any](elements []T, f func(item T) bool) []T {
	return KeepOrDiscard(elements, f, true)
}

func Discard[T any](elements []T, f func(item T) bool) []T {
	return KeepOrDiscard(elements, f, false)
}

func KeepOrDiscard[T any](elements []T, f func(item T) bool, keep bool) []T {
	r := []T{}
	for _, t := range elements {
		if f(t) == keep {
			r = append(r, t)
		}
	}
	return r
}
