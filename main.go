package main

import (
	"register/app"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Output:
	// Hello, world.
	println("Hello, world.")

	// Start new fiber app
	server := fiber.New()

	app.SetupRoutes(server)

	// Listen on PORT 3000
	server.Listen(":3000")

}
