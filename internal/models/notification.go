package models

import (
	"time"

	"github.com/google/uuid"
)

// Notification representa una notificación enviada a un usuario
type Notification struct {
	ID          uuid.UUID  `json:"id"`
	RecipientID uuid.UUID  `json:"recipient_id"`
	AlertID     *uuid.UUID `json:"alert_id,omitempty"` // Puntero para manejar NULL
	Title       string     `json:"title"`
	Message     string     `json:"message"`
	SentAt      time.Time  `json:"sent_at"`
	ReadAt      *time.Time `json:"read_at,omitempty"` // Puntero para manejar NULL
	Type        string     `json:"notification_type"`
}

// NotificationUpdateRequest representa la solicitud para actualizar una notificación
type NotificationUpdateRequest struct {
	Read bool `json:"read"` // Marcar como leída/no leída
}

// FamilyMemberAddRequest representa la solicitud para agregar un familiar a un paciente
type FamilyMemberAddRequest struct {
	UserID              string `json:"user_id" validate:"required,uuid4"`
	PatientID           string `json:"patient_id" validate:"required,uuid4"`
	Relationship        string `json:"relationship" validate:"required"`
	NotificationEnabled bool   `json:"notification_enabled"`
}

// PatientMonitoringRequest representa la solicitud para actualizar parámetros de monitoreo
type PatientMonitoringRequest struct {
	MinHeartRate     *int  `json:"min_heart_rate" validate:"omitempty,min=20,max=200"`
	MaxHeartRate     *int  `json:"max_heart_rate" validate:"omitempty,min=20,max=200"`
	MonitoringActive *bool `json:"monitoring_active"`
}

type NotificationDetailResponse struct {
	Notification
	RelatedAlert *struct {
		AlertType string    `json:"type,omitempty"`
		Severity  string    `json:"severity,omitempty"`
		Time      time.Time `json:"time,omitempty"`
	} `json:"related_alert,omitempty"`
	RecipientName string `json:"recipient_name"`
}
