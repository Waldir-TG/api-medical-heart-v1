package services

import (
	"context"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/repositories"
)

type DoctorService struct {
	doctorRepo *repositories.DoctorRepository
}

func NewDoctorService(doctorRepo *repositories.DoctorRepository) *DoctorService {
	return &DoctorService{doctorRepo: doctorRepo}
}

func (s *DoctorService) GetDoctors(ctx context.Context) ([]*models.Doctor, error) {
	return s.doctorRepo.GetDoctors(ctx)
}
