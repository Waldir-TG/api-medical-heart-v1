// cmd/api/main.go
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Waldir-TG/api-medical-heart-v1/internal/db"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/repositories"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/routes"
	"github.com/Waldir-TG/api-medical-heart-v1/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
)

// @title Medical Heart API
// @version 1.0
// @description API for managing medical heart data
// @BasePath /
func main() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Inicializar conexión a base de datos
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	// Inicializar repositorios
	userRepo := repositories.NewUserRepository(database)
	sessionRepo := repositories.NewSessionRepository(database)
	doctorRepo := repositories.NewDoctorRepository(database)
	patientRepo := repositories.NewPatientRepo(database)
	deviceRepo := repositories.NewDeviceRepository(database)
	heartReadingRepo := repositories.NewHeartReadingRepository(database)

	// Inicializar servicios
	authService := services.NewAuthService(userRepo, sessionRepo)
	userService := services.NewUserService(userRepo)
	doctorService := services.NewDoctorService(doctorRepo)
	patientService := services.NewPatientService(patientRepo)
	deviceService := services.NewDeviceService(deviceRepo)
	heartReadingService := services.NewHeartReadingService(heartReadingRepo)

	// Configurar aplicación Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Ruta para servir el archivo swagger.json
	app.Get("/swagger.json", func(c *fiber.Ctx) error {
		return c.SendFile("./docs/swagger.json")
	})

	// Ruta para servir RapiDoc (HTML interactivo)
	app.Get("/docs", func(c *fiber.Ctx) error {
		html := `
		<!doctype html>
		<html>
			<head>
				<meta charset="utf-8">
				<script 
					type="module" 
					src="https://unpkg.com/rapidoc/dist/rapidoc-min.js"
				></script>
			</head>
			<body>
				<rapi-doc 
					spec-url="/swagger.json" 
					theme="light"
				></rapi-doc>
			</body>
		</html>
		`
		c.Type("html")
		return c.SendString(html)
	})

	// Middlewares globales
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("CORS_ALLOW_ORIGINS"),
		AllowMethods:     "GET,POST,PUT,DELETE",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
	}))

	// Configurar rutas
	routes.SetupAuthRoutes(app, authService)
	routes.SetupUserRoutes(app, authService, userService)
	routes.SetupDoctorRoutes(app, doctorService)
	routes.SetupPatientRoutes(app, authService, patientService)
	routes.SetupDeviceRoutes(app, authService, deviceService)
	routes.SetupHeartReadingRoutes(app, authService, heartReadingService)
	// Iniciar servidor
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
		}

		if err := app.Listen(":" + port); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}
}
