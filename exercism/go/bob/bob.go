package bob

import "strings"

// Hey responds to remarks in weird but well known ways.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion := len(remark) > 0 && remark[len(remark)-1] == '?'
	hasCaps, hasLower, hasNum := false, false, false
	for _, r := range remark {
		if r >= 'A' && r <= 'Z' {
			hasCaps = true
		} else if r >= 'a' && r <= 'z' {
			hasLower = true
		} else if r >= '0' && r <= '9' {
			hasNum = true
		}
	}

	// greet now
	if hasCaps && !hasLower {
		if isQuestion {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	}
	if isQuestion {
		return "Sure."
	}
	if !hasCaps && !hasLower && !hasNum {
		return "Fine. Be that way!"
	}
	return "Whatever."
}
