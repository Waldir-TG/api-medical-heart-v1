package routes

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/controllers"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/middleware"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupPatientRoutes(app *fiber.App, authService *services.AuthService, patientService *services.PatientService) {
	patientController := controllers.NewPatientController(patientService)
	// Group of routes for patients
	patients := app.Group("/api/patient", middleware.AuthMiddleware(authService))

	// Additional routes can be added as needed:
	patients.Get("/", patientController.GetAllPatientsBasicDetails)
	patients.Get("/details", patientController.GetAllPatientsDetails)
	patients.Get("/details/:id", patientController.GetPatientDetailsByID)
	patients.Get("/basicDetails/:id", patientController.GetPatientBasicDetailsByID)
	// Patient routes
	patients.Post("/", patientController.CreatePatient)
	patients.Patch("/:id", patientController.UpdatePatient)

	// patients.Delete("/:id", patientController.DeletePatient)

	// Routes protected by role
	// doctor := patients.Group("/doctor/patient", middleware.RoleMiddleware("doctor"))
	doctor := patients.Group("/doctor")
	doctor.Post("/", patientController.AssignDoctorToPatient)
	// doctor.Get("/", patientController.GetPatientsByDoctor) // This would need to be implemented
}
