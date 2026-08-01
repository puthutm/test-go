package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ValueScaleRequest struct {
	Value       string  `json:"value" validate:"required,max=225"`
	Description *string `json:"description"`
}

func ValidateValueScale(c *fiber.Ctx) error {
	var req ValueScaleRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
