package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type GradeRequest struct {
	Code        string   `json:"code" validate:"required,max=3"`
	Name        string   `json:"name" validate:"required,max=255"`
	LowerLimit  *float64 `json:"lower_limit"`
	UpperLimit  *float64 `json:"upper_limit"`
	Description *string  `json:"description"`
}

func ValidateGrade(c *fiber.Ctx) error {
	var req GradeRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
