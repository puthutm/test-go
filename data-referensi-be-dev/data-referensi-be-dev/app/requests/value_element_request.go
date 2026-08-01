package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ValueElementRequest struct {
	Name        string  `json:"name" validate:"required,max=225"`
	Description *string `json:"description"`
}

func ValidateValueElement(c *fiber.Ctx) error {
	var req ValueElementRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
