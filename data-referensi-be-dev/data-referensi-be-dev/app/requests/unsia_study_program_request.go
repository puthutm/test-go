package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type UnsiaStudyProgramRequest struct {
	Code string `json:"code" validate:"required,max=3"`
	Name string `json:"name" validate:"required,max=255"`
}

func ValidateUnsiaStudyProgram(c *fiber.Ctx) error {
	var req UnsiaStudyProgramRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
