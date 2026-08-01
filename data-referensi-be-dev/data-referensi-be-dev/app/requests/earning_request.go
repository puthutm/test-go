package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type EarningRequest struct {
	Code  string `json:"code" validate:"required,max=50"`
	Name  string `json:"name" validate:"required,max=255"`
	Range string `json:"range" validate:"required,max=255"`
}

func ValidateEarning(c *fiber.Ctx) error {
	var req EarningRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
