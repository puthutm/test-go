package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type StudentStatusRequest struct {
	Name      string `json:"name" validate:"required,max=255"`
	IsDefault bool   `json:"is_default"`
}

func ValidateStudentStatus(c *fiber.Ctx) error {
	var req StudentStatusRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
