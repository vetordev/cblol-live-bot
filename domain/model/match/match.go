package match

import (
	"cblol-bot/domain/model/team"
	"time"
)

type State = int

const (
	Unstarted State = iota
	Completed
)

type Match struct {
	schedule time.Time
	state    State

	team1 team.Team
	team2 team.Team
}

func (m *Match) Winner() *team.Team {
	return nil
}

func (m *Match) Loser() *team.Team {
	return nil
}

func New(schedule time.Time, state State, team1 team.Team, team2 team.Team) *Match {
	return &Match{schedule, state, team1, team2}
}
