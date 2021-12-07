package app

import (
	"register/controllers"
	"register/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App) {
	// Connect to the database
	db := NewDB()

	// Middleware
	api := app.Group("/api", logger.New())

	// Routes
	auth := api.Group("/auth")
	authService := service.NewAuthService(db)
	authController := controllers.NewAuthController(authService)

	authController.Router(auth)

}
