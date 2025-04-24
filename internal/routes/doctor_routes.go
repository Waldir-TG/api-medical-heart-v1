package routes

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/controllers"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupDoctorRoutes(app *fiber.App, doctorServices *services.DoctorService) {
	doctorController := controllers.NewDoctorController(doctorServices)

	doctor := app.Group("/api/doctor")

	doctor.Get("/", doctorController.GetDoctors)
	doctor.Get("/:id", doctorController.GetDoctorDetail)
	doctor.Post("/", doctorController.CreateDoctor)
	doctor.Patch("/:id", doctorController.UpdateDoctor)
}
