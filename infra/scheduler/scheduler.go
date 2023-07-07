package scheduler

import (
	"cblol-bot/infra/scheduler/job"
	"github.com/robfig/cron/v3"
	"time"
)

type Scheduler struct {
	cron *cron.Cron
}

func (s *Scheduler) Load() {

	scheduleNotification := job.NewScheduleNotification()

	s.cron.AddJob("@midnight", scheduleNotification)
}

func New() *Scheduler {
	location, _ := time.LoadLocation("America/Sao_Paulo")

	c := cron.New(cron.WithLocation(location))

	return &Scheduler{c}
}
