package models

import (
	"time"

	"github.com/google/uuid"
)

// HeartReading representa una lectura cardíaca (equivalente a heart_readings)
type HeartReading struct {
	Time                 time.Time  `json:"time"`
	PatientID            uuid.UUID  `json:"patient_id"`
	DeviceID             *uuid.UUID `json:"device_id,omitempty"`
	EntryMethod          string     `json:"entry_method"` // automatic, manual
	EnteredBy            *uuid.UUID `json:"entered_by,omitempty"`
	ReadingType          string     `json:"reading_type"` // normal, exercise, rest, medication
	Source               string     `json:"source"`       // watch, ecg_monitor, phone_app, doctor, family, self
	BPM                  int        `json:"bpm"`
	Variability          *float64   `json:"variability,omitempty"`
	IrregularityDetected bool       `json:"irregularity_detected"`
	OxygenLevel          *int       `json:"oxygen_level,omitempty"`
	SystolicPressure     *int       `json:"systolic_pressure,omitempty"`
	DiastolicPressure    *int       `json:"diastolic_pressure,omitempty"`
	Temperature          *float64   `json:"temperature,omitempty"`
	ActivityLevel        *int       `json:"activity_level,omitempty"`
	Notes                *string    `json:"notes,omitempty"`
	ReliabilityScore     *float64   `json:"reliability_score,omitempty"`
	Processed            bool       `json:"processed"`
}

// HeartReadingRequest representa la solicitud para registrar una lectura cardíaca
type HeartReadingRequest struct {
	PatientID            string   `json:"patient_id" validate:"required,uuid4"`
	DeviceID             string   `json:"device_id" validate:"omitempty,uuid4"`
	EntryMethod          string   `json:"entry_method" validate:"required,oneof=automatic manual"`
	ReadingType          string   `json:"reading_type" validate:"required,oneof=normal exercise rest medication"`
	Source               string   `json:"source" validate:"required,oneof=watch ecg_monitor phone_app doctor family self"`
	BPM                  int      `json:"bpm" validate:"required,min=20,max=250"`
	Variability          *float64 `json:"variability" validate:"omitempty,min=0"`
	IrregularityDetected bool     `json:"irregularity_detected"`
	OxygenLevel          *int     `json:"oxygen_level" validate:"omitempty,min=50,max=100"`
	Notes                *string  `json:"notes" validate:"omitempty,max=500"`
}

// AlertAcknowledgeRequest representa la solicitud para reconocer una alerta
type AlertAcknowledgeRequest struct {
	AlertID string `json:"alert_id" validate:"required,uuid4"`
}

type HeartReadingWithDeviceResponse struct {
	HeartReading
	DeviceInfo *struct {
		Type         string `json:"type,omitempty"`
		SerialNumber string `json:"serial_number,omitempty"`
	} `json:"device_info,omitempty"`
	EnteredByName *string `json:"entered_by,omitempty"`
}
