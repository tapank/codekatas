package stringset

// Set type.
type Set map[string]bool

func New() Set {
	return Set{}
}

func NewFromSlice(l []string) Set {
	set := New()
	for _, v := range l {
		set.Add(v)
	}
	return set
}

func (s Set) String() string {
	out := []rune{'{'}
	for k := range s {
		elem := "\"" + k + "\", "
		out = append(out, []rune(elem)...)
	}
	if len(out) >= 3 {
		out = out[:len(out)-2]
	}
	out = append(out, '}')
	return string(out)
}

func (s Set) IsEmpty() bool {
	return len(s) == 0
}

func (s Set) Has(elem string) bool {
	return s[elem]
}

func (s Set) Add(elem string) {
	s[elem] = true
}

func Subset(s1, s2 Set) bool {
	for k := range s1 {
		if !s2[k] {
			return false
		}
	}
	return true
}

func Disjoint(s1, s2 Set) bool {
	for k := range s1 {
		if s2[k] {
			return false
		}
	}
	return true
}

func Equal(s1, s2 Set) bool {
	return Subset(s1, s2) && Subset(s2, s1)
}

func Intersection(s1, s2 Set) Set {
	s := New()
	for k := range s1 {
		if s2[k] {
			s.Add(k)
		}
	}
	return s
}

func Difference(s1, s2 Set) Set {
	s := New()
	for k := range s1 {
		if !s2[k] {
			s.Add(k)
		}
	}
	return s
}

func Union(s1, s2 Set) Set {
	s := New()
	for k := range s1 {
		s.Add(k)
	}
	for k := range s2 {
		s.Add(k)
	}
	return s
}
