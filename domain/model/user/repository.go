package user

import "errors"

var CouldNotCreate = errors.New("Could not crate a new user")

type Repository interface {
	Create(int64, string) error
	Exists(int64) bool
}
