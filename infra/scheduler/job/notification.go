package job

import "cblol-bot/application/notification"

type ScheduleNotification struct {
	notificationApplication *notification.Application
}

func (j *ScheduleNotification) Run() {
	j.notificationApplication.ScheduleDailyNotificationOfMatches()
}

func NewScheduleNotification(notificationApplication notification.Application) *ScheduleNotification {
	return &ScheduleNotification{}
}
