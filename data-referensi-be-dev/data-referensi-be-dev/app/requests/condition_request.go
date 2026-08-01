package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ConditionRequest struct {
	Code  string  `json:"code" validate:"required,max=100"`
	Name  string  `json:"name" validate:"required,max=255"`
	Point float64 `json:"point" validate:"required"`
}

func ValidateCondition(c *fiber.Ctx) error {
	var req ConditionRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
