package secret

import "slices"

// Handshake can be optimized, but premature optimization is the root of evil
func Handshake(code uint) []string {
	actions := []string{}
	if 0b00001&code != 0 {
		actions = append(actions, "wink")
	}
	if 0b00010&code != 0 {
		actions = append(actions, "double blink")
	}
	if 0b00100&code != 0 {
		actions = append(actions, "close your eyes")
	}
	if 0b01000&code != 0 {
		actions = append(actions, "jump")
	}
	if 0b10000&code != 0 {
		slices.Reverse(actions)
	}
	return actions
}
