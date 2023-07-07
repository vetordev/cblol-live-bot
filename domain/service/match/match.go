package match

import (
	"cblol-bot/domain/model/match"
	"cblol-bot/util/date"
	"fmt"
	"strings"
	"time"
)

func MatchesByDay(day time.Time, allMatches []*match.Match) string {

	var formattedMatches []string

	for _, m := range allMatches {
		formattedMatches = append(formattedMatches, m.Format())
	}

	weekDay := date.GetWeekDayInPt(day.Weekday())
	j := strings.Join(formattedMatches, "\n")

	return fmt.Sprintf("\n📆 %s - %d/%d\n\n%s", weekDay, day.Day(), day.Month(), j)
}
