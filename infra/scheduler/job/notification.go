package job

import (
	"cblol-bot/application/notification"
	"cblol-bot/infra/scheduler"
)

type ScheduleNotification struct {
	notificationApplication *notification.Application
	scheduler               *scheduler.Scheduler
}

func (j *ScheduleNotification) Schedule() {
	j.scheduler.Add("@midnight", func() {
		j.scheduler.RemoveAll()

		j.notificationApplication.ScheduleDailyNotificationOfMatches()
	})
}

func NewJobNotification(notificationApplication *notification.Application, scheduler *scheduler.Scheduler) *ScheduleNotification {
	return &ScheduleNotification{notificationApplication, scheduler}
}
