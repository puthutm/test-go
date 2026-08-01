package requests

import (
	"data-referensi/handlers"

	"github.com/gofiber/fiber/v2"
)

type ReportCardAssessmentRequest struct {
	Code      string  `json:"code" validate:"required,max=100"`
	Name      string  `json:"name" validate:"required,max=255"`
	Value     float64 `json:"value" validate:"required"`
	SubjectId string  `json:"subject_id" validate:"required,max=36"`
	Note      string  `json:"note" validate:"required"`
}

func ValidateReportCardAssessment(c *fiber.Ctx) error {
	var req ReportCardAssessmentRequest

	errorMessages, err := handlers.ModelValidate(c, &req)
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"error":   true,
			"message": errorMessages,
		})
	}

	return c.Next()
}
