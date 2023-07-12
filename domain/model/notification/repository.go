package notification

import (
	"cblol-bot/domain/model/user"
	"errors"
	"time"
)

var NotFoundByUser = errors.New("notification not found for this user")
var CouldNotCreate = errors.New("could not create a new notification")
var CouldNotUpdate = errors.New("could not update this notification")

type Repository interface {
	Create(time.Time, bool, *user.User) (int64, error)
	Update(*Notification) error
	FindByUser(*user.User) (*Notification, error)
}
