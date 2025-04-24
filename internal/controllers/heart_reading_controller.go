package controllers

import (
	"strconv"
	"time"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HeartReadingController struct {
	heartReadingService *services.HeartReadingService
}

func NewHeartReadingController(heartReadingService *services.HeartReadingService) *HeartReadingController {
	return &HeartReadingController{
		heartReadingService: heartReadingService,
	}
}

// CreateHeartReading handles the creation of a new heart reading
func (c *HeartReadingController) CreateHeartReading(ctx *fiber.Ctx) error {
	var request models.HeartReadingCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.heartReadingService.CreateHeartReading(ctx.Context(), &request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Heart reading created successfully",
	})
}

// GetPatientHeartReadings retrieves heart readings for a specific patient
func (c *HeartReadingController) GetPatientHeartReadings(ctx *fiber.Ctx) error {
	patientID, err := uuid.Parse(ctx.Params("patientId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	// Parse query parameters
	var startTime, endTime *time.Time
	if startTimeStr := ctx.Query("start_time"); startTimeStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid start_time format. Use RFC3339 format.",
			})
		}
		startTime = &parsedTime
	}

	if endTimeStr := ctx.Query("end_time"); endTimeStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid end_time format. Use RFC3339 format.",
			})
		}
		endTime = &parsedTime
	}

	var limit, offset *int
	if limitStr := ctx.Query("limit"); limitStr != "" {
		parsedLimit, err := strconv.Atoi(limitStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid limit parameter",
			})
		}
		limit = &parsedLimit
	}

	if offsetStr := ctx.Query("offset"); offsetStr != "" {
		parsedOffset, err := strconv.Atoi(offsetStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid offset parameter",
			})
		}
		offset = &parsedOffset
	}

	params := &models.HeartReadingQueryParams{
		PatientID: patientID,
		StartTime: startTime,
		EndTime:   endTime,
		Limit:     limit,
		Offset:    offset,
	}

	readings, err := c.heartReadingService.GetPatientHeartReadings(ctx.Context(), params)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(readings)
}

// GetHeartRateStats retrieves heart rate statistics for a patient
func (c *HeartReadingController) GetHeartRateStats(ctx *fiber.Ctx) error {
	patientID, err := uuid.Parse(ctx.Params("patientId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	// Parse query parameters
	var startTime, endTime *time.Time
	if startTimeStr := ctx.Query("start_time"); startTimeStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid start_time format. Use RFC3339 format.",
			})
		}
		startTime = &parsedTime
	}

	if endTimeStr := ctx.Query("end_time"); endTimeStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid end_time format. Use RFC3339 format.",
			})
		}
		endTime = &parsedTime
	}

	stats, err := c.heartReadingService.GetHeartRateStats(ctx.Context(), patientID, startTime, endTime)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(stats)
}

// GetHeartRateTrends retrieves heart rate trends for a patient
func (c *HeartReadingController) GetHeartRateTrends(ctx *fiber.Ctx) error {
	patientID, err := uuid.Parse(ctx.Params("patientId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	// Parse query parameters
	var startTime, endTime *time.Time
	if startTimeStr := ctx.Query("start_time"); startTimeStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid start_time format. Use RFC3339 format.",
			})
		}
		startTime = &parsedTime
	}

	if endTimeStr := ctx.Query("end_time"); endTimeStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid end_time format. Use RFC3339 format.",
			})
		}
		endTime = &parsedTime
	}

	resolution := ctx.Query("resolution", "day")

	trends, err := c.heartReadingService.GetHeartRateTrends(ctx.Context(), patientID, startTime, endTime, resolution)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(trends)
}

// GetHeartRateAnomalies retrieves heart rate anomalies for a patient
func (c *HeartReadingController) GetHeartRateAnomalies(ctx *fiber.Ctx) error {
	patientID, err := uuid.Parse(ctx.Params("patientId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	// Parse query parameters
	var startTime, endTime *time.Time
	if startTimeStr := ctx.Query("start_time"); startTimeStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, startTimeStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid start_time format. Use RFC3339 format.",
			})
		}
		startTime = &parsedTime
	}

	if endTimeStr := ctx.Query("end_time"); endTimeStr != "" {
		parsedTime, err := time.Parse(time.RFC3339, endTimeStr)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid end_time format. Use RFC3339 format.",
			})
		}
		endTime = &parsedTime
	}

	thresholdStr := ctx.Query("threshold", "0.2")
	threshold, err := strconv.ParseFloat(thresholdStr, 64)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid threshold parameter",
		})
	}

	anomalies, err := c.heartReadingService.GetHeartRateAnomalies(ctx.Context(), patientID, startTime, endTime, threshold)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(anomalies)
}

// GetHeartReadingByID retrieves a specific heart reading by ID
func (c *HeartReadingController) GetHeartReadingByID(ctx *fiber.Ctx) error {
	readingID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid reading ID",
		})
	}

	reading, err := c.heartReadingService.GetHeartReadingByID(ctx.Context(), readingID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if reading == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Heart reading not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(reading)
}

// UpdateHeartReading updates an existing heart reading
func (c *HeartReadingController) UpdateHeartReading(ctx *fiber.Ctx) error {
	readingID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid reading ID",
		})
	}

	patientID, err := uuid.Parse(ctx.Params("patientId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	var request models.HeartReadingUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.heartReadingService.UpdateHeartReading(ctx.Context(), readingID, patientID, &request); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Heart reading updated successfully",
	})
}

// ProcessUnprocessedReadings processes any unprocessed heart readings
func (c *HeartReadingController) ProcessUnprocessedReadings(ctx *fiber.Ctx) error {
	if err := c.heartReadingService.ProcessUnprocessedReadings(ctx.Context()); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Unprocessed readings processed successfully",
	})
}
