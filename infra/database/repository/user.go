package repository

import (
	"cblol-bot/domain/model/user"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func (r *UserRepository) Create(chatId int64, name string) error {

	stmt, err := r.db.Prepare("INSERT INTO users(chat_id, name) VALUES(?, ?)")

	if err != nil {
		fmt.Println(err)
		return user.CouldNotCreate
	}

	defer stmt.Close()

	_, err = stmt.Exec(chatId, name)

	if err != nil {
		fmt.Println(err)
		return user.CouldNotCreate
	}

	return nil
}

func (r *UserRepository) Exists(chatId int64) bool {
	stmt, err := r.db.Prepare("SELECT EXISTS(SELECT chat_id FROM users WHERE chat_id = ?)")

	if err != nil {
		fmt.Println(err)
		return false
	}

	var exists int
	err = stmt.QueryRow(chatId).Scan(&exists)

	return exists != 0
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}
