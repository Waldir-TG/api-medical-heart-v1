package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/db"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/google/uuid"
)

type HeartReadingRepository struct {
	db *db.PostgresDB
}

func NewHeartReadingRepository(database *db.PostgresDB) *HeartReadingRepository {
	return &HeartReadingRepository{db: database}
}

// CreateHeartReading inserts a new heart reading record
func (r *HeartReadingRepository) CreateHeartReading(ctx context.Context, reading *models.HeartReadingCreateRequest) error {
	query := `CALL insert_heart_reading($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);`

	_, err := r.db.Pool.Exec(ctx, query,
		reading.PatientID,
		reading.DeviceID,
		reading.EntryMethod,
		reading.EnteredBy,
		reading.ReadingType,
		reading.Source,
		reading.BPM,
		reading.Variability,
		reading.IrregularityDetected,
		reading.OxygenLevel,
		reading.SystolicPressure,
		reading.DiastolicPressure,
		reading.Temperature,
		reading.ActivityLevel,
		reading.Notes,
		reading.ReliabilityScore,
	)

	if err != nil {
		return fmt.Errorf("failed to create heart reading: %w", err)
	}

	return nil
}

// GetPatientHeartReadings retrieves heart readings for a specific patient
func (r *HeartReadingRepository) GetPatientHeartReadings(
	ctx context.Context,
	params *models.HeartReadingQueryParams,
) ([]*models.HeartReadingResponse, error) {
	query := `SELECT * FROM get_patient_heart_readings($1, $2, $3, $4, $5);`

	limit := 1000
	offset := 0

	if params.Limit != nil {
		limit = *params.Limit
	}

	if params.Offset != nil {
		offset = *params.Offset
	}

	rows, err := r.db.Pool.Query(ctx, query,
		params.PatientID,
		params.StartTime,
		params.EndTime,
		limit,
		offset,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get patient heart readings: %w", err)
	}
	defer rows.Close()

	var readings []*models.HeartReadingResponse

	for rows.Next() {
		var reading models.HeartReadingResponse
		var deviceType *string
		var variability *float64
		var oxygenLevel *int
		var notes *string
		var reliabilityScore *float64

		if err := rows.Scan(
			&reading.ReadingTime,
			&deviceType,
			&reading.EntryMethod,
			&reading.ReadingType,
			&reading.Source,
			&reading.BPM,
			&variability,
			&reading.IrregularityDetected,
			&oxygenLevel,
			&notes,
			&reliabilityScore,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		reading.DeviceType = deviceType
		reading.Variability = variability
		reading.OxygenLevel = oxygenLevel
		reading.Notes = notes
		reading.ReliabilityScore = reliabilityScore

		readings = append(readings, &reading)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return readings, nil
}

// GetHeartRateStats calculates statistics for heart readings
func (r *HeartReadingRepository) GetHeartRateStats(
	ctx context.Context,
	patientID uuid.UUID,
	startTime *time.Time,
	endTime *time.Time,
) (*models.HeartRateStats, error) {
	query := `SELECT * FROM calculate_heart_rate_stats($1, $2, $3);`

	row := r.db.Pool.QueryRow(ctx, query, patientID, startTime, endTime)

	var stats models.HeartRateStats

	if err := row.Scan(
		&stats.AvgBPM,
		&stats.MinBPM,
		&stats.MaxBPM,
		&stats.VariabilityAvg,
		&stats.ReadingCount,
		&stats.IrregularityCount,
		&stats.LowReadingsCount,
		&stats.HighReadingsCount,
	); err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	return &stats, nil
}

// GetHeartRateTrends retrieves heart rate trends over time
func (r *HeartReadingRepository) GetHeartRateTrends(
	ctx context.Context,
	patientID uuid.UUID,
	startTime *time.Time,
	endTime *time.Time,
	resolution string,
) ([]*models.HeartRateTrend, error) {
	query := `SELECT * FROM get_heart_rate_trend($1, $2, $3, $4);`

	rows, err := r.db.Pool.Query(ctx, query, patientID, startTime, endTime, resolution)
	if err != nil {
		return nil, fmt.Errorf("failed to get heart rate trends: %w", err)
	}
	defer rows.Close()

	var trends []*models.HeartRateTrend

	for rows.Next() {
		var trend models.HeartRateTrend

		if err := rows.Scan(
			&trend.PeriodStart,
			&trend.PeriodEnd,
			&trend.AvgBPM,
			&trend.MinBPM,
			&trend.MaxBPM,
			&trend.ReadingCount,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		trends = append(trends, &trend)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return trends, nil
}

// GetHeartRateAnomalies detects anomalies in heart rate readings
func (r *HeartReadingRepository) GetHeartRateAnomalies(
	ctx context.Context,
	patientID uuid.UUID,
	startTime *time.Time,
	endTime *time.Time,
	threshold float64,
) ([]*models.HeartRateAnomaly, error) {
	query := `SELECT * FROM detect_heart_rate_anomalies($1, $2, $3, $4);`

	rows, err := r.db.Pool.Query(ctx, query, patientID, startTime, endTime, threshold)
	if err != nil {
		return nil, fmt.Errorf("failed to get heart rate anomalies: %w", err)
	}
	defer rows.Close()

	var anomalies []*models.HeartRateAnomaly

	for rows.Next() {
		var anomaly models.HeartRateAnomaly

		if err := rows.Scan(
			&anomaly.ReadingTime,
			&anomaly.BPM,
			&anomaly.ExpectedRangeLow,
			&anomaly.ExpectedRangeHigh,
			&anomaly.DeviationPercentage,
			&anomaly.IsIrregular,
			&anomaly.AnomalyType,
		); err != nil {
			return nil, fmt.Errorf("scan failed: %w", err)
		}

		anomalies = append(anomalies, &anomaly)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration failed: %w", err)
	}

	return anomalies, nil
}

// ProcessUnprocessedReadings processes any unprocessed heart readings
func (r *HeartReadingRepository) ProcessUnprocessedReadings(ctx context.Context) error {
	query := `CALL process_unprocessed_readings();`

	_, err := r.db.Pool.Exec(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to process unprocessed readings: %w", err)
	}

	return nil
}

// UpdateHeartReading updates an existing heart reading
func (r *HeartReadingRepository) UpdateHeartReading(
	ctx context.Context,
	readingID uuid.UUID,
	patientID uuid.UUID,
	update *models.HeartReadingUpdateRequest,
) error {
	query := `CALL update_heart_reading($1, $2, $3, $4, $5, $6, $7, $8);`

	_, err := r.db.Pool.Exec(ctx, query,
		patientID,
		readingID,
		update.ReadingType,
		update.BPM,
		update.Variability,
		update.IrregularityDetected,
		update.OxygenLevel,
		update.Notes,
	)

	if err != nil {
		return fmt.Errorf("failed to update heart reading: %w", err)
	}

	return nil
}

// GetHeartReadingByID retrieves a specific heart reading by ID
func (r *HeartReadingRepository) GetHeartReadingByID(
	ctx context.Context,
	readingID uuid.UUID,
) (*models.HeartReading, error) {
	query := `SELECT * FROM get_heart_reading_by_id($1);`

	row := r.db.Pool.QueryRow(ctx, query, readingID)

	var reading models.HeartReading
	var deviceID *uuid.UUID
	var enteredBy *uuid.UUID
	var variability *float64
	var oxygenLevel *int
	var systolicPressure *int
	var diastolicPressure *int
	var temperature *float64
	var activityLevel *int
	var notes *string
	var reliabilityScore *float64

	if err := row.Scan(
		&reading.ID,
		&reading.PatientID,
		&deviceID,
		&reading.EntryMethod,
		&enteredBy,
		&reading.ReadingType,
		&reading.Source,
		&reading.BPM,
		&variability,
		&reading.IrregularityDetected,
		&oxygenLevel,
		&systolicPressure,
		&diastolicPressure,
		&temperature,
		&activityLevel,
		&notes,
		&reliabilityScore,
		&reading.Processed,
		&reading.Time,
	); err != nil {
		return nil, fmt.Errorf("scan failed: %w", err)
	}

	reading.DeviceID = deviceID
	reading.EnteredBy = enteredBy
	reading.Variability = variability
	reading.OxygenLevel = oxygenLevel
	reading.SystolicPressure = systolicPressure
	reading.DiastolicPressure = diastolicPressure
	reading.Temperature = temperature
	reading.ActivityLevel = activityLevel
	reading.Notes = notes
	reading.ReliabilityScore = reliabilityScore

	return &reading, nil
}
