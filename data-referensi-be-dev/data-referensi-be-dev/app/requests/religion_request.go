package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ReligionRequest struct {
	Code string `json:"code" validate:"required,max=2"`
	Name string `json:"name" validate:"required,max=255"`
}

func ValidateReligion(c *fiber.Ctx) error {
	var req ReligionRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
