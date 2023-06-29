package week

import "cblol-bot/domain/model/match"

type Week struct {
	matches []*match.Match
}

func (w *Week) MatchesPerDay() {

}

func New(matches []*match.Match) *Week {
	return &Week{matches}
}
