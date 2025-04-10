package models

import "time"

type Alert struct {
	ID             string    `json:"id"`
	PatientID      string    `json:"patient_id"`
	AlertType      string    `json:"alert_type"`
	Severity       string    `json:"severity"`
	Message        string    `json:"message"`
	ReadingTime    time.Time `json:"reading_time"`
	Acknowledged   bool      `json:"acknowledged"`
	AcknowledgedBy string    `json:"acknowledged_by"`
	AcknowledgedAt time.Time `json:"acknowledged_at"`
	CreatedAt      time.Time `json:"created_at"`
}
