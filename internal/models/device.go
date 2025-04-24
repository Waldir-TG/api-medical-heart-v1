package models

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID              uuid.UUID  `json:"id"`
	PatientID       *uuid.UUID `json:"patient_id,omitempty"` // Puntero para manejar NULL
	DeviceType      string     `json:"device_type"`
	SerialNumber    string     `json:"serial_number"`
	FirmwareVersion string     `json:"firmware_version"`
	LastSync        *time.Time `json:"last_sync,omitempty"`
	BatteryLevel    *int       `json:"battery_level,omitempty"` // Puntero para manejar NULL
	IsActive        bool       `json:"is_active"`
	RegisteredAt    time.Time  `json:"registered_at"`
}

type DeviceResponse struct {
	DeviceID        uuid.UUID  `json:"device_id"`
	PatientID       *uuid.UUID `json:"patient_id,omitempty"` // Puntero para manejar NULL
	DeviceType      string     `json:"device_type"`
	SerialNumber    string     `json:"serial_number"`
	FirmwareVersion string     `json:"firmware_version"`
	LastSync        *time.Time `json:"last_sync,omitempty"`
	BatteryLevel    *int       `json:"battery_level,omitempty"` // Puntero para manejar NULL
	IsActive        bool       `json:"is_active"`
	RegisteredAt    time.Time  `json:"registered_at"`
}

type DeviceRegisterRequest struct {
	PatientID       string `json:"patient_id" validate:"omitempty,uuid4"`
	DeviceType      string `json:"device_type" validate:"required"`
	SerialNumber    string `json:"serial_number" validate:"required"`
	FirmwareVersion string `json:"firmware_version"`
}

// DeviceUpdateRequest representa la solicitud para actualizar un dispositivo
type DeviceUpdateRequest struct {
	PatientID       *uuid.UUID `json:"patient_id" validate:"omitempty,uuid4"`
	DeviceType      *string    `json:"device_type" validate:"omitempty"`
	FirmwareVersion *string    `json:"firmware_version" validate:"omitempty"`
	BatteryLevel    *int       `json:"battery_level" validate:"omitempty,min=0,max=100"`
	IsActive        *bool      `json:"is_active" validate:"omitempty"`
}
