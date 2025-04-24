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

// DoctorCreateRequest representa la solicitud para crear un médico
type DoctorCreateRequest struct {
	UserID              string `json:"user_id" validate:"required,uuid4"`
	IdentityType        string `json:"identity_type" validate:"required,oneof=CC TE PP PEP DIE"`
	IdentityNumber      string `json:"identity_number" validate:"required"`
	Specialty           string `json:"specialty" validate:"required"`
	LicenseNumber       string `json:"license_number" validate:"required"`
	HospitalAffiliation string `json:"hospital_affiliation"`
}

type DoctorUpdateRequest struct {
	IdentityType        *string `json:"identity_type" validate:"omitempty,oneof=CC TE PP PEP DIE"`
	IdentityNumber      *string `json:"identity_number" validate:"omitempty"`
	Specialty           *string `json:"specialty" validate:"omitempty"`
	LicenseNumber       *string `json:"license_number" validate:"omitempty"`
	HospitalAffiliation *string `json:"hospital_affiliation" validate:"omitempty"`
	IsActive            *bool   `json:"is_active" validate:"omitempty"`
}

type DoctorDetailsResponse struct {
	DoctorID           uuid.UUID `json:"doctor_id"`
	UserID             uuid.UUID `json:"user_id"`
	Email              string    `json:"email"`
	FirstName          string    `json:"first_name"`
	LastName           string    `json:"last_name"`
	PhoneNumber        string    `json:"phone_number"`
	IdentityType       string    `json:"identity_type"`
	IdentityNumber     string    `json:"identity_number"`
	Specialty          string    `json:"specialty"`
	LicenseNumber      string    `json:"license_number"`
	HospitalAffiliaton string    `json:"hospital_affiliation"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	IsActive           bool      `json:"is_active"`
}

type DoctorWithPatientsResponse struct {
	Doctor
	UserInfo struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"user_info"`
	Patients []PatientShortInfo `json:"patients,omitempty"`
}

// DoctorShortInfo para la respuesta de joins
type DoctorShortInfo struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	Specialty string    `json:"specialty"`
}
