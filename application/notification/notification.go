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

const CouldNotEnableNotifications = "Oops! Não foi possível habilitar as notificações"
const NotificationsEnabled = "Notificações habilitadadas!"

type Application struct {
	matchApplication       *match.Application
	userRepository         user.Repository
	notificationRepository notification.Repository
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

	notificationService := notificationsvc.New()

	notificationService.ScheduleNotifications(matches)
}

func (a *Application) EnableDailyNotificationOfMatches(chatId int64, userName string, notificationTime string) string {

	if notificationTime == "" {
		notificationTime = notification.DefaultScheduleTime
	}

	if exists := a.userRepository.Exists(chatId); !exists {

		err := a.userRepository.Create(chatId, userName)

		if err != nil {
			fmt.Println(err)

			return CouldNotEnableNotifications
		}
	}

	u := user.New(chatId, userName)

	scheduledFor, err := time.Parse(time.TimeOnly, notificationTime)

	n, err := a.notificationRepository.FindByUser(u)

	if err != nil {
		if !errors.Is(err, notification.NotFoundByUser) {
			return CouldNotEnableNotifications
		}

		_, err := a.notificationRepository.Create(scheduledFor, true, u)

		if err != nil {
			return CouldNotEnableNotifications
		}

		return NotificationsEnabled
	}

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
) *Application {
	return &Application{matchApplication, userRepository, notificationRepository}
}
