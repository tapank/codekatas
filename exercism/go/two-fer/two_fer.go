package twofer

import "fmt"

// ShareWith returns what to say when giving a cookie
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
