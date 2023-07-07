package notification

import (
	"cblol-bot/application/match"
	"cblol-bot/domain/service/notification"
	"fmt"
)

type Application struct {
	matchApplication *match.Application
}

func (a *Application) ScheduleDailyNotificationOfMatches() {
	matches, err := a.matchApplication.ListTodayMatchesFromAPI()

	if err != nil {
		fmt.Println(err)
		return
	}

	if len(matches) == 0 {
		return
	}

	notificationService := notification.New(matches)

	notificationService.ScheduleNotification()

}

func New(matchApplication *match.Application) *Application {
	return &Application{matchApplication}
}
