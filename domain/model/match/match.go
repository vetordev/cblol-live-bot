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
	Schedule *time.Time
	Week     int
	State    State

	Team1 *team.Team
	Team2 *team.Team

	Winner *team.Team
}

func New(schedule *time.Time, week int, state State, team1 *team.Team, team2 *team.Team, winner *team.Team) *Match {
	return &Match{schedule, week, state, team1, team2, winner}
}
