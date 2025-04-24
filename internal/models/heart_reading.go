package models

import (
	"time"

	"github.com/google/uuid"
)

// HeartReading represents a heart rate reading record
type HeartReading struct {
	ID                   uuid.UUID  `json:"id"`
	PatientID            uuid.UUID  `json:"patient_id"`
	DeviceID             *uuid.UUID `json:"device_id,omitempty"`
	EntryMethod          string     `json:"entry_method"` // 'device', 'manual', 'imported'
	EnteredBy            *uuid.UUID `json:"entered_by,omitempty"`
	ReadingType          string     `json:"reading_type"` // 'resting', 'active', 'sleep'
	Source               string     `json:"source"`       // 'smartwatch', 'chest_strap', 'manual_entry'
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
	Time                 time.Time  `json:"time"`
}

// HeartReadingCreateRequest represents the request to create a new heart reading
type HeartReadingCreateRequest struct {
	PatientID            uuid.UUID  `json:"patient_id" validate:"required"`
	DeviceID             *uuid.UUID `json:"device_id,omitempty"`
	EntryMethod          string     `json:"entry_method" validate:"required,oneof=device manual imported"`
	EnteredBy            *uuid.UUID `json:"entered_by,omitempty"`
	ReadingType          string     `json:"reading_type" validate:"required,oneof=resting active sleep"`
	Source               string     `json:"source" validate:"required"`
	BPM                  int        `json:"bpm" validate:"required,min=1,max=300"`
	Variability          *float64   `json:"variability,omitempty"`
	IrregularityDetected bool       `json:"irregularity_detected"`
	OxygenLevel          *int       `json:"oxygen_level,omitempty" validate:"omitempty,min=0,max=100"`
	SystolicPressure     *int       `json:"systolic_pressure,omitempty" validate:"omitempty,min=0,max=300"`
	DiastolicPressure    *int       `json:"diastolic_pressure,omitempty" validate:"omitempty,min=0,max=200"`
	Temperature          *float64   `json:"temperature,omitempty"`
	ActivityLevel        *int       `json:"activity_level,omitempty" validate:"omitempty,min=0,max=10"`
	Notes                *string    `json:"notes,omitempty"`
	ReliabilityScore     *float64   `json:"reliability_score,omitempty" validate:"omitempty,min=0,max=1"`
}

// HeartReadingUpdateRequest represents the request to update an existing heart reading
type HeartReadingUpdateRequest struct {
	ReadingType          *string  `json:"reading_type,omitempty" validate:"omitempty,oneof=resting active sleep"`
	BPM                  *int     `json:"bpm,omitempty" validate:"omitempty,min=1,max=300"`
	Variability          *float64 `json:"variability,omitempty"`
	IrregularityDetected *bool    `json:"irregularity_detected,omitempty"`
	OxygenLevel          *int     `json:"oxygen_level,omitempty" validate:"omitempty,min=0,max=100"`
	Notes                *string  `json:"notes,omitempty"`
	ReliabilityScore     *float64 `json:"reliability_score,omitempty" validate:"omitempty,min=0,max=1"`
}

// HeartReadingResponse represents a heart reading with additional information
type HeartReadingResponse struct {
	ReadingTime          time.Time `json:"reading_time"`
	DeviceType           *string   `json:"device_type,omitempty"`
	EntryMethod          string    `json:"entry_method"`
	ReadingType          string    `json:"reading_type"`
	Source               string    `json:"source"`
	BPM                  int       `json:"bpm"`
	Variability          *float64  `json:"variability,omitempty"`
	IrregularityDetected bool      `json:"irregularity_detected"`
	OxygenLevel          *int      `json:"oxygen_level,omitempty"`
	Notes                *string   `json:"notes,omitempty"`
	ReliabilityScore     *float64  `json:"reliability_score,omitempty"`
}

// HeartRateStats represents statistics about heart rate readings
type HeartRateStats struct {
	AvgBPM            float64 `json:"avg_bpm"`
	MinBPM            int     `json:"min_bpm"`
	MaxBPM            int     `json:"max_bpm"`
	VariabilityAvg    float64 `json:"variability_avg"`
	ReadingCount      int64   `json:"reading_count"`
	IrregularityCount int64   `json:"irregularity_count"`
	LowReadingsCount  int64   `json:"low_readings_count"`
	HighReadingsCount int64   `json:"high_readings_count"`
}

// HeartRateAnomaly represents an anomaly in heart rate readings
type HeartRateAnomaly struct {
	ReadingTime         time.Time `json:"reading_time"`
	BPM                 int       `json:"bpm"`
	ExpectedRangeLow    int       `json:"expected_range_low"`
	ExpectedRangeHigh   int       `json:"expected_range_high"`
	DeviationPercentage float64   `json:"deviation_percentage"`
	IsIrregular         bool      `json:"is_irregular"`
	AnomalyType         string    `json:"anomaly_type"` // 'extreme_low', 'extreme_high', 'low', 'high', 'irregular', 'normal'
}

// HeartRateTrend represents a trend in heart rate readings over time
type HeartRateTrend struct {
	PeriodStart  time.Time `json:"period_start"`
	PeriodEnd    time.Time `json:"period_end"`
	AvgBPM       float64   `json:"avg_bpm"`
	MinBPM       int       `json:"min_bpm"`
	MaxBPM       int       `json:"max_bpm"`
	ReadingCount int64     `json:"reading_count"`
}

// HeartReadingQueryParams represents query parameters for fetching heart readings
type HeartReadingQueryParams struct {
	PatientID  uuid.UUID  `json:"patient_id" validate:"required"`
	StartTime  *time.Time `json:"start_time,omitempty"`
	EndTime    *time.Time `json:"end_time,omitempty"`
	Limit      *int       `json:"limit,omitempty" validate:"omitempty,min=1,max=10000"`
	Offset     *int       `json:"offset,omitempty" validate:"omitempty,min=0"`
	HoursBack  *int       `json:"hours_back,omitempty" validate:"omitempty,min=1,max=8760"` // Max 1 year in hours
	Resolution *string    `json:"resolution,omitempty" validate:"omitempty,oneof=minute hour day week month"`
}

// HeartReadingWithDeviceResponse extends HeartReadingResponse with device information
type HeartReadingWithDeviceResponse struct {
	HeartReadingResponse
	DeviceInfo *struct {
		ID           uuid.UUID `json:"id,omitempty"`
		DeviceType   string    `json:"device_type,omitempty"`
		SerialNumber string    `json:"serial_number,omitempty"`
	} `json:"device_info,omitempty"`
	EnteredByInfo *struct {
		ID        uuid.UUID `json:"id,omitempty"`
		FirstName string    `json:"first_name,omitempty"`
		LastName  string    `json:"last_name,omitempty"`
	} `json:"entered_by_info,omitempty"`
}
