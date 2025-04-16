package models

import (
	"time"

	"github.com/google/uuid"
)

type FamilyMemberPatient struct {
	UserID              uuid.UUID `json:"user_id"`
	PatientID           uuid.UUID `json:"patient_id"`
	Relationship        string    `json:"relationship"`
	NotificationEnabled bool      `json:"notification_enabled"`
	AssignedDate        time.Time `json:"assigned_date"`
}

// FamilyMemberInfo para la respuesta de joins
type FamilyMemberInfo struct {
	UserID       uuid.UUID `json:"user_id"`
	FullName     string    `json:"full_name"`
	Relationship string    `json:"relationship"`
}
