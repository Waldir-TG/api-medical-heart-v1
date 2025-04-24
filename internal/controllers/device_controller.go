package controllers

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeviceController struct {
	deviceService *services.DeviceService
}

func NewDeviceController(deviceService *services.DeviceService) *DeviceController {
	return &DeviceController{
		deviceService: deviceService,
	}
}

func (c *DeviceController) GetDevices(ctx *fiber.Ctx) error {
	devices, err := c.deviceService.GetDevices(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(devices)
}

func (c *DeviceController) GetDeviceByID(ctx *fiber.Ctx) error {
	id, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid device ID",
		})
	}

	device, err := c.deviceService.GetDeviceByID(ctx.Context(), id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if device == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Device not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(device)
}

func (c *DeviceController) GetDevicesByPatientID(ctx *fiber.Ctx) error {
	patientID, err := uuid.Parse(ctx.Params("patientId"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	devices, err := c.deviceService.GetDevicesByPatientID(ctx.Context(), patientID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(devices)
}

func (c *DeviceController) RegisterDevice(ctx *fiber.Ctx) error {
	var request models.DeviceRegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	deviceID, err := c.deviceService.RegisterDevice(ctx.Context(), &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":   "Device registered successfully",
		"device_id": deviceID,
	})
}

func (c *DeviceController) UpdateDevice(ctx *fiber.Ctx) error {
	deviceID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid device ID",
		})
	}

	var request models.DeviceUpdateRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	message, err := c.deviceService.UpdateDevice(ctx.Context(), deviceID, &request)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}

func (c *DeviceController) UpdateDeviceSync(ctx *fiber.Ctx) error {
	deviceID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid device ID",
		})
	}

	type SyncRequest struct {
		BatteryLevel int `json:"battery_level"`
	}

	var request SyncRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	err = c.deviceService.UpdateDeviceSync(ctx.Context(), deviceID, request.BatteryLevel)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Device sync updated successfully",
	})
}

func (c *DeviceController) DeactivateDevice(ctx *fiber.Ctx) error {
	deviceID, err := uuid.Parse(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid device ID",
		})
	}

	message, err := c.deviceService.DeactivateDevice(ctx.Context(), deviceID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": message,
	})
}