package models

import (
	"time"

	"github.com/google/uuid"
)

type Patient struct {
	ID                       uuid.UUID `json:"id"`
	UserID                   uuid.UUID `json:"user_id"`
	DateOfBirth              string    `json:"date_of_birth"`
	Gender                   string    `json:"gender"`
	BloodType                string    `json:"blood_type"`
	HeightCM                 float32   `json:"height_cm"`
	WeightKG                 float32   `json:"weight_kg"`
	MedicalConditions        string    `json:"medical_conditions"`
	Allergies                string    `json:"allergies"`
	EmergencyContactName     string    `json:"emergency_contact_name"`
	EmergencyContactPhone    string    `json:"emergency_contact_phone"`
	EmergencyContactRelation string    `json:"emergency_contact_relation"`
	MinHeartRate             int       `json:"min_heart_rate"`
	MaxHeartRate             int       `json:"max_heart_rate"`
	MonitoringActive         bool      `json:"monitoring_active"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}
