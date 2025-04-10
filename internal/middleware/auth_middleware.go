// internal/middleware/auth_middleware.go
package middleware

import (
	"fmt"
	"strings"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/utils"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware verifica que el usuario esté autenticado
func AuthMiddleware(authService *services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener el token del header Authorization
		authHeader := c.Cookies("Authorization")
		fmt.Println("Token recibido:", authHeader)
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Authorization header is required",
			})
		}

		// Verificar formato del token
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid authorization format",
			})
		}
		token := parts[1]

		// Validar el token en la base de datos
		session, user, err := authService.ValidateToken(c.Context(), token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		// También validar el JWT
		claims, err := utils.ValidateJWT(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		// Verificar que el usuario del token coincida con el de la sesión
		if claims.UserID != user.ID {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		// Almacenar información del usuario y la sesión en el contexto
		c.Locals("user", user)
		c.Locals("session", session)
		c.Locals("claims", claims)

		return c.Next()
	}
}
