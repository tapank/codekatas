package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Define the Robot type here.
type Robot struct {
	RobotName string
}

var index = 0
var names, maxIndex = func() ([]string, int) {
	allnames := make([]string, 0, 26*26*1000)
	var name string
	for ch1 := 'A'; ch1 <= 'Z'; ch1++ {
		for ch2 := 'A'; ch2 <= 'Z'; ch2++ {
			for d := 0; d <= 999; d++ {
				name = string(ch1) + string(ch2) + fmt.Sprintf("%03d", d)
				allnames = append(allnames, name)
			}
		}
	}
	// shuffle allnames to make it random
	rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Shuffle(len(allnames), func(i, j int) {
		allnames[i], allnames[j] = allnames[j], allnames[i]
	})
	return allnames, len(allnames) - 1
}()

func (r *Robot) Name() (string, error) {
	if r.RobotName != "" {
		return r.RobotName, nil
	}
	if index <= maxIndex {
		r.RobotName = names[index]
		index++
		return r.RobotName, nil
	}
	return "", fmt.Errorf("no more names available")
}

func (r *Robot) Reset() {
	r.RobotName = ""
}
