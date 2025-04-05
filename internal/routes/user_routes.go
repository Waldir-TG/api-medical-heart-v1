// internal/routes/user_routes.go
package routes

import (
	"github.com/Waldir-TG/api-medical-heart-v1/internal/controllers"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/middleware"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App, authService *services.AuthService, userService *services.UserService) {
	// Crear controlador
	userController := controllers.NewUserController(userService)

	// Grupo de rutas para usuarios
	users := app.Group("/api/users", middleware.AuthMiddleware(authService))

	// Rutas de usuarios
	users.Get("/profile", userController.GetProfile)

	// Rutas protegidas por rol
	admin := users.Group("/admin", middleware.RoleMiddleware("admin"))
	admin.Get("/", userController.GetAllUsers)
	admin.Post("/", userController.CreateUser)
	admin.Put("/:id", userController.UpdateUser)
	admin.Delete("/:id", userController.DeleteUser)

	// Rutas específicas para médicos
	doctor := users.Group("/doctor", middleware.RoleMiddleware("doctor"))
	doctor.Get("/patients", userController.GetDoctorPatients)
}
