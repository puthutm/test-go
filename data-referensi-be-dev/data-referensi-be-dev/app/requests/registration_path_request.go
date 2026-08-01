package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type RegistrationPathRequest struct {
	Code             string `json:"code" validate:"required,max=255"`
	RegistrationPath string `json:"registration_path" validate:"required,max=255"`
	RegistrationType string `json:"registration_type" validate:"required,max=255"`
	Description      string `json:"description" `
}

func ValidateRegistrationPath(c *fiber.Ctx) error {
	var req RegistrationPathRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
