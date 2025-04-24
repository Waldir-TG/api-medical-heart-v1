package services

import (
	"context"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/repositories"
	"github.com/google/uuid"
)

type PatientService struct {
	patientRepo *repositories.PatientRepo
}

func NewPatientService(patientRepo *repositories.PatientRepo) *PatientService {
	return &PatientService{
		patientRepo: patientRepo,
	}
}

func (s *PatientService) GetAllPatientsBasicDetails(ctx context.Context) ([]*models.PatientDetailBasicResponse, error) {
	return s.patientRepo.GetPatientsBasicDetail(ctx, nil)
}

func (s *PatientService) GetPatientBasicDetailsByID(ctx context.Context, patientID *uuid.UUID) ([]*models.PatientDetailBasicResponse, error) {
	return s.patientRepo.GetPatientsBasicDetail(ctx, patientID)
}

func (s *PatientService) GetPatientsDetails(ctx context.Context) ([]*models.PatientDetailResponse, error) {
	return s.patientRepo.GetPatientsDetails(ctx, nil)
}

func (s *PatientService) GetPatientDetailsByID(ctx context.Context, patientID *uuid.UUID) ([]*models.PatientDetailResponse, error) {
	return s.patientRepo.GetPatientsDetails(ctx, patientID)
}

func (s *PatientService) CreatePatient(ctx context.Context, patient *models.PatientCreateRequest) (uuid.UUID, error) {
	return s.patientRepo.CreatePatient(ctx, patient)
}

func (s *PatientService) UpdatePatient(ctx context.Context, patientID uuid.UUID, patient *models.PatientUpdatedRequest) (string, error) {
	return s.patientRepo.UpdatePatient(ctx, patientID, patient)
}

func (s *PatientService) AssignDoctorToPatient(ctx context.Context, DoctorPatientAssign *models.DoctorPatientAssignRequest) (string, error) {
	return s.patientRepo.AssignDoctorToPatient(ctx, DoctorPatientAssign)
}

// Additional methods can be added as needed, such as:
// - GetPatientByID
// - GetAllPatients
// - DeletePatient
// - GetPatientsByDoctorID
