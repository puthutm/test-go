package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ScienceRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" `
}

func ValidateScience(c *fiber.Ctx) error {
	var req ScienceRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
