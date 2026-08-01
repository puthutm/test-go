package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type SksWeightRequest struct {
	Amount int `json:"amount" validate:"required,number"`
}

func ValidateSksWeight(c *fiber.Ctx) error {
	var req SksWeightRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
