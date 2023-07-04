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

	alertJob := job.NewAlert()
	s.cron.AddJob("a", alertJob)
}

func New() *Scheduler {
	location, _ := time.LoadLocation("America/Sao_Paulo")

	c := cron.New(cron.WithLocation(location))

	return &Scheduler{c}
}
