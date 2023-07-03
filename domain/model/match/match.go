package match

import (
	"cblol-bot/domain/model/team"
	"fmt"
	"strconv"
	"time"
)

type State = int

const (
	Unstarted State = iota
	Completed
)

type Match struct {
	Schedule *time.Time
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

	s := fmt.Sprintf("<b>%s</b> X <b>%s</b> - %d:%s", m.Team1.Name, m.Team2.Name, hour, minute)

	if m.State == Completed {
		s += fmt.Sprintf(" (%s 🏆)", m.Winner.Name)
	}

	return s
}

func New(schedule *time.Time, block string, state State, team1 *team.Team, team2 *team.Team, winner *team.Team) *Match {
	return &Match{schedule, block, state, team1, team2, winner}
}
