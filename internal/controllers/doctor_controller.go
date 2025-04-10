package controllers

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

type DoctorController struct {
	doctorService *services.DoctorService
}

func NewDoctorController(doctorService *services.DoctorService) *DoctorController {
	return &DoctorController{
		doctorService: doctorService,
	}
}

// @Summary      Obtener lista de doctores
// @Description  Retorna todos los doctores registrados
// @Tags         doctors
// @Accept       json
// @Produce      json
// @Success      200  {array}   models.Doctor
// @Failure      500  {object}  map[string]string
// @Router       /api/doctor [get]
func (c *DoctorController) GetDoctors(ctx *fiber.Ctx) error {
	doctors, err := c.doctorService.GetDoctors(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting doctors",
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(doctors)
}
