// internal/middleware/role_middleware.go
package middleware

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/models"
	"github.com/gofiber/fiber/v2"
)

// RoleMiddleware verifica que el usuario tenga el rol adecuado
func RoleMiddleware(allowedRoles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Obtener el usuario del contexto (establecido por AuthMiddleware)
		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "User not found in context",
			})
		}

		// Verificar si el rol del usuario está permitido
		for _, role := range allowedRoles {
			if user.RoleName == role {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied: insufficient permissions",
		})
	}
}
