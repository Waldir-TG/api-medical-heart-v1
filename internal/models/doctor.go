package models

import (
	"time"

	"github.com/google/uuid"
)

// Doctor model info
// @Description Información de un médico
type Doctor struct {
	ID                 uuid.UUID `json:"id" example:"Dr. House"`
	UserId             uuid.UUID `json:"user_id" example:"Dr. House"`
	IdentityType       string    `json:"identity_type" example:"Dr. House"`
	IdentityNumber     string    `json:"identity_number" example:"Dr. House"`
	Specialty          string    `json:"specialty" example:"Dr. House"`
	LicenseNumber      string    `json:"license_number" example:"Dr. House"`
	HospitalAffiliaton string    `json:"hospital_affiliation" example:"Dr. House"`
	CreatedAt          time.Time `json:"created_at" example:"Dr. House"`
	UpdatedAt          time.Time `json:"updated_at" example:"Dr. House"`
}
