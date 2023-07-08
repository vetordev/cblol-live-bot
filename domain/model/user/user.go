package user

type User struct {
	Id     int
	ChatId string
	Name   string
}

func New(id int, chatId string, name string) *User {
	return &User{id, chatId, name}
}
