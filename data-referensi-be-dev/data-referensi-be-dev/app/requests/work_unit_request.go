package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type WorkUnitRequest struct {
	Code     string `json:"code" validate:"required,max=5"`
	Acronym  string `json:"acronym" validate:"required,max=50"`
	Fullname string `json:"fullname" validate:"required"`
}

func ValidateWorkUnit(c *fiber.Ctx) error {
	var req WorkUnitRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
