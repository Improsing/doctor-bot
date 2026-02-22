package service

import (
	"doctor-bot/internal/repository"
)

type PatientService struct {
	doctorRepo  *repository.DoctorRepository
	patientRepo *repository.PatientRepository
}

func NewPatientService(
	doctorRepo *repository.DoctorRepository,
	patientRepo *repository.PatientRepository,
) *PatientService {
	return &PatientService{
		doctorRepo:  doctorRepo,
		patientRepo: patientRepo}
}

func (s *PatientService) DeletePatient(tgID int64, patientID int64) error {

}
