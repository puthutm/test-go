package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type SubjectRequest struct {
	Code         string  `json:"code" validate:"required,max=100"`
	Name         string  `json:"name" validate:"required,max=255"`
	MinimumPoint float64 `json:"minimum_point" validate:"required"`
}

func ValidateSubject(c *fiber.Ctx) error {
	var req SubjectRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
