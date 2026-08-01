package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type RegistrationStatusRequest struct {
	Name           string `json:"name" validate:"required,max=255"`
	IsDefault      bool   `json:"is_default" validate:"boolean"`
	AcceptedMarker bool   `json:"accepted_marker" validate:"boolean"`
}

func ValidateRegistrationStatus(c *fiber.Ctx) error {
	var req RegistrationStatusRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
