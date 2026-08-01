package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ReferralRequest struct {
	Code        string `json:"code" validate:"required,max=50"`
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=255"`
	Status      string `json:"status" validate:"required,max=50"`
}

func ValidateReferral(c *fiber.Ctx) error {
	var req ReferralRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
