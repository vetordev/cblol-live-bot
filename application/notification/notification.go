package notification

import (
	"cblol-bot/application/match"
	"cblol-bot/domain/model/notification"
	"cblol-bot/domain/model/user"
	notificationsvc "cblol-bot/domain/service/notification"
	"errors"
	"fmt"
	"time"
)

const NotificationsEnabled = "Notificações habilitadas!"

type Application struct {
	matchApplication       *match.Application
	userRepository         user.Repository
	notificationRepository notification.Repository
	notificationService    *notificationsvc.Service
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

	notifications, err := a.notificationRepository.List()

	if err != nil {
		return
	}

	a.notificationService.ScheduleNotifications(matches, notifications)
}

func (a *Application) EnableDailyNotificationOfMatches(chatId int64, userName string, scheduledTime string) string {

	scheduledTime = toSchedulingFormat(scheduledTime)

	if scheduledTime == "" {
		return InvalidScheduledTime
	}

	if exists := a.userRepository.Exists(chatId); !exists {

		err := a.userRepository.Create(chatId, userName)

		if err != nil {
			fmt.Println(err)

			return CouldNotEnableNotifications
		}
	}

	u := user.New(chatId, userName)

	n, err := a.notificationRepository.FindByUser(u)

	if err != nil {
		if !errors.Is(err, notification.NotFoundByUser) {
			return CouldNotEnableNotifications
		}

		_, err = a.notificationRepository.Create(scheduledTime, true, u)

		if err != nil {
			return CouldNotEnableNotifications
		}

		return NotificationsEnabled
	}

	scheduledFor, err := time.Parse(time.TimeOnly, scheduledTime)

	n.ScheduledFor = scheduledFor

	err = a.notificationRepository.Update(n)

	if err != nil {
		return CouldNotEnableNotifications
	}

	return NotificationsEnabled
}

func New(
	matchApplication *match.Application,
	userRepository user.Repository,
	notificationRepository notification.Repository,
	notificationService *notificationsvc.Service,
) *Application {
	return &Application{
		matchApplication,
		userRepository,
		notificationRepository,
		notificationService,
	}
}
