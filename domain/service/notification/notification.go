package notification

import (
	"cblol-bot/domain/model/match"
	"cblol-bot/domain/model/notification"
	matchsvc "cblol-bot/domain/service/match"
	"fmt"
)

type Service struct {
	scheduler Scheduler
	notifier  Notifier
}

func (s *Service) ScheduleNotifications(matches []*match.Match, notifications []*notification.Notification) {

	formattedMatches := matchsvc.FormatMatchesPerDay(matches[0].Schedule, matches)

	for _, n := range notifications {

		text := fmt.Sprintf("Olá, %s. Segue os jogos de hoje:\n%s", n.User.Name, formattedMatches)

		spec := fmt.Sprintf("%d %d %d * * *", n.ScheduledFor.Second(), n.ScheduledFor.Minute(), n.ScheduledFor.Hour())

		chatID := n.User.ChatId

		s.scheduler.Add(spec, func() {
			s.notifier.Send(chatID, text)
		})
	}

}

func New(scheduler Scheduler, notifier Notifier) *Service {
	return &Service{scheduler, notifier}
}

type Scheduler interface {
	Add(string, func())
}

type Notifier interface {
	Send(chatId int64, text string)
}
