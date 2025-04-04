// internal/controllers/user_controller.go
package controllers

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// GetProfile obtiene el perfil del usuario autenticado
func (c *UserController) GetProfile(ctx *fiber.Ctx) error {
	user, ok := ctx.Locals("user").(*models.User)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found in context",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

// GetAllUsers obtiene todos los usuarios (solo para administradores)
func (c *UserController) GetAllUsers(ctx *fiber.Ctx) error {
	// Implementar lógica para obtener todos los usuarios
	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Method not implemented",
	})
}

// CreateUser crea un nuevo usuario (solo para administradores)
func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	// Implementar lógica para crear un usuario
	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Method not implemented",
	})
}

// UpdateUser actualiza un usuario existente (solo para administradores)
func (c *UserController) UpdateUser(ctx *fiber.Ctx) error {
	// Implementar lógica para actualizar un usuario
	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Method not implemented",
	})
}

// DeleteUser elimina un usuario (solo para administradores)
func (c *UserController) DeleteUser(ctx *fiber.Ctx) error {
	// Implementar lógica para eliminar un usuario
	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Method not implemented",
	})
}

// GetDoctorPatients obtiene todos los pacientes asignados a un médico
func (c *UserController) GetDoctorPatients(ctx *fiber.Ctx) error {
	// Implementar lógica para obtener los pacientes de un médico
	return ctx.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"message": "Method not implemented",
	})
}
