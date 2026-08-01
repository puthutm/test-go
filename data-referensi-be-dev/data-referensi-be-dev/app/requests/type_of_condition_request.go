package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type TypeOfConditionRequest struct {
	Code            string `json:"code" validate:"required,max=100"`
	TypeOfCondition string `json:"type_of_condition" validate:"required,max=255"`
	Note            string `json:"note" validate:"required"`
}

func ValidateTypeOfCondition(c *fiber.Ctx) error {
	var req TypeOfConditionRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
