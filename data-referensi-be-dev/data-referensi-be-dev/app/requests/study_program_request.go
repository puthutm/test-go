package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type StudyProgramRequest struct {
	Code string `json:"code" validate:"required,max=5"`
	Name string `json:"name" validate:"required,max=255"`
	Type string `json:"type" validate:"required,max=255"`
}

func ValidateStudyProgram(c *fiber.Ctx) error {
	var req StudyProgramRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
