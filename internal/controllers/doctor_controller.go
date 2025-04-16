package controllers

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (c *DoctorController) GetDoctorDetail(ctx *fiber.Ctx) error {
	doctorID := ctx.Params("id")
	if doctorID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Doctor ID is required",
		})
	}
	// Convertir doctorID a UUID
	doctorUUID, err := uuid.Parse(doctorID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid doctor ID",
		})
	}

	doctor, err := c.doctorService.GetDoctorDetail(ctx.Context(), doctorUUID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message":   "Error getting doctor detail",
			"doctor_id": doctorID,
			"error":     err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(doctor)
}

func (c *DoctorController) CreateDoctor(ctx *fiber.Ctx) error {
	var doctorData models.DoctorCreateRequest
	if err := ctx.BodyParser(&doctorData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	doctorID, err := c.doctorService.CreateDoctor(ctx.Context(), &doctorData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error creating doctor",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":  "Doctor created successfully",
		"doctorID": doctorID,
	})
}

func (c *DoctorController) UpdateDoctor(ctx *fiber.Ctx) error {
	doctorID := ctx.Params("id")
	if doctorID == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Doctor ID is required",
		})
	}
	// Convertir doctorID a UUID
	doctorUUID, err := uuid.Parse(doctorID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid doctor ID",
		})
	}

	var doctorData models.DoctorUpdateRequest
	if err := ctx.BodyParser(&doctorData); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	message, err := c.doctorService.UpdateDoctor(ctx.Context(), doctorUUID, &doctorData)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error updating doctor",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}
