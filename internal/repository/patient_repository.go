package repository

import (
	"doctor-bot/internal/models"

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
		INSERT INTO patients (full_name, age, diagnosis)
		VALUES ($1, $2, $3)
	`, patient.FullName, patient.Age, patient.Diagnosis)
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

func (h *PatientRepository) DeleteByID(id int) error {
	_, err := h.db.Exec(`
		DELETE FROM patients WHERE id=$1`, id)
	return err
}
