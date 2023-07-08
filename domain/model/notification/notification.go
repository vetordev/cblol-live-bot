package notification

import (
	"cblol-bot/domain/model/user"
	"time"
)

type Notification struct {
	Id       int
	Schedule time.Time
	Enable   bool
	User     *user.User
}

func New() *Notification {
	return &Notification{}
}
