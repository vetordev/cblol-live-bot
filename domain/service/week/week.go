package week

import (
	"cblol-bot/domain/model/match"
	matchsvc "cblol-bot/domain/service/match"
	"cblol-bot/util/date"
	"fmt"
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

	block := w.findBlock(allMatches)
	weekMatches := w.filterMatchesByBlock(allMatches, block)

	matchDays := make(map[time.Time][]*match.Match)

	for _, m := range weekMatches {
		matchDay := date.ResetHours(m.Schedule)

		matchDays[matchDay] = append(matchDays[matchDay], m)
	}

	var formattedMatches []string

	for day, matches := range matchDays {
		formattedMatches = append(formattedMatches, matchsvc.FormatMatchesPerDay(day, matches))
	}

	formattedWeek := strings.Join(formattedMatches, "\n\n")

	return fmt.Sprintf("<b>%s</b>\n", block) + formattedWeek
}

func New() *Week {

	return &Week{}
}
