package ranking

import (
	"cblol-bot/domain/model/team"
	"fmt"
	"strings"
)

type Ranking struct {
	teams []*team.Team
}

func (r *Ranking) FormatToString() string {

	var ranking []string

	for pos, team := range r.teams {
		ranking = append(ranking, fmt.Sprintf("%d - %s (%d-%d)", pos+1, team.Name, team.Wins, team.Losses))
	}

	return strings.Join(ranking, "\n")
}

func New(teams []*team.Team) *Ranking {
	return &Ranking{teams}
}
