package notification

import (
	"cblol-bot/domain/model/match"
	"cblol-bot/domain/model/user"
	"time"
)

type Service struct {
}

func (s *Service) ScheduleNotifications(matches []*match.Match) {

}

func (s *Service) EnableMatchNotification(user *user.User, time *time.Time) {

}

func New() *Service {
	return &Service{}
}
