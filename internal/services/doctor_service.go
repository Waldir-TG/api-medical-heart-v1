package services

import (
	"context"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/repositories"
	"github.com/google/uuid"
)

type DoctorService struct {
	doctorRepo *repositories.DoctorRepository
}

func NewDoctorService(doctorRepo *repositories.DoctorRepository) *DoctorService {
	return &DoctorService{doctorRepo: doctorRepo}
}

func (s *DoctorService) GetDoctors(ctx context.Context) ([]*models.DoctorDetailsResponse, error) {
	return s.doctorRepo.GetDoctors(ctx)
}

func (s *DoctorService) GetDoctorDetail(ctx context.Context, id uuid.UUID) (*models.DoctorDetailsResponse, error) {
	return s.doctorRepo.GetDoctorDetail(ctx, id)
}

func (s *DoctorService) CreateDoctor(ctx context.Context, doctor *models.DoctorCreateRequest) (uuid.UUID, error) {
	return s.doctorRepo.CreateDoctor(ctx, doctor)
}

func (s *DoctorService) UpdateDoctor(ctx context.Context, doctorID uuid.UUID, doctor *models.DoctorUpdateRequest) (string, error) {
	return s.doctorRepo.UpdateDoctor(ctx, doctorID, doctor)
}
