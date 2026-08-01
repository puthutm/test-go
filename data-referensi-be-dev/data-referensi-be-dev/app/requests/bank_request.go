package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type BankRequest struct {
	Code string `json:"code" validate:"required,max=12"`
	Name string `json:"name" validate:"required,max=255"`
}

func ValidateBank(c *fiber.Ctx) error {
	var req BankRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
