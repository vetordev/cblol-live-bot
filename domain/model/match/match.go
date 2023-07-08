package match

import (
	"cblol-bot/domain/model/team"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type State = int

const (
	Unstarted State = iota
	Completed
)

type Match struct {
	Schedule time.Time
	Block    string
	State    State

	Team1 *team.Team
	Team2 *team.Team

	Winner *team.Team
}

func (m *Match) Format() string {

	hour := m.Schedule.Hour()

	minute := strconv.Itoa(m.Schedule.Minute())

	if len(minute) < 2 {
		minute = "0" + minute
	}

	team1ShortName := getShortName(m.Team1.Name)
	team2ShortName := getShortName(m.Team2.Name)

	s := fmt.Sprintf("<b>%s</b> X <b>%s</b> - %d:%s", team1ShortName, team2ShortName, hour, minute)

	if m.State == Completed {
		s += fmt.Sprintf(" (%s 🏆)", getShortName(m.Winner.Name))
	}

	return s
}

func getShortName(name string) string {

	split := strings.Split(name, " ")

	shortName := split[0]

	if len(split) > 2 {
		shortName += " " + split[1]
	}

	return shortName
}

func New(schedule time.Time, block string, state State, team1 *team.Team, team2 *team.Team, winner *team.Team) *Match {
	return &Match{schedule, block, state, team1, team2, winner}
}
