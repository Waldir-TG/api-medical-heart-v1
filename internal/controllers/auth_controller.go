// internal/controllers/auth_controller.go
package controllers

import (
	"fmt"
	"time"

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
// @Summary      Registrar un nuevo usuario
// @Description  Crea una nueva cuenta de usuario en el sistema
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.RegisterRequest true "Datos de registro"
// @Success      201  {object}  map[string]interface{}  "Usuario registrado exitosamente"
// @Failure      400  {object}  map[string]string       "Cuerpo de solicitud inválido"
// @Failure      400  {object}  map[string]string       "Error en los datos de registro"
// @Router       /api/auth/register [post]
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

	//imprime el request
	fmt.Println(&req)

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user_id": userID,
	})
}

// Login maneja el inicio de sesión de usuarios
// @Summary      Iniciar sesión
// @Description  Autentica a un usuario y genera un token de acceso
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body models.LoginRequest true "Credenciales de acceso"
// @Success      200  {object}  models.AuthResponse  "Autenticación exitosa"
// @Failure      400  {object}  map[string]string    "Cuerpo de solicitud inválido o campos faltantes"
// @Failure      401  {object}  map[string]string    "Credenciales inválidas"
// @Router       /api/auth/login [post]
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

	// ExpiresAt := time.Now().Add(time.Hour)

	cookie := new(fiber.Cookie)
	cookie.Name = "Authorization"
	cookie.Value = "Bearer " + authResponse.Token
	cookie.HTTPOnly = true  // Importante para seguridad
	cookie.Secure = false   // Solo en HTTPS (en producción)
	cookie.SameSite = "Lax" // o "Strict" para más seguridad
	cookie.Expires = time.Now().Add(24 * time.Hour)

	ctx.Cookie(cookie)

	return ctx.Status(fiber.StatusOK).JSON(authResponse)
}

// Logout maneja el cierre de sesión de usuarios
// @Summary      Cerrar sesión
// @Description  Invalida el token de acceso del usuario
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  map[string]string  "Sesión cerrada exitosamente"
// @Failure      400  {object}  map[string]string  "Header de autorización faltante"
// @Failure      500  {object}  map[string]string  "Error al invalidar el token"
// @Router       /api/auth/logout [post]
func (c *AuthController) Logout(ctx *fiber.Ctx) error {
	// Obtener el token del header Authorization
	authCookie := ctx.Cookies("Authorization")
	if authCookie == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Authorization header is required",
		})
	}

	// Extraer el token
	token := authCookie[7:] // Eliminar "Bearer "

	// Invalidar sesión
	err := c.authService.Logout(ctx.Context(), token)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	ctx.ClearCookie("Authorization")

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

// ValidateSession verifica si una sesión es válida
// @Summary      Validar sesión
// @Description  Verifica si el token de acceso del usuario es válido
// @Tags         auth
// @Accept       json
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {object}  map[string]interface{}  "Token válido con información del usuario"
// @Router       /api/auth/validate [get]
func (c *AuthController) ValidateSession(ctx *fiber.Ctx) error {
	// El usuario ya está verificado por el middleware de autenticación
	user := ctx.Locals("user").(*models.User)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Token is valid",
		"user":    user,
	})
}
