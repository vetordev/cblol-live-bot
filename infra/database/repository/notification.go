package repository

import "database/sql"

type NotificationRepository struct {
	db *sql.DB
}

func (r *NotificationRepository) Create() {

}

func (r *NotificationRepository) Update() {

}

func (r *NotificationRepository) Find() {

}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db}
}
