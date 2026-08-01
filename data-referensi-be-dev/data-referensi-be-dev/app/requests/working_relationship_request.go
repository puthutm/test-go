package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type WorkingRelationshipRequest struct {
	Code string `json:"code" validate:"required,max=3"`
	Name string `json:"name" validate:"required,max=255"`
}

func ValidateWorkingRelationship(c *fiber.Ctx) error {
	var req WorkingRelationshipRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
