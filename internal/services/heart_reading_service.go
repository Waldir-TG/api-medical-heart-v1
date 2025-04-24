package services

import (
	"context"
	"time"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/repositories"
	"github.com/google/uuid"
)

type HeartReadingService struct {
	heartReadingRepo *repositories.HeartReadingRepository
}

func NewHeartReadingService(heartReadingRepo *repositories.HeartReadingRepository) *HeartReadingService {
	return &HeartReadingService{
		heartReadingRepo: heartReadingRepo,
	}
}

// CreateHeartReading creates a new heart reading
func (s *HeartReadingService) CreateHeartReading(ctx context.Context, reading *models.HeartReadingCreateRequest) error {
	return s.heartReadingRepo.CreateHeartReading(ctx, reading)
}

// GetPatientHeartReadings retrieves heart readings for a specific patient
func (s *HeartReadingService) GetPatientHeartReadings(
	ctx context.Context,
	params *models.HeartReadingQueryParams,
) ([]*models.HeartReadingResponse, error) {
	return s.heartReadingRepo.GetPatientHeartReadings(ctx, params)
}

// GetHeartRateStats calculates statistics for heart readings
func (s *HeartReadingService) GetHeartRateStats(
	ctx context.Context,
	patientID uuid.UUID,
	startTime *time.Time,
	endTime *time.Time,
) (*models.HeartRateStats, error) {
	return s.heartReadingRepo.GetHeartRateStats(ctx, patientID, startTime, endTime)
}

// GetHeartRateTrends retrieves heart rate trends over time
func (s *HeartReadingService) GetHeartRateTrends(
	ctx context.Context,
	patientID uuid.UUID,
	startTime *time.Time,
	endTime *time.Time,
	resolution string,
) ([]*models.HeartRateTrend, error) {
	return s.heartReadingRepo.GetHeartRateTrends(ctx, patientID, startTime, endTime, resolution)
}

// GetHeartRateAnomalies detects anomalies in heart rate readings
func (s *HeartReadingService) GetHeartRateAnomalies(
	ctx context.Context,
	patientID uuid.UUID,
	startTime *time.Time,
	endTime *time.Time,
	threshold float64,
) ([]*models.HeartRateAnomaly, error) {
	return s.heartReadingRepo.GetHeartRateAnomalies(ctx, patientID, startTime, endTime, threshold)
}

// ProcessUnprocessedReadings processes any unprocessed heart readings
func (s *HeartReadingService) ProcessUnprocessedReadings(ctx context.Context) error {
	return s.heartReadingRepo.ProcessUnprocessedReadings(ctx)
}

// UpdateHeartReading updates an existing heart reading
func (s *HeartReadingService) UpdateHeartReading(
	ctx context.Context,
	readingID uuid.UUID,
	patientID uuid.UUID,
	update *models.HeartReadingUpdateRequest,
) error {
	return s.heartReadingRepo.UpdateHeartReading(ctx, readingID, patientID, update)
}

// GetHeartReadingByID retrieves a specific heart reading by ID
func (s *HeartReadingService) GetHeartReadingByID(
	ctx context.Context,
	readingID uuid.UUID,
) (*models.HeartReading, error) {
	return s.heartReadingRepo.GetHeartReadingByID(ctx, readingID)
}

// ScheduleReadingsProcessing schedules the processing of unprocessed readings
// This could be called by a background worker or scheduler
func (s *HeartReadingService) ScheduleReadingsProcessing(ctx context.Context) error {
	return s.ProcessUnprocessedReadings(ctx)
}

// GetPatientHeartReadingsWithTimeRange is a convenience method that creates a query params object
func (s *HeartReadingService) GetPatientHeartReadingsWithTimeRange(
	ctx context.Context,
	patientID uuid.UUID,
	startTime *time.Time,
	endTime *time.Time,
	limit *int,
	offset *int,
) ([]*models.HeartReadingResponse, error) {
	params := &models.HeartReadingQueryParams{
		PatientID: patientID,
		StartTime: startTime,
		EndTime:   endTime,
		Limit:     limit,
		Offset:    offset,
	}
	
	return s.GetPatientHeartReadings(ctx, params)
}