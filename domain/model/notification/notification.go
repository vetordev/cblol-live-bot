package notification

import (
	"cblol-bot/domain/model/user"
	"time"
)

const DefaultScheduleTime = "10:00:00"

type Notification struct {
	Id           int
	ScheduledFor time.Time
	Enable       bool
	User         *user.User
}

func New(id int, schedule time.Time, enable bool, user *user.User) *Notification {
	return &Notification{id, schedule, enable, user}
}
