package series

func All(n int, s string) []string {
	in := []rune(s)
	if n < 1 || len(in) < n {
		return nil
	}

	all := []string{}
	for i := 0; i < len(in)+1-n; i++ {
		all = append(all, string(in[i:i+n]))
	}
	return all
}

func UnsafeFirst(n int, s string) string {
	all := All(n, s)
	if len(all) > 0 {
		return all[0]
	}
	return ""
}

func First(n int, s string) (first string, ok bool) {
	all := All(n, s)
	if len(all) > 0 {
		return all[0], true
	}
	return "", false
}
