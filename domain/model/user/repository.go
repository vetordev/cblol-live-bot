package user

type Repository interface {
	Create(int64, string) error
	Update(*User) error
	Exists(int64) bool
}
