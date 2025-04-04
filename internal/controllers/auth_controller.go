// internal/controllers/auth_controller.go
package controllers

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

// Register maneja el registro de nuevos usuarios
func (c *AuthController) Register(ctx *fiber.Ctx) error {
	var req models.RegisterRequest

	// Parsear el cuerpo de la solicitud
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validar los datos de entrada (implementar validación según necesidad)

	// Registrar al usuario
	userID, err := c.authService.Register(ctx.Context(), &req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user_id": userID,
	})
}

// Login maneja el inicio de sesión de usuarios
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	var req models.LoginRequest

	// Parsear el cuerpo de la solicitud
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Validar los datos de entrada
	if req.Email == "" || req.Password == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Email and password are required",
		})
	}

	// Obtener información del dispositivo y dirección IP
	deviceInfo := ctx.Get("User-Agent")
	ipAddress := ctx.IP()

	// Iniciar sesión
	authResponse, err := c.authService.Login(ctx.Context(), &req, &deviceInfo, &ipAddress)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(authResponse)
}

// Logout maneja el cierre de sesión de usuarios
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	// Obtener el token del header Authorization
	authHeader := ctx.Get("Authorization")
	if authHeader == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	// Extraer el token
	token := authHeader[7:] // Eliminar "Bearer "

	// Invalidar sesión
	err := c.authService.Logout(ctx.Context(), token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

// ValidateSession verifica si una sesión es válida
func (c *AuthController) ValidateSession(ctx *fiber.Ctx) error {
	// El usuario ya está verificado por el middleware de autenticación
	user := ctx.Locals("user").(*models.User)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Token is valid",
		"user":    user,
	})
}
