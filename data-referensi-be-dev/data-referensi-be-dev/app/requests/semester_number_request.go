package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type SemesterNumberRequest struct {
	SemesterNumber string `json:"semester_number" validate:"required"`
}

func ValidateSemesterNumber(c *fiber.Ctx) error {
	var req SemesterNumberRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
