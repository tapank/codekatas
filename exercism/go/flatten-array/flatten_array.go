package flatten

func Flatten(nested interface{}) []interface{} {
	out := make([]interface{}, 0)
	if value, ok := nested.([]interface{}); ok {
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
