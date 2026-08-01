package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ProgramTypeRequest struct {
	Code        string `json:"code" validate:"required,max=255"`
	ProgramType string `json:"program_type" validate:"required,max=255"`
	Description string `json:"description" `
	IsIPC       bool   `json:"is_ipc" `
}

func ValidateProgramType(c *fiber.Ctx) error {
	var req ProgramTypeRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
