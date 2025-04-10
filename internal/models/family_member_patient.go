package models

import "time"

type FamilyMemberPatient struct {
	UserID              string    `json:"user_id"`
	PatientID           string    `json:"patient_id"`
	Relationship        string    `json:"relationship"`
	NotificationEnabled bool      `json:"notification_enabled"`
	AssignedDate        time.Time `json:"assigned_date"`
}
