package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type CurriculumYearRequest struct {
	Years       string  `json:"years" validate:"required,max=4"`
	Starts      string  `json:"starts" validate:"required"`
	StartDate   string  `json:"start_date" validate:"required"`
	EndDate     string  `json:"end_date" validate:"required"`
	Description *string `json:"description"`
}

func ValidateCurriculumYear(c *fiber.Ctx) error {
	var req CurriculumYearRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
