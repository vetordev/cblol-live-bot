package match

import (
	"cblol-bot/domain/model/match"
	"cblol-bot/util/date"
	"fmt"
	"strings"
	"time"
)

func FormatMatchesPerDay(day time.Time, matches []*match.Match) string {

	var formattedMatches []string

	for _, m := range matches {
		formattedMatches = append(formattedMatches, m.Format())
	}

	weekDay := date.GetWeekDayInPt(day.Weekday())
	j := strings.Join(formattedMatches, "\n")

	return fmt.Sprintf("\n📆 %s - %d/%d (<b>%s</b>)\n\n%s", weekDay, day.Day(), day.Month(), matches[0].Block, j)
}
