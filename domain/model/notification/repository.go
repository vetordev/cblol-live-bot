package notification

import (
	"cblol-bot/domain/model/user"
	"errors"
	"time"
)

var NotFoundByUser = errors.New("notification not found for this user")

type Repository interface {
	Create(time.Time, bool, *user.User) (int, error)
	Update(*Notification) error
	Find(int) *Notification
	FindByUser(*user.User) (*Notification, error)
}
