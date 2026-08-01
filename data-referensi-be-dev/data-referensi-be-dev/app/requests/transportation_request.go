package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type TransportationRequest struct {
	Code string `json:"code" validate:"required,max=50"`
	Name string `json:"name" validate:"required"`
}

func ValidateTransportation(c *fiber.Ctx) error {
	var req TransportationRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
