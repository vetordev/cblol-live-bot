package notification

import (
	"cblol-bot/domain/model/user"
	"errors"
)

var NotFoundByUser = errors.New("notification not found for this user")
var CouldNotCreate = errors.New("could not create a new notification")
var CouldNotUpdate = errors.New("could not update this notification")

type Repository interface {
	Create(string, bool, *user.User) (int64, error)
	Update(*Notification) error
	List() ([]*Notification, error)
	FindByUser(*user.User) (*Notification, error)
}
