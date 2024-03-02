package flatten

func Flatten(nested interface{}) []interface{} {
	out := make([]interface{}, 0)
	switch value := nested.(type) {
	case int:
		out = append(out, value)
	case []interface{}:
		for _, v := range value {
			switch v.(type) {
			case int:
				out = append(out, v)
			case []interface{}:
				out = append(out, Flatten(v)...)
			}
		}
	}
	return out
}
