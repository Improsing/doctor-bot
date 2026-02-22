package models

import "time"

type Doctor struct {
	ID             int64     `db:"id"`
	TelegramID     int64     `db:"telegram_id"`
	FullName       string    `db:"full_name"`
	Specialization string    `db:"specialization"`
	CreatedAt      time.Time `db:"created_at"`
}
