package controllers

import (
	"register/models/dtos"
	"register/models/interfaces"

	"github.com/gofiber/fiber/v2"
)

// type AuthController struct {
//   Register()
// }

type AuthController struct {
	service interfaces.AuthService
}

func NewAuthController(service interfaces.AuthService) *AuthController {
	return &AuthController{
		service: service,
	}
}

func (controller *AuthController) Router(c fiber.Router) {
	c.Post("/register", controller.Register)
	// c.Post("/register", controller.Register)
}

func (controller *AuthController) Register(c *fiber.Ctx) error {
	request := dtos.RegisterRequest{}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	user, err := controller.service.Register(request)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})

}
