package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	plants map[string][]string
}

var codes = map[byte]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if len(diagram) == 0 || diagram[0] != '\n' {
		return nil, errors.New("empty diagram")
	}

	rows := strings.Split(strings.Trim(diagram, "\n"), "\n")
	if len(rows) != 2 || // must have two rows
		len(rows[0]) != len(rows[1]) || // both rows should be same length
		len(rows[0])%2 != 0 || // must have even number of elements
		len(children)*2 != len(rows[0]) { // rows must be twice the length of children
		return nil, errors.New("bad input")
	}

	// sort children
	c := make([]string, len(children))
	copy(c, children)
	sort.Strings(c)

	garden := Garden{make(map[string][]string)}
	for i := 0; i < len(c); i++ {
		c1, c2, c3, c4 := codes[rows[0][i*2]],
			codes[rows[0][i*2+1]],
			codes[rows[1][i*2]],
			codes[rows[1][i*2+1]]
		if len(c1) == 0 || len(c2) == 0 || len(c3) == 0 || len(c4) == 0 {
			return nil, errors.New("bad plant")
		}
		p := []string{c1, c2, c3, c4}
		child := c[i]
		if _, exits := garden.plants[child]; exits {
			return nil, errors.New("duplicate child")
		}
		garden.plants[c[i]] = p
	}
	return &garden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.plants[child]
	return plants, ok
}
