// internal/routes/auth_routes.go
package routes

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/controllers"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/middleware"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App, authService *services.AuthService) {
	// Crear controlador
	authController := controllers.NewAuthController(authService)

	// Grupo de rutas para autenticación
	auth := app.Group("/api/auth")

	// Rutas públicas
	auth.Post("/register", authController.Register)
	auth.Post("/login", authController.Login)

	// Rutas protegidas
	auth.Post("/logout", middleware.AuthMiddleware(authService), authController.Logout)
	auth.Get("/validate", middleware.AuthMiddleware(authService), authController.ValidateSession)
}
