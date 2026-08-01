package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type FieldStudyConcentrationRequest struct {
	Code           string `json:"code" validate:"required,max=3"`
	StudyProgramID string `json:"study_program_id" validate:"required,max=255"`
	Name           string `json:"name" validate:"required,max=255"`
	Description    string `json:"description" validate:"required,max=255"`
}

func ValidateFieldStudyConcentration(c *fiber.Ctx) error {
	var req FieldStudyConcentrationRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
