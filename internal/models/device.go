package models

import "time"

type Device struct {
	ID              string    `json:"id"`
	PatientID       string    `json:"patient_id"`
	DeviceType      string    `json:"device_type"`
	SerialNumber    string    `json:"serial_number"`
	FirmwareVersion string    `json:"firmware_version"`
	LastSync        string    `json:"last_sync"`
	BatteryLevel    int       `json:"battery_level"`
	IsActive        bool      `json:"is_active"`
	RegisteredAt    time.Time `json:"registered_at"`
}
