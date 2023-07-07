package notification

import "cblol-bot/domain/model/match"

type Service struct {
	matches []*match.Match
}

func (s *Service) ScheduleNotification() {

}

func New(matches []*match.Match) *Service {
	return &Service{matches}
}
