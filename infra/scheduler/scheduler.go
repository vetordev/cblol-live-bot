package scheduler

import (
	"github.com/robfig/cron/v3"
	"time"
)

type Scheduler struct {
	cron *cron.Cron
}

func (s *Scheduler) RemoveAll() {
	jobs := s.cron.Entries()

	for _, entry := range jobs {
		s.cron.Remove(entry.ID)
	}
}

func (s *Scheduler) Add(spec string, cmd func()) {
	s.cron.AddFunc(spec, cmd)
}

func New() *Scheduler {
	location, _ := time.LoadLocation("America/Sao_Paulo")

	c := cron.New(cron.WithLocation(location))

	c.Start()

	return &Scheduler{c}
}
