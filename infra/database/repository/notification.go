package repository

import (
	"cblol-bot/domain/model/notification"
	"cblol-bot/domain/model/user"
	"database/sql"
	"fmt"
	"time"
)

type NotificationRepository struct {
	db *sql.DB
}

func (r *NotificationRepository) Create(scheduledFor string, enable bool, u *user.User) (int64, error) {
	stmt, err := r.db.Prepare("INSERT INTO notifications (scheduled_for, enable, user_id) VALUES(?, ?, ?)")

	if err != nil {
		fmt.Println(err)
		return 0, notification.CouldNotCreate
	}

	defer stmt.Close()

	result, err := stmt.Exec(scheduledFor, enable, u.ChatId)

	if err != nil {
		fmt.Println(err)
		return 0, notification.CouldNotCreate
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Println(err)
		return 0, notification.CouldNotCreate
	}

	return id, nil
}

func (r *NotificationRepository) Update(n *notification.Notification) error {
	stmt, err := r.db.Prepare("UPDATE notifications SET scheduled_for = ?, enable = ?, user_id = ? WHERE id = ?")

	if err != nil {
		fmt.Println(err)
		return notification.CouldNotUpdate
	}

	defer stmt.Close()

	scheduledFor := n.ScheduledFor.Format(time.TimeOnly)

	_, err = stmt.Exec(scheduledFor, n.Enable, n.User.ChatId, n.Id)

	if err != nil {
		fmt.Println(err)
		return notification.CouldNotUpdate
	}

	return nil
}

func (r *NotificationRepository) List() ([]*notification.Notification, error) {
	var notifications []*notification.Notification

	stmt, err := r.db.Prepare(`
		SELECT
			id,
			schedule_for,
			enable,
			user.chat_id,
			user.name
		FROM notifications
		INNER JOIN users ON notifications.user_id = users.chat_id
	`)

	if err != nil {
		fmt.Println(err)
		return nil, notification.CouldNotList
	}

	defer stmt.Close()

	rows, err := stmt.Query()

	if err != nil {
		if err == sql.ErrNoRows {
			return notifications, nil
		}

		fmt.Println(err)
		return notifications, err
	}

	for rows.Next() {
		var n notification.Notification
		var u user.User

		var schedule string
		err = rows.Scan(&n.Id, &schedule, &n.Enable, &u.ChatId, &u.Name)

		n.ScheduledFor, _ = time.Parse(time.TimeOnly, schedule)
		n.User = &u

		notifications = append(notifications, &n)
	}

	return notifications, nil
}

func (r *NotificationRepository) FindByUser(u *user.User) (*notification.Notification, error) {
	stmt, err := r.db.Prepare("SELECT id, scheduled_for, enable from notifications WHERE user_id = ?")

	if err != nil {
		fmt.Println(err)
		return nil, notification.NotFoundByUser
	}

	defer stmt.Close()

	var n notification.Notification

	n.User = u
	var scheduleTime string

	err = stmt.QueryRow(u.ChatId).Scan(&n.Id, &scheduleTime, &n.Enable)

	if err != nil {
		fmt.Println(err)
		return nil, notification.NotFoundByUser
	}

	n.ScheduledFor, _ = time.Parse(time.TimeOnly, scheduleTime)

	return &n, nil
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{db}
}
