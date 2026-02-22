package models

import "time"

type Patient struct {
	ID        int64     `db:"id"`
	FullName  string    `db:"full_name"`
	Age       int       `db:"age"`
	Diagnosis string    `db:"diagnosis"`
	CreatedAt time.Time `db:"created_at"`
	DoctorID  int64     `db:"doctor_id"`
}
