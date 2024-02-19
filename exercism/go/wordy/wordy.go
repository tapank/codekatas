package wordy

import (
	"strconv"
	"strings"
)

const (
	Plus     = "plus"
	Minus    = "minus"
	Multiply = "multiplied_by"
	Divide   = "divided_by"
)

var operations = map[string]func(v1, v2 int) (int, bool){
	Plus:     func(v1, v2 int) (int, bool) { return v1 + v2, true },
	Minus:    func(v1, v2 int) (int, bool) { return v1 - v2, true },
	Multiply: func(v1, v2 int) (int, bool) { return v1 * v2, true },
	Divide: func(v1, v2 int) (int, bool) {
		if v2 == 0 {
			return 0, false
		}
		return v1 / v2, true
	},
}

func Answer(question string) (int, bool) {
	// sanitize input
	q := strings.TrimPrefix(question, "What is ")
	q = strings.ReplaceAll(q, "multiplied by", "multiplied_by")
	q = strings.ReplaceAll(q, "divided by", "divided_by")
	q = strings.TrimRight(q, "?")

	var v1 int
	var s1, op, s2 string
	var err error

	// process first number and store result in v1
	// later we continue to use v1 for interim results
	s1, q, _ = strings.Cut(q, " ")
	if v1, err = strconv.Atoi(s1); err != nil {
		return 0, false
	}

	// process remaining expression
	for q != "" {
		var sep bool
		if op, q, sep = strings.Cut(q, " "); !sep {
			// no more arguments left after the operator, it is an error
			return 0, false
		}
		s2, q, _ = strings.Cut(q, " ")
		if v2, err := strconv.Atoi(s2); err != nil {
			return 0, false
		} else if operation, ok := operations[op]; ok {
			v1, ok = operation(v1, v2)
			if !ok {
				return 0, false
			}
		} else {
			// unsupported operation
			return 0, false
		}
	}
	return v1, true
}
