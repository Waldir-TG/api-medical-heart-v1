package models

import "time"

type Notification struct {
	ID          string    `json:"id"`
	RecipientID string    `json:"recipient_id"`
	AlertID     string    `json:"alert_id"`
	Title       string    `json:"title"`
	Message     string    `json:"message"`
	SentAt      time.Time `json:"sent_at"`
	ReadAt      time.Time `json:"read_at"`
	Type        string    `json:"notification_type"`
}
