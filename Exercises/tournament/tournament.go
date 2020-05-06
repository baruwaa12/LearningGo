package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type team struct {
	name  string
	mp    int
	wins  int
	draws int
	loses int
}

func (t *team) win() {
	t.wins++
}
func (t *team) loose() {
	t.loses++
}
func (t *team) draw() {
	t.draws++
}
func (t *team) playMatch() {
	t.mp++
}
func (t *team) points() int {
	return t.wins*3 + t.draws
}

//Tally is function for talying the results from a given output
func Tally(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	teams := map[string]*team{}
	var sb strings.Builder
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == byte('#') {
			continue
		}
		splitLine := strings.Split(line, ";")
		if len(splitLine) == 3 {
			if _, ok := teams[splitLine[0]]; !ok {
				teams[splitLine[0]] = &team{
					name: splitLine[0],
				}
			}
			if _, ok := teams[splitLine[1]]; !ok {
				teams[splitLine[1]] = &team{
					name: splitLine[1],
				}
			}
			teams[splitLine[0]].playMatch()
			teams[splitLine[1]].playMatch()
			switch splitLine[2] {
			case "win":
				{
					teams[splitLine[0]].win()
					teams[splitLine[1]].loose()
				}
			case "loss":
				{
					teams[splitLine[0]].loose()
					teams[splitLine[1]].win()
				}
			case "draw":
				{
					teams[splitLine[0]].draw()
					teams[splitLine[1]].draw()
				}
			default:
				return errors.New("not recognized result")
			}
		} else {
			if len(splitLine) < 3 {
				return errors.New("Not formated ok.")
			}
		}
	}
	sortedTeams := sortTeams(teams)
	sb.WriteString(fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P"))
	for _, team := range sortedTeams {
		sb.WriteString(fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", team.name, team.mp, team.wins, team.draws, team.loses, team.points()))
	}
	io.WriteString(w, sb.String())
	return nil
}

func sortTeams(teams map[string]*team) []*team {
	teamSlice := []*team{}
	for _, team := range teams {
		teamSlice = append(teamSlice, team)
	}
	sort.Slice(teamSlice, func(i, j int) bool {
		if teamSlice[i].points() != teamSlice[j].points() {
			return teamSlice[i].points() > teamSlice[j].points()
		}
		if teamSlice[i].mp != teamSlice[j].mp {
			return teamSlice[i].mp > teamSlice[j].mp
		}
		return teamSlice[i].name < teamSlice[j].name
	})
	return teamSlice
}
