package tournament

import (
	"bufio"
	"cmp"
	"fmt"
	"io"
	"slices"
	"strings"
)

// Team is a team type where,
// Name: team name
// MP: Matches Played
// W: Matches Won
// D: Matches Drawn (Tied)
// L: Matches Lost
// P: Points
type Team struct {
	Name           string
	MP, W, D, L, P int
}

const teamfmt = "%-30s | %2d | %2d | %2d | %2d | %2d"

func NewTeam(name string) *Team {
	return &Team{Name: name}
}

func (t *Team) String() string {
	return fmt.Sprintf(teamfmt, t.Name, t.MP, t.W, t.D, t.L, t.P)
}

func Tally(reader io.Reader, writer io.Writer) error {
	scanner := bufio.NewScanner(reader)
	teams := map[string]*Team{}
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		// fmt.Printf("Got line: ***%s***\n", line)
		tokens := strings.Split(line, ";")
		if len(tokens) != 3 {
			return fmt.Errorf("invalid line: '%s'", line)
		}
		var team1, team2 *Team
		if team1 = teams[tokens[0]]; team1 == nil {
			team1 = NewTeam(tokens[0])
			teams[team1.Name] = team1
		}
		if team2 = teams[tokens[1]]; team2 == nil {
			team2 = NewTeam(tokens[1])
			teams[team2.Name] = team2
		}
		outcome := tokens[2]
		team1.MP++
		team2.MP++
		switch outcome {
		case "win":
			team1.W++
			team1.P += 3
			team2.L++
		case "loss":
			team1.L++
			team2.W++
			team2.P += 3
		case "draw":
			team1.D++
			team1.P++
			team2.D++
			team2.P++
		default:
			return fmt.Errorf("invalid outcome: '%s'", outcome)
		}
	}
	writeresult(teamslice(teams), writer)
	return nil
}

func writeresult(sortedTeams []Team, writer io.Writer) {
	w := bufio.NewWriter(writer)
	w.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	for _, t := range sortedTeams {
		w.WriteString(t.String())
		w.WriteRune('\n')
	}
	w.Flush()
}

// Sort is ordered by points, descending.
// In case of a tie, teams are ordered alphabetically.
func teamslice(teams map[string]*Team) []Team {
	tslice := []Team{}
	for _, v := range teams {
		tslice = append(tslice, *v)
	}
	slices.SortFunc(tslice, func(a, b Team) int {
		if n := cmp.Compare(b.P, a.P); n != 0 {
			return n
		}
		// If points are equal, order by name
		return cmp.Compare(a.Name, b.Name)
	})
	return tslice
}
