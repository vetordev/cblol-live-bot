package scheduler

import (
	"cblol-bot/infra/scheduler/job"
	"github.com/robfig/cron/v3"
	"time"
)

type Scheduler struct {
	cron *cron.Cron

	scheduleNotificationJob *job.ScheduleNotification
}

func (s *Scheduler) Load() {

	s.cron.AddJob("@midnight", s.scheduleNotificationJob)

	s.cron.Start()
}

func New(scheduleNotificationJob *job.ScheduleNotification) *Scheduler {
	location, _ := time.LoadLocation("America/Sao_Paulo")

	c := cron.New(cron.WithLocation(location))

	return &Scheduler{c, scheduleNotificationJob}
}
