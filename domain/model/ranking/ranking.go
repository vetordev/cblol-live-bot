package ranking

import (
	"cblol-bot/domain/model/team"
	"fmt"
	"strings"
)

type Ranking struct {
	teams []*team.Team
}

func getIcon(i int) string {
	icons := []string{"🥇", "🥈", "🥉", "4⃣", "5⃣", "6⃣", "7⃣", "8⃣", "9⃣", "🔟"}

	return icons[i]
}

func (r *Ranking) Format() string {

	var ranking []string

	for pos, team := range r.teams {
		formatted := fmt.Sprintf("%s - <b>%s</b> (%dv-%dd)", getIcon(pos), team.Name, team.Wins, team.Losses)
		ranking = append(ranking, formatted)
	}

	return strings.Join(ranking, "\n")
}

func New(teams []*team.Team) *Ranking {
	return &Ranking{teams}
}
