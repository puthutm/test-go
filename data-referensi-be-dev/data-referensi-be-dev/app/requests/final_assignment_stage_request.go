package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type FinalAssignmentStageRequest struct {
	Name        string  `json:"name" validate:"required,max=225"`
	Description *string `json:"description"`
}

func ValidateFinalAssignmentStage(c *fiber.Ctx) error {
	var req FinalAssignmentStageRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
