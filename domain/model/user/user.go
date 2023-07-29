package user

type User struct {
	ChatId int64
	Name   string
}

func New(chatId int64, name string) *User {
	return &User{chatId, name}
}
