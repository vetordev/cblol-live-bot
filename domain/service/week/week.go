package week

import (
	"cblol-bot/domain/model/match"
	matchsvc "cblol-bot/domain/service/match"
	"cblol-bot/util/date"
	"strings"
	"time"
)

type Week struct {
}

func (w *Week) findBlock(matches []*match.Match) string {

	now := time.Now()
	today := date.ResetHours(now)

	var blockName string

	for _, m := range matches {
		matchDay := date.ResetHours(m.Schedule)

		if matchDay.After(today) || matchDay.Equal(today) {
			blockName = m.Block
			break
		}

	}

	return blockName
}

func (w *Week) filterMatchesByBlock(allMatches []*match.Match, block string) []*match.Match {
	var matches []*match.Match

	for _, m := range allMatches {
		if m.Block == block {
			matches = append(matches, m)
		}
	}

	return matches
}

func (w *Week) FormatWeekMatches(allMatches []*match.Match) string {

	var formattedMatches []string

	location, _ := time.LoadLocation("America/Sao_Paulo")
	lastMonday := date.LastMonday(time.Now().In(location))

	day := date.ResetHours(lastMonday)

	for {

		var matches []*match.Match

		for _, m := range allMatches {
			if date.ResetHours(m.Schedule).Equal(day) {
				matches = append(matches, m)
			}
		}

		if len(matches) > 0 {
			formattedMatches = append(formattedMatches, matchsvc.FormatMatchesPerDay(day, matches))
		}

		if day.Weekday() == 0 {
			break
		}
		day = day.AddDate(0, 0, 1)
	}

	return strings.Join(formattedMatches, "\n")
}

func New() *Week {

	return &Week{}
}
