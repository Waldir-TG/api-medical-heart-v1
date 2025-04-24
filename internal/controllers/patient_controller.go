package controllers

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type PatientController struct {
	patientService *services.PatientService
}

func NewPatientController(patientService *services.PatientService) *PatientController {
	return &PatientController{
		patientService: patientService,
	}
}

func (c *PatientController) GetAllPatientsBasicDetails(ctx *fiber.Ctx) error {
	patients, err := c.patientService.GetAllPatientsBasicDetails(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(patients)
}

func (c *PatientController) GetPatientBasicDetailsByID(ctx *fiber.Ctx) error {
	patientID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	patient, err := c.patientService.GetPatientBasicDetailsByID(ctx.Context(), &patientID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(patient)
}

func (c *PatientController) GetAllPatientsDetails(ctx *fiber.Ctx) error {
	patients, err := c.patientService.GetPatientsDetails(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(patients)
}

func (c *PatientController) GetPatientDetailsByID(ctx *fiber.Ctx) error {
	patientID, err := uuid.Parse(ctx.Params("id"))

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}
	patient, err := c.patientService.GetPatientDetailsByID(ctx.Context(), &patientID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(patient)
}

func (c *PatientController) CreatePatient(ctx *fiber.Ctx) error {
	var request models.PatientCreateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	patientID, err := c.patientService.CreatePatient(ctx.Context(), &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":    "Patient created successfully",
		"patient_id": patientID,
	})
}

func (c *PatientController) UpdatePatient(ctx *fiber.Ctx) error {
	patientID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	var request models.PatientUpdatedRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	message, err := c.patientService.UpdatePatient(ctx.Context(), patientID, &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}

func (c *PatientController) AssignDoctorToPatient(ctx *fiber.Ctx) error {
	var request models.DoctorPatientAssignRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	message, err := c.patientService.AssignDoctorToPatient(ctx.Context(), &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}

// Additional handler methods can be added as needed
