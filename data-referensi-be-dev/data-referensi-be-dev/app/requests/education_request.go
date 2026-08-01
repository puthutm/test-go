package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type EducationRequest struct {
	EducationalLevelId string `json:"educational_level_id" validate:"required,max=36"`
	StudyProgramId     string `json:"study_program_id" validate:"max=36"`
	Name               string `json:"name" validate:"required,max=255"`
}

func ValidateEducation(c *fiber.Ctx) error {
	var req EducationRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
