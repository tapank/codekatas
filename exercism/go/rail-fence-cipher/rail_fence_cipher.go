package railfence

func Encode(message string, rails int) string {
	if rails < 2 || len(message) <= rails {
		return message
	}

	// spread the letters in rails
	rows := make([][]rune, rails)
	n, delta := 0, 1
	for _, r := range message {
		if rows[n] == nil {
			rows[n] = []rune{r}
		} else {
			rows[n] = append(rows[n], r)
		}

		if n == rails-1 {
			delta = -1
		} else if n == 0 {
			delta = 1
		}
		n += delta
	}

	// assemble encoded message
	msg := ""
	for _, row := range rows {
		msg += string(row)
	}

	return msg
}

func Decode(message string, rails int) string {
	if rails < 2 || len(message) <= rails {
		// nothing to do
		return message
	}

	// figure out the length of each rail
	railLengths := map[int]int{}
	n, delta := 0, 1
	for range message {
		railLengths[n]++

		if n == rails-1 {
			delta = -1
		} else if n == 0 {
			delta = 1
		}
		n += delta
	}

	// now split the message into rails
	rows := []*RuneStack{}
	pos := 0
	for i := 0; i < rails; i++ {
		l := railLengths[i]
		rows = append(rows, NewRuneStack([]rune(message)[pos:pos+l]))
		pos += l
	}

	// construct decoded message
	msg := []rune{}
	n, delta = 0, 1
	for {
		if r, ok := rows[n].Shift(); ok {
			msg = append(msg, r)
		} else {
			break
		}

		if n == rails-1 {
			delta = -1
		} else if n == 0 {
			delta = 1
		}
		n += delta
	}
	return string(msg)
}

type RuneStack struct {
	data []rune
	pos  int
	size int
}

func NewRuneStack(runes []rune) *RuneStack {
	return &RuneStack{data: runes, size: len(runes)}
}

func (rs *RuneStack) Shift() (r rune, ok bool) {
	if rs.pos == rs.size {
		ok = false
		return
	}
	r, ok = rs.data[rs.pos], true
	rs.pos++
	return
}
