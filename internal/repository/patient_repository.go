package repository

import (
	"doctor-bot/internal/models"
	"errors"

	"github.com/jmoiron/sqlx"
)

type PatientRepository struct {
	db *sqlx.DB
}

func NewPatientRepository(db *sqlx.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (h *PatientRepository) Create(patient *models.Patient) error {
	_, err := h.db.Exec(`
		INSERT INTO patients (full_name, age, diagnosis, doctor_id)
		VALUES ($1, $2, $3, $4)
	`, patient.FullName, patient.Age, patient.Diagnosis, patient.DoctorID)
	return err
}

func (h *PatientRepository) GetAll() ([]models.Patient, error) {
	var patients []models.Patient

	err := h.db.Select(&patients, "SELECT * FROM patients ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}

	return patients, nil

}

func (h *PatientRepository) DeleteByID(id int64, doctorID int64) error {
	result, err := h.db.Exec(`
		DELETE FROM patients WHERE id= $1 AND doctor_id = $2`, id, doctorID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("patient not found or access denied")
	}

	return nil
}
