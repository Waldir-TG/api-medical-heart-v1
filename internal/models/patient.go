package models

import (
	"time"

	"github.com/google/uuid"
)

// Patient representa la información médica de un paciente
type Patient struct {
	PatientID                uuid.UUID `json:"patient_id"`
	UserID                   uuid.UUID `json:"user_id"`
	DateOfBirth              time.Time `json:"date_of_birth"`
	Gender                   string    `json:"gender"` // Considerar usar un tipo enum personalizado
	BloodType                string    `json:"blood_type"`
	HeightCm                 int       `json:"height_cm"`
	WeightKg                 float64   `json:"weight_kg"`
	MedicalConditions        []string  `json:"medical_conditions"`
	Allergies                []string  `json:"allergies"`
	EmergencyContactName     string    `json:"emergency_contact_name"`
	EmergencyContactPhone    string    `json:"emergency_contact_phone"`
	EmergencyContactRelation string    `json:"emergency_contact_relation"`
	MinHeartRate             int       `json:"min_heart_rate"`
	MaxHeartRate             int       `json:"max_heart_rate"`
	MonitoringActive         bool      `json:"monitoring_active"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

// Doctor representa la información profesional de un médico

// DoctorPatient representa la relación entre médicos y pacientes
type DoctorPatient struct {
	DoctorID     uuid.UUID `json:"doctor_id"`
	PatientID    uuid.UUID `json:"patient_id"`
	AssignedDate time.Time `json:"assigned_date"`
}

// PatientCreateRequest representa la solicitud para crear un paciente
type PatientCreateRequest struct {
	UserID                   string   `json:"user_id" validate:"required,uuid4"`
	DateOfBirth              string   `json:"date_of_birth" validate:"required,datetime=2006-01-02"`
	Gender                   string   `json:"gender" validate:"required,oneof=male female other"`
	BloodType                string   `json:"blood_type" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	HeightCm                 int      `json:"height_cm" validate:"required,min=30,max=300"`
	WeightKg                 float64  `json:"weight_kg" validate:"required,min=0.5,max=500"`
	MedicalConditions        []string `json:"medical_conditions"`
	Allergies                []string `json:"allergies"`
	EmergencyContactName     string   `json:"emergency_contact_name" validate:"required"`
	EmergencyContactPhone    string   `json:"emergency_contact_phone" validate:"required"`
	EmergencyContactRelation string   `json:"emergency_contact_relation" validate:"required"`
	MinHeartRate             int      `json:"min_heart_rate" validate:"required,min=20,max=200"`
	MaxHeartRate             int      `json:"max_heart_rate" validate:"required,min=20,max=200,gtfield=MinHeartRate"`
}

type PatientUpdatedRequest struct {
	// DateOfBirth              *string   `json:"date_of_birth" validate:"datetime=2006-01-02"`
	// Gender                   *string   `json:"gender" validate:"oneof=male female other"`
	// BloodType                *string   `json:"blood_type" validate:"omitempty,oneof=A+ A- B+ B- AB+ AB- O+ O-"`
	HeightCm                 *int      `json:"height_cm" validate:"min=30,max=300"`
	WeightKg                 *float64  `json:"weight_kg" validate:"min=0.5,max=500"`
	MedicalConditions        *[]string `json:"medical_conditions"`
	Allergies                *[]string `json:"allergies"`
	EmergencyContactName     *string   `json:"emergency_contact_name"`
	EmergencyContactPhone    *string   `json:"emergency_contact_phone"`
	EmergencyContactRelation *string   `json:"emergency_contact_relation"`
	MinHeartRate             *int      `json:"min_heart_rate" validate:"min=20,max=200"`
	MaxHeartRate             *int      `json:"max_heart_rate" validate:"min=20,max=200,gtfield=MinHeartRate"`
	MonitoringActive         *bool     `json:"monitoring_active"`
	AlertRecipients          *bool     `json:"alert_recipients"`
}

// DoctorPatientAssignRequest representa la solicitud para asignar un médico a un paciente
type DoctorPatientAssignRequest struct {
	DoctorID  string `json:"doctor_id" validate:"required,uuid4"`
	PatientID string `json:"patient_id" validate:"required,uuid4"`
}

type PatientDetailBasicResponse struct {
	Patient
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	IsActive    bool   `json:"is_active"`
}

type PatientDetailResponse struct {
	PatientDetailBasicResponse
	AssignedDoctors *[]struct {
		DoctorID      uuid.UUID `json:"doctor_id"`
		UserID        uuid.UUID `json:"user_id"`
		FullName      string    `json:"full_name"`
		Specialty     string    `json:"specialty"`
		LicenseNumber string    `json:"license_number"`
		Hospital      string    `json:"hospital_affiliation"`
		AssignedDate  time.Time `json:"assigned_date"`
	}
	FamilyMembers *[]struct {
		UserID       uuid.UUID `json:"user_id"`
		FullName     string    `json:"full_name"`
		Email        string    `json:"email"`
		PhoneNumber  string    `json:"phone_number"`
		Relationship string    `json:"relationship"`
		Notification bool      `json:"notification_enabled"`
		AssignedDate time.Time `json:"assigned_date"`
	}
	LastHeartReadings *[]struct {
		Time             time.Time `json:"time"`
		Bpm              int       `json:"bpm"`
		Variability      string    `json:"variability"`
		Source           string    `json:"source"`
		ReadingType      string    `json:"reading_type"`
		ReliabilityScore float64   `json:"reliability_score"`
	}
}

type PatientShortInfo struct {
	ID        uuid.UUID `json:"id"`
	FullName  string    `json:"full_name"`
	BirthDate time.Time `json:"birth_date"`
}
