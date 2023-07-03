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
	block   string
	matches []*match.Match
}

func findBlock(matches []*match.Match) string {

	now := time.Now()
	today := date.ResetHours(now)

	var blockName string

	for _, m := range matches {
		matchDay := date.ResetHours(*m.Schedule)

		if matchDay.After(today) || matchDay.Equal(today) {
			blockName = m.Block
			break
		}

	}

	return blockName
}

func filterMatchesByBlock(allMatches []*match.Match, block string) []*match.Match {
	var matches []*match.Match

	for _, m := range allMatches {
		if m.Block == block {
			matches = append(matches, m)
		}
	}

	return matches
}

func (w *Week) Matches() string {

	matchDays := make(map[time.Time][]*match.Match)

	for _, m := range w.matches {
		matchDay := date.ResetHours(*m.Schedule)

		matchDays[matchDay] = append(matchDays[matchDay], m)
	}

	var weekMatches []string

	for day, matches := range matchDays {
		weekMatches = append(weekMatches, matchsvc.MatchesPerDay(day, matches))
	}

	formattedWeek := strings.Join(weekMatches, "\n\n")

	return fmt.Sprintf("<b>%s</b>\n", w.block) + formattedWeek
}

func New(allMatches []*match.Match) *Week {

	block := findBlock(allMatches)
	weekMatches := filterMatchesByBlock(allMatches, block)

	return &Week{block, weekMatches}
}
