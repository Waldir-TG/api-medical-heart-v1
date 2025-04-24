package routes

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/controllers"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/middleware"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupHeartReadingRoutes(app *fiber.App, authService *services.AuthService, heartReadingService *services.HeartReadingService) {
	heartReadingController := controllers.NewHeartReadingController(heartReadingService)

	// Group of routes for heart readings
	heartReadings := app.Group("/api/heart-readings", middleware.AuthMiddleware(authService))

	// Routes for heart readings
	heartReadings.Post("/", heartReadingController.CreateHeartReading)
	heartReadings.Get("/patient/:patientId", heartReadingController.GetPatientHeartReadings)
	heartReadings.Get("/patient/:patientId/stats", heartReadingController.GetHeartRateStats)
	heartReadings.Get("/patient/:patientId/trends", heartReadingController.GetHeartRateTrends)
	heartReadings.Get("/patient/:patientId/anomalies", heartReadingController.GetHeartRateAnomalies)
	heartReadings.Get("/:id", heartReadingController.GetHeartReadingByID)
	heartReadings.Put("/patient/:patientId/:id", heartReadingController.UpdateHeartReading)

	// Admin-only routes
	admin := heartReadings.Group("/admin", middleware.RoleMiddleware("admin"))
	admin.Post("/process", heartReadingController.ProcessUnprocessedReadings)
}