package week

import (
	"cblol-bot/domain/model/match"
	"cblol-bot/util/date"
	"fmt"
	"strings"
	"time"
)

type Week struct {
	matches []*match.Match
}

func findWeek(matches []*match.Match) int {

	now := time.Now()
	today := date.ResetHours(now)

	var week int

	for _, match := range matches {
		matchDay := date.ResetHours(*match.Schedule)

		if matchDay.After(today) || matchDay.Equal(today) {
			week = match.Week
			break
		}

	}

	return week
}

func filterMatchesByWeek(allMatches []*match.Match, week int) []*match.Match {
	var matches []*match.Match

	for _, m := range allMatches {
		if m.Week == week {
			matches = append(matches, m)
		}
	}

	return matches
}

func (w *Week) MatchesPerDay() string {

	matchDays := make(map[time.Time][]*match.Match)

	for _, m := range w.matches {
		matchDay := date.ResetHours(*m.Schedule)

		matchDays[matchDay] = append(matchDays[matchDay], m)
	}

	var week []string

	for day, matches := range matchDays {
		var formattedMatches []string

		for _, m := range matches {
			s := fmt.Sprintf("%s x %s - %d:%d", m.Team1.Name, m.Team2.Name, m.Schedule.Hour(), m.Schedule.Minute())

			if m.State == match.Completed {

			}

			formattedMatches = append(formattedMatches, s)
		}

		weekDay := date.GetWeekDayInPt(day.Weekday())
		j := strings.Join(formattedMatches, "\n")

		week = append(
			week,
			fmt.Sprintf("%s - %d/%d\n\n%s", weekDay, day.Day(), day.Month(), j),
		)
	}

	return strings.Join(week, "\n\n")
}

func New(allMatches []*match.Match) *Week {

	weekMatches := filterMatchesByWeek(allMatches, findWeek(allMatches))

	return &Week{weekMatches}
}
