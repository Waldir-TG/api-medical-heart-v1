package models

import (
	"time"

	"github.com/google/uuid"
)

type Alert struct {
	ID             uuid.UUID  `json:"id"`
	PatientID      uuid.UUID  `json:"patient_id"`
	AlertType      string     `json:"alert_type"`
	Severity       string     `json:"severity"` // low, medium, high, critical
	Message        string     `json:"message"`
	ReadingTime    time.Time  `json:"reading_time"`
	Acknowledged   bool       `json:"acknowledged"`
	AcknowledgedBy *uuid.UUID `json:"acknowledged_by,omitempty"`
	AcknowledgedAt *time.Time `json:"acknowledged_at,omitempty"`
	CreatedAt      time.Time  `json:"created_at"`
}

type AlertWithPatientResponse struct {
	Alert
	PatientInfo struct {
		FullName  string    `json:"full_name"`
		BirthDate time.Time `json:"birth_date"`
	} `json:"patient_info"`
	AcknowledgedByUser *string `json:"acknowledged_by,omitempty"`
}
