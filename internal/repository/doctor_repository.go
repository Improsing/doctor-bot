package repository

import (
	"doctor-bot/internal/models"

	"github.com/jmoiron/sqlx"
)

type DoctorRepository struct {
	db *sqlx.DB
}

func NewDoctorRepository(db *sqlx.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) GetByTelegramID(tgID int64) (*models.Doctor, error) {
	var doctor models.Doctor
	err := r.db.Get(&doctor, "SELECT * FROM doctors WHERE telegram_id=$1", tgID)
	if err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *DoctorRepository) Create(doctor *models.Doctor) error {
	_, err := r.db.Exec(`
		INSERT INTO doctors (telegram_id, full_name, specialization)
		VALUES ($1, $2, $3)
	`, doctor.TelegramID, doctor.FullName, doctor.Specialization)
	return err
}
