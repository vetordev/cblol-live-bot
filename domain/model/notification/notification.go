package notification

import (
	"cblol-bot/domain/model/user"
	"time"
)

type Notification struct {
	Id           int64
	ScheduledFor time.Time
	Enable       bool
	User         *user.User
}

func New(id int64, schedule time.Time, enable bool, user *user.User) *Notification {
	return &Notification{id, schedule, enable, user}
}

const DefaultScheduledTime = "10:00:00"
