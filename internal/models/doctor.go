package models

type Doctor struct {
	ID             int64  `db:"id"`
	TelegramID     int64  `db:"telegram_id"`
	FullName       string `db:"full_name"`
	Specialization string `db:"specialization"`
}
