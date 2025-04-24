package routes

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/controllers"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/middleware"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupDeviceRoutes(app *fiber.App, authService *services.AuthService, deviceService *services.DeviceService) {
	deviceController := controllers.NewDeviceController(deviceService)

	// Group of routes for devices
	devices := app.Group("/api/devices", middleware.AuthMiddleware(authService))

	// Public device routes (still require authentication)
	devices.Get("/", deviceController.GetDevices)
	devices.Get("/:id", deviceController.GetDeviceByID)
	devices.Get("/patient/:patientId", deviceController.GetDevicesByPatientID)
	devices.Post("/", deviceController.RegisterDevice)
	devices.Patch("/:id", deviceController.UpdateDevice)
	devices.Post("/:id/sync", deviceController.UpdateDeviceSync)
	devices.Delete("/:id", deviceController.DeactivateDevice)

	// Admin-only routes
	// admin := devices.Group("/admin", middleware.RoleMiddleware("admin"))
	// Add admin-specific routes if needed
}
