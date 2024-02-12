package listops

// IntList is an abstraction of a list of integers
type IntList []int

// Foldl folds (reduces) each item into the accumulator from the left
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	f := initial
	for _, n := range s {
		f = fn(f, n)
	}
	return f
}

// Foldr folds (reduces) each item into the accumulator from the right
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	f := initial
	mi := len(s) - 1 //max index
	for i := range s {
		f = fn(s[mi-i], f)
	}
	return f
}

// Filter returns the list of all items for which predicate(item) is True
func (s IntList) Filter(fn func(int) bool) IntList {
	list := IntList{}
	for _, v := range s {
		if fn(v) {
			list = append(list, v)
		}
	}
	return list
}

// Length returns the total number of items within it
func (s IntList) Length() int {
	return len(s)
}

// Map returns the list of the results of applying function(item) on all items
func (s IntList) Map(fn func(int) int) IntList {
	list := make(IntList, s.Length())
	for i, v := range s {
		list[i] = fn(v)
	}
	return list
}

// Reverse returns a list with all the original items, but in reversed order
func (s IntList) Reverse() IntList {
	mi := s.Length() - 1
	list := make(IntList, s.Length())
	for i, v := range s {
		list[mi-i] = v
	}
	return list
}

// Append adds all items in the second list to the end of the first list
func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

// Concat combines all items in all lists into one flattened list
func (s IntList) Concat(lists []IntList) IntList {
	list := make(IntList, s.Length())
	copy(list, s)
	for _, l := range lists {
		list = append(list, l...)
	}
	return list
}
